package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	r := gin.Default()
	r.PUT("/cart", accessLoggMiddleware, monitorMiddlware, parseCart, cartUpdateHandler)
	_ = r.Run()
}

func parseCart(c *gin.Context) {
	req := Cart{}
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.Set("cart", req)
}
func cartUpdateHandler(c *gin.Context) {
	cart := getCart(c)
	logrus.WithField("cart", cart).Info("inserting")
	c.JSON(http.StatusOK, cart)
}

func getCart(c *gin.Context) Cart {
	_cart, exists := c.Get("cart")
	if !exists {
		c.AbortWithStatus(http.StatusBadRequest)
		panic("duck..")
	}
	return _cart.(Cart)
}

func accessLoggMiddleware(c *gin.Context) {
	logrus.WithField("request", c.Request.RequestURI).Info("Logging request")
}

func monitorMiddlware(c *gin.Context) {
	c.Next()
	logrus.WithFields(logrus.Fields{
		"request": c.Request.RequestURI,
		"failed":  c.IsAborted(),
	}).Info("Monitoring Request")
}
