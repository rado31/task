package api

import (
	"task/src/api/schemas"

	"github.com/gin-gonic/gin"
)

func Controller(router *gin.RouterGroup) {
	router.GET("/:id", schemas.Validate_id, get_one)
	router.GET("/all", schemas.Validate_query, get_all)
	router.POST("/", schemas.Validate_create, create)
	router.PUT("/", schemas.Validate_update, update)
	router.DELETE("/:id", schemas.Validate_id, remove)
}
