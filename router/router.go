package router

import (
	"hotelPMS/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Handler(db *gorm.DB, router *gin.Engine, apiPath string, authService auth.Service) *gin.Engine {
	//api := router.Group(apiPath)

	return router
}
