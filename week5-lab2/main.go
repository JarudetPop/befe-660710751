package main

import (
	"github.com/gin-gonic/gin"
)

type Users struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func getUsers(c *gin.Context) {
	users := []Users{{ID: "1", Name: "Jarudet"}}
	c.JSON(200, users) // Marshal to JSON
}
func main() {
	r := gin.Default()
	r.GET("/users", getUsers)
	r.Run(":8080") // Run the server on port 8080
}
