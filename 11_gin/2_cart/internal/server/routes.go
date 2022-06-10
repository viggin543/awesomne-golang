package server

import (
	"github.com/gin-gonic/gin"
	"github.com/viggin543/awesomne-golang/11_gin/2_cart/internal/module"
)

func SetupRoutes(r *gin.Engine, m module.Module) {
	r.PUT(
		"/cart",
		accessLoggMiddleware,
		monitorMiddleware,
		parseCartMiddleware,
		cartUpdateHandler(m),
	)
}
