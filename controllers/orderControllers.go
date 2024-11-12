package controllers

import (
	"net/http"
	"rest-api/database"
	"rest-api/models"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var db = database.GetDB()
	var order models.Order

	// rawData, _ := c.GetRawData()
	// fmt.Println("Request Body:", string(rawData))

	// Bind JSON request to the Order struct
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Save order in DB, need the ID
	if err := db.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create order",
		})
		return
	}

	// Slap OrderID to each items in the request
	for i := range order.Items {
		order.Items[i].OrderID = order.ID
		order.Items[i].Model.ID = 0 // ChatGPT solution since I'm unsure of neater way to do this... Ensure ID is unset so it auto-generates
	}

	// Create items with the associated OrderID
	if err := db.Create(&order.Items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create order items",
		})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Order created successfully",
		"order":   order,
	})
}

func GetAllData(c *gin.Context) {
	var db = database.GetDB()
	var orders []models.Order

	err := db.Find(&orders).Error
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orders})
}
