package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/viggin543/awesomne-golang/11_gin/2_cart/internal/models"
	"github.com/viggin543/awesomne-golang/11_gin/2_cart/internal/module"
	"github.com/viggin543/awesomne-golang/11_gin/2_cart/internal/server"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Info("cart service starting")
	cfg := models.ParseConfig()
	cartModule := module.Create(cfg)
	engine := gin.Default()
	server.SetupRoutes(engine, cartModule)
	_ = engine.Run()
}
