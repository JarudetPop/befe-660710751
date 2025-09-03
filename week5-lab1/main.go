package main

import (
	"github.com/gin-gonic/gin"
)

type Users struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {
		users := []Users{{ID: "1", Name: "Jarudet"}}
		c.JSON(200, users) //ทำ marshal เป็น JSON
	})
	r.Run(":8080") //รันเซิร์ฟเวอร์ที่พอร์ต 8080
}
