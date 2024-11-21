package main

import (
	"github.com/gin-gonic/gin"
	"github.com/suresh-02/Iragu-booking/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"age":     "200",
		})
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:8080
}
