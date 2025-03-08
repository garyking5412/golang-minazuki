package main

import (
	"context"
	"github.com/gin-gonic/gin"
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
	"syscall"
)

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

func (s *CategoryServiceServer) mustEmbedUnimplementedCategoryServiceServer() {
	//TODO implement me
	panic("implement me")
}

func initGrpcServer() {
	log.Printf("GRPC server INIT >>>")
	lis, err := net.Listen("tcp", ":50051")
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

type Category struct {
	gorm.Model
	Name   string `json:"name"`
	Detail string `json:"detail"`
}

func (Category) TableName() string {
	return "category"
}

var db *gorm.DB

func connectDatabase() *gorm.DB {
	log.Printf("Database server connecting >>>")
	dsn := "host=host.docker.internal user=postgres password=2716 dbname=postgres port=5432 sslmode=disable search_path=local"
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	db.AutoMigrate(&Category{})
	return db
}

func createCategory(c *gin.Context) {
	var newCategory Category
	if err := c.BindJSON(&newCategory); err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}
	if err := db.Create(&newCategory).Error; err != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}
	c.IndentedJSON(http.StatusCreated, newCategory)
}

func getAllCategory(c *gin.Context) {
	var allCategory []Category
	if err := db.Find(&allCategory).Error; err != nil {
		log.Printf("Failed to get all category: %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, allCategory)
}

func getCategory(c *gin.Context) {
	var allCategory []Category
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
	rout.GET("/minazuki", func(ctx *gin.Context) {
		getAllCategory(ctx)
	})
	rout.GET("/minazuki/getCategoryByID", func(ctx *gin.Context) {
		getCategoryById(ctx, initGrpcClient())
	})
	rout.POST("/minazuki", createCategory)
	err := rout.Run(port)
	if err != nil {
		log.Fatalf("Failed to start GIN server: %v", err)
	}
	log.Printf("GIN server listening at %v", port)
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

	<-stopChan
}
