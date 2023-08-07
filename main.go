package main

import (
	"cms-tizi/db"
	"cms-tizi/db/schema"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()
	db.Run()

	r.POST("/createTodo", createTodo)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func createTodo(c *gin.Context) {
	var todo schema.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the user record in the database
	db.DB.Create(&todo)
	c.JSON(http.StatusOK, gin.H{"message": "Todo created successfully", "user": todo})
}
