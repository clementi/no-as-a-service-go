package main

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"os"

	"github.com/gin-gonic/gin"
)

func setupRouter(reasons []string) *gin.Engine {
	r := gin.Default()

	r.GET("/api/no", func(c *gin.Context) {
		reason := reasons[rand.Intn(len(reasons))]

		c.JSON(200, gin.H{"reason": reason})
	})

	return r
}

func main() {
	reasons, err := loadReasons()
	if err != nil {
		log.Fatal("Could not load reasons")
	}

	r := setupRouter(reasons)
	r.Run(":8080")
}

func loadReasons() ([]string, error) {
	reasonsFile, err := os.Open("reasons.json")

	if err != nil {
		return nil, err
	}

	defer reasonsFile.Close()

	bytes, err := io.ReadAll(reasonsFile)
	if err != nil {
		return nil, err
	}

	var reasons []string

	json.Unmarshal(bytes, &reasons)

	return reasons, nil
}
