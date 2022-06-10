package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/viggin543/awesomne-golang/11_gin/2_cart/internal/models"
	"github.com/viggin543/awesomne-golang/11_gin/2_cart/internal/module"
	"net/http"
)

func cartUpdateHandler(m module.Module) func(c *gin.Context) {
	return func(c *gin.Context) {
		cart := getCart(c)
		logrus.WithField("cart", cart).Info("upserting user cart")
		if err := m.CartService.UpdateCart(c, cart); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, cart)
	}
}

func getCart(c *gin.Context) models.Cart {
	_cart, exists := c.Get("cart")
	if !exists {
		c.AbortWithStatus(http.StatusBadRequest)
		panic("duck..")
	}
	return _cart.(models.Cart)
}

func parseCartMiddleware(c *gin.Context) {
	req := models.Cart{}
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.Set("cart", req)
}
