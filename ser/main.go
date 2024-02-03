package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type juice struct {
	ID    string  `json:"id"`
	Type  string  `json:"type"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var juices = []juice{
	{ID: "1", Type: "Fruit Juice", Name: "Orange Delight", Price: 3.49},
	{ID: "2", Type: "Vegetable Juice", Name: "Tomato Explosion", Price: 2.99},
	{ID: "3", Type: "Mixed Juice", Name: "Tropical Energy", Price: 4.99},
	{ID: "4", Type: "Berry Juice", Name: "Strawberry Sweetness", Price: 3.79},
}

func main() {
	router := gin.Default()
	router.GET("/juices", getJuices)
	router.GET("/juices/:id", getJuiceByID)
	router.POST("/juices", postJuice)
	router.DELETE("/juices/:id", deleteJuiceByID)
	router.PUT("/juices/:id", updateJuiceByID)
	router.Run("localhost:8080")
}

func getJuices(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, juices)
}

func postJuice(c *gin.Context) {
	var newJuice juice

	if err := c.BindJSON(&newJuice); err != nil {
		return
	}

	juices = append(juices, newJuice)
	c.IndentedJSON(http.StatusCreated, newJuice)
}

func getJuiceByID(c *gin.Context) {
	id := c.Param("id")

	for _, j := range juices {
		if j.ID == id {
			c.IndentedJSON(http.StatusOK, j)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "juice not found"})
}

func deleteJuiceByID(c *gin.Context) {
	id := c.Param("id")

	for i, j := range juices {
		if j.ID == id {

			juices = append(juices[:i], juices[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "juice deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "juice not found"})
}

func updateJuiceByID(c *gin.Context) {
	id := c.Param("id")

	for i, j := range juices {
		if j.ID == id {
			var updatedJuice juice

			if err := c.BindJSON(&updatedJuice); err != nil {
				return
			}

			juices[i] = updatedJuice
			c.IndentedJSON(http.StatusOK, updatedJuice)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "juice not found"})
}
