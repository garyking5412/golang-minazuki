package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

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

func main() {

	//set up database connection
	connectDatabase()

	//set up application
	rout := gin.Default()
	rout.GET("/minazuki", func(ctx *gin.Context) {
		getAllCategory(ctx)
	})
	rout.POST("/minazuki", createCategory)
	rout.Run(":3004")
}
