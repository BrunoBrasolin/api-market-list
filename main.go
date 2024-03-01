package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MarketList struct {
	ID    int `json:"id"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

var lists = []MarketList{
	{ID: 1, Month: 01, Year: 2024},
	{ID: 1, Month: 02, Year: 2024},
	{ID: 1, Month: 03, Year: 2024},
}

func get(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, lists)
}

func main() {
	router := gin.Default()
	router.GET("/", get)

	router.Run("localhost:8080")
}
