package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-arc/internal/config"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB, cfg *config.Config) {
	r.GET("/api/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
