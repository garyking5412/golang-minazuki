package goService

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"golang-minazuki/models"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func CreateCategory(ctx *gin.Context, db *gorm.DB, rdb *redis.Client) {
	log.Printf(">>> creating category")
	var newCategory models.Category
	if err := ctx.BindJSON(&newCategory); err != nil {
		panic(err)
	}
	if err := db.Create(&newCategory).Error; err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if err := rdb.Set(ctx, newCategory.Name, newCategory, 0).Err(); err != nil {
		panic(err)
	}
	ctx.IndentedJSON(http.StatusCreated, newCategory)
}

func CachingCategory(ctx *gin.Context, rdb *redis.Client) {
	log.Printf(">>> caching category")
	var newCategory models.Category
	if err := ctx.BindJSON(&newCategory); err != nil {
		panic(err)
	}
	category, err := json.Marshal(newCategory)
	if err != nil {
		log.Fatalf("Could not serialize user: %v", err)
	}
	if err := rdb.Set(ctx, newCategory.Name, category, 0).Err(); err != nil {
		panic(err)
	}
	ctx.IndentedJSON(http.StatusCreated, newCategory)
}

func GetAllCategory(c *gin.Context, db *gorm.DB) {
	var allCategory []models.Category
	if err := db.Find(&allCategory).Error; err != nil {
		log.Printf("Failed to get all category: %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, allCategory)
}

func GetCategoryById(c *gin.Context, rdb *redis.Client) {
	log.Printf("Getting category by id from redis cache: %v", c.Param("key"))
	id := c.DefaultQuery("key", "")
	val, err := rdb.Get(c, id).Result()
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
	}
	var result models.Category
	err = json.Unmarshal([]byte(val), &result)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error!"})
	}
	c.IndentedJSON(http.StatusOK, result)
}
