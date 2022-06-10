package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func accessLoggMiddleware(c *gin.Context) {
	logrus.WithField("request", c.Request.RequestURI).Info("Logging request")
}

func monitorMiddleware(c *gin.Context) {
	c.Next()
	logrus.WithFields(logrus.Fields{
		"request": c.Request.RequestURI,
		"failed":  c.IsAborted(),
	}).Info("Monitoring Request")
}
