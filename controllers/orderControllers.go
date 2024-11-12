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

	// // Create items with the associated OrderID
	// if err := db.Create(&order.Items).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": "Failed to create order items",
	// 	})
	// 	return
	// }

	// Return a success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Order created successfully",
		"order":   order,
	})
}

func GetAllData(c *gin.Context) {
	var db = database.GetDB()
	var orders []models.Order

	err := db.Preload("Items").Find(&orders).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orders})
}

func UpdateDataOrderAndItem(c *gin.Context) {
	var db = database.GetDB()
	var order models.Order

	// Bind JSON request to the Order struct
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := db.First(&order, "customer_name = ?", order.CustomerName).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	for _, item := range order.Items {
		// Find the existing item by item_code and order_id
		var existingItem models.Item
		if err := db.Where("item_code = ? AND order_id = ?", item.ItemCode, order.ID).First(&existingItem).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Item not found for update!"})
			return
		}

		// Update fields
		existingItem.Description = item.Description
		existingItem.Quantity = item.Quantity

		// Save the updated item
		if err := db.Save(&existingItem).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order and items updated successfully"})
}

func DeleteDataOrderAndItem(c *gin.Context) {
	var db = database.GetDB()
	var order models.Order

	// Bind JSON request to the Order struct
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Find order by customer name, while also getting the items associated with customer.
	err := db.Preload("Items").Where("customer_name = ?", order.CustomerName).First(&order).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Delete items first, as they are associated with the order
	if len(order.Items) > 0 {
		if err := db.Delete(&order.Items).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete items"})
			return
		}
	}

	// Delete the order itself
	if err := db.Delete(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order and associated items deleted successfully"})
}
