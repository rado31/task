package src

import (
	"task/src/api"
	"task/src/middlewares"

	"github.com/gin-gonic/gin"
)

func Init_app() *gin.Engine {
	router := gin.New()
	router.SetTrustedProxies(nil)

	router.Use(gin.Logger())
	router.Use(middlewares.Cors)

	set_controllers(router)

	return router
}

func set_controllers(router *gin.Engine) {
	create_group_router(router, "users", api.Controller)
}

func create_group_router(
	router *gin.Engine,
	name string,
	controller func(router *gin.RouterGroup),
	middlewares ...gin.HandlerFunc,
) {
	group_name := router.Group(name, middlewares...)
	controller(group_name)
}
