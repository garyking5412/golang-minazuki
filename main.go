package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	"golang-minazuki/LocalService"
	"golang-minazuki/global"
	"golang-minazuki/models"
	service "golang-minazuki/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

var grpcClients sync.Map

type CategoryServiceServer struct {
	service.UnimplementedCategoryServiceServer
}

func (s *CategoryServiceServer) GetCategory(ctx context.Context, request *service.CategoryRequest) (*service.CategoryResponse, error) {
	log.Printf("Category Request Received: %v", request)
	return &service.CategoryResponse{
		Id:     request.GetId(),
		Name:   "thaitx",
		Detail: "test detail",
	}, nil
}

func (s *CategoryServiceServer) Chat(stream service.CategoryService_ChatServer) error {
	log.Print("Connection established >>>>>>>>>>")

	clientID := strconv.Itoa(int(time.Now().Unix()))

	welcome := &service.ChatMessage{
		Sender:    clientID,
		Content:   "Access Granted! Welcome to Golang Minazuki GRPC Server",
		Timestamp: time.Now().Unix(),
	}
	if err := stream.Send(welcome); err != nil {
		panic(err)
	}
	return nil
}

func (s *CategoryServiceServer) mustEmbedUnimplementedCategoryServiceServer() {
	//TODO implement me
	panic("implement me")
}

func initGrpcServer() {
	log.Printf("GRPC server INIT >>>")
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()

	service.RegisterCategoryServiceServer(srv, &CategoryServiceServer{})
	log.Printf("gRPC server listening at %v", lis.Addr())
	err = srv.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}
}

func initGrpcClient() (connection *grpc.ClientConn) {
	log.Printf("GRPC client INIT >>>")
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create gRPC client: %v", err)
	}
	port := conn.GetState().String()
	log.Printf("gRPC client listening at %v", port)
	return conn
}

func initLocalRedisConnection() (connection *redis.Client) {
	log.Printf("REDIS client INIT >>>")
	connection = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})
	var ctx = context.Background()
	pong, err := connection.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Println(pong)
	return connection
}

var db *gorm.DB

func connectDatabase() *gorm.DB {
	log.Printf("Database server connecting >>>")
	dsn := "host=localhost user=postgres password=2716 dbname=postgres port=5432 sslmode=disable search_path=public"
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}

func getCategory(c *gin.Context) {
	var allCategory []models.Category
	if err := db.Find(&allCategory).Error; err != nil {
		log.Printf("Failed to get all category: %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, allCategory)
}

func getCategoryById(c *gin.Context, connection *grpc.ClientConn) {
	log.Printf("Getting category by id: %v", c.Param("id"))
	id, err := strconv.Atoi(c.DefaultQuery("id", "0"))
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
	}
	conn := service.NewCategoryServiceClient(connection)
	response, err := conn.GetCategory(c, &service.CategoryRequest{
		Id: int32(id),
	})
	c.IndentedJSON(http.StatusOK, response)
}

func initGinServer() {
	log.Printf("GIN server INIT >>>")
	rout := gin.Default()
	port := ":3004"
	redisConnection := initLocalRedisConnection()
	rout.GET("/minazuki", func(ctx *gin.Context) {
		LocalService.GetAllCategory(ctx, db)
	})
	rout.GET("/minazuki/getCategoryByID", func(ctx *gin.Context) {
		LocalService.GetCategoryById(ctx, redisConnection)
	})
	rout.POST("/minazuki", func(ctx *gin.Context) {
		LocalService.CachingCategory(ctx, redisConnection)
	})
	err := rout.Run(port)
	if err != nil {
		log.Fatalf("Failed to start GIN server: %v", err)
	}
	log.Printf("GIN server listening at %v", port)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func initWebsocketServer() {
	http.HandleFunc("/ws", handleConn)
	go func() {
		log.Println("Starting server on :8081")
		if err := http.ListenAndServe(":8081", nil); err != nil {
			log.Fatal("ListenAndServe:", err)
		}
	}()
	select {}
}

func handleConn(writer http.ResponseWriter, request *http.Request) {
	log.Printf("Websocket server INIT >>>")
	conn, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
	}
	defer conn.Close()
	fmt.Println("Client Connected")

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Failed to read message: %v", err)
		}
		response := LocalService.HandleWSMessage(messageType, message)
		log.Printf("Echo: %v", response)
		if err := conn.WriteMessage(messageType, []byte(response)); err != nil {
			log.Printf("Failed to write message: %v", err)
			break
		}
	}
}

func main() {

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	//set up database connection
	go connectDatabase()

	//set up grpc server
	go initGrpcServer()

	//set up application
	go initGinServer()

	global.Ctx = &global.ApplicationContext{
		DatabaseConnection: db,
	}

	//initWebsocketServer()

	<-stopChan
}
