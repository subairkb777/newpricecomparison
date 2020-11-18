package controllers

import (
	"net/http"

	"../services"
	"github.com/gin-gonic/gin"
)

type ItemReq struct {
	ItemName    string `json:"item_name"`
	MarketPlace string `json:"market_place"`
}

func GetItem(c *gin.Context) {
	var req ItemReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})
		return
	}

	// itemData, err := services.GetItem(req.ItemName, req.MarketPlace)
	itemData, err := services.GetItem(req.ItemName, req.MarketPlace)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, itemData)
}
