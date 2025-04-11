package LocalService

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"golang-minazuki/global"
	"golang-minazuki/models"
	"log"
)

func HandleWSMessage(messageType int, data []byte) string {
	if messageType == websocket.TextMessage {
	}
	var wsCategory models.Category
	var err error
	if err := json.Unmarshal(data, &wsCategory); err != nil {
		log.Printf("Failed to parse message: %v", err)
	}
	if err := global.Ctx.DatabaseConnection.Create(&wsCategory).Error; err != nil {
		log.Printf("Failed to create new category: %v", err)
	}
	response, err := json.Marshal(wsCategory)
	if err != nil {
		log.Printf("Failed to marshal new category: %v", err)
	}
	return string(response)
}
