package main

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"os"
)

func GetAlive(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Yo I'm alive!",
	})
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	router := gin.Default()
	router.GET("/alive", GetAlive)

	err := router.Run(os.Getenv("HOST") + ":" + os.Getenv("PORT"))
	if err != nil {
		slog.Error("Error starting server: ", err)
		return
	}
}
