package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shakezidin/interview"
)

var webhookURL = "https://webhook.site/5412ebdf-ed92-4e67-827c-97967ea85186"

func main() {
	engine := gin.Default()

	engine.POST("/test", Calling)

	engine.Run(":8080")
}

func Calling(c *gin.Context) {
	ch := make(chan map[string]interface{})
	go interview.InterviewTask(c, ch, webhookURL)
	log.Println("waiting for responses")
	ans := <-ch
	c.JSON(http.StatusAccepted, gin.H{
		"data": ans,
	})
}

