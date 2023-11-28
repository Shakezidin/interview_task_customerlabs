package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shakezidin/interview"
)

func main() {
	engine := gin.Default()

	engine.POST("/test", Calling)

	engine.Run(":8080")

}

func Calling(c *gin.Context) {
	ch := make(chan map[string]interface{})
	go interview.InterviewTask(c, ch)
	log.Println("waiting for responseS")
	ans := <-ch
	c.JSON(http.StatusAccepted, gin.H{
		"data": ans,
	})
}
