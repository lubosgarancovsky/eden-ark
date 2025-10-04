package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-arc/internal/config"
	"github.com/lubosgarancovsky/eden-arc/internal/handler"
	"github.com/lubosgarancovsky/eden-arc/internal/middleware"
	"github.com/lubosgarancovsky/eden-arc/internal/repository"
	"github.com/lubosgarancovsky/eden-arc/internal/service"
	"github.com/lubosgarancovsky/go-kit/rsql"
	"gorm.io/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB, cfg *config.Config) {
	parser := rsql.New()

	// Clients
	clientRepo := repository.NewClientRepository(db)
	clientService := service.NewClientService(clientRepo)
	clientHandler := handler.NewClientHandler(parser, clientService)

	// Secrets
	clientSecretRepo := repository.NewClientSecretRepository(db)
	clientSecretService := service.NewClientSecretService(clientSecretRepo)
	clientSecretHandler := handler.NewClientSecretHandler(parser, clientSecretService)

	v1 := r.Group("/v1/arc")
	v1.Use(middleware.ErrorMiddleware())

	protected := v1.Use(middleware.AuthMiddleware())
	{
		protected.GET("/admin/clients", clientHandler.FindAll)
		protected.GET("/admin/clients/:clientId", clientHandler.FindByID)
		protected.POST("/admin/clients", clientHandler.Create)
		protected.PUT("/admin/clients/:clientId", clientHandler.Update)
		protected.DELETE("/admin/clients/:clientId", clientHandler.Delete)
	}
	{
		protected.GET("/admin/clients/:clientId/secrets", clientSecretHandler.FindAll)
		protected.POST("/admin/clients/:clientId/secrets", clientSecretHandler.Create)
		protected.DELETE("/admin/clients/:clientId/secrets/:clientSecretId", clientSecretHandler.Delete)
	}
}
