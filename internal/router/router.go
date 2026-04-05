package router

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-ark/internal/config"
	"github.com/lubosgarancovsky/eden-ark/internal/handler"
	"github.com/lubosgarancovsky/eden-ark/internal/middleware"
	"github.com/lubosgarancovsky/eden-ark/internal/repository"
	"github.com/lubosgarancovsky/eden-ark/internal/service"
	"github.com/lubosgarancovsky/go-kit"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB, cfg *config.Config) {
	parser := go_kit.NewRSQLParser()

	// Global middleware
	r.Use(middleware.ErrorMiddleware(), middleware.AppStateValidationMiddleware(db))

	// Email
	emailService := service.NewEmailService(cfg)

	// Clients
	clientRepo := repository.NewClientRepository(db)
	clientService := service.NewClientService(clientRepo)
	clientHandler := handler.NewClientHandler(parser, clientService)

	// Secrets
	clientSecretRepo := repository.NewClientSecretRepository(db)
	clientSecretService := service.NewClientSecretService(clientSecretRepo)
	clientSecretHandler := handler.NewClientSecretHandler(parser, clientSecretService)

	// Authorization code
	authCodeRepo := repository.NewAuthCodeRepository(db)
	authCodeService := service.NewAuthCodeService(authCodeRepo)

	// Session
	sessionRepo := repository.NewSessionRepository(db)
	sessionService := service.NewSessionService(cfg, sessionRepo)

	// Recovery token
	recoveryTokenRepo := repository.NewRecoveryTokenRepository(db)
	recoveryTokenService := service.NewRecoveryTokenService(cfg, recoveryTokenRepo)

	// User
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(cfg, userRepo, recoveryTokenService, emailService)
	userHandler := handler.NewUserHandler(parser, userService)

	// User avatar
	userAvatarRepo := repository.NewUserAvatarRepository(db)
	userAvatarService := service.NewUserAvatarService(cfg, userService, userAvatarRepo)
	userAvatarHandler := handler.NewUserAvatarHandler(userAvatarService)

	// OAuth2
	oauth2Service := service.NewOAuthService(cfg, clientService, clientSecretService, sessionService, userService, authCodeService)
	oauth2Handler := handler.NewOAuth2Handler(oauth2Service)

	v1 := r.Group("/v1/ark")
	protected := v1.Group("", middleware.AuthMiddleware())
	admin := protected.Group("/admin", middleware.RoleMiddleware([]string{"admin"}))

	// OAuth2 endpoints
	oauth2 := r.Group("/oauth2")
	{
		oauth2.GET("/authorize", oauth2Handler.Authorize)
		oauth2.POST("/token", oauth2Handler.Token)
		oauth2.POST("/login", oauth2Handler.Login)
		oauth2.GET("/logout", oauth2Handler.Logout)
		oauth2.GET("/profile", middleware.AuthMiddleware(), oauth2Handler.Profile)
		oauth2.POST("/register", oauth2Handler.Register)
	}

	// API endpoints
	adminClient := admin.Group("/clients")
	{
		adminClient.GET("", clientHandler.FindAll)
		adminClient.GET("/:clientId", clientHandler.FindByID)
		adminClient.POST("", clientHandler.Create)
		adminClient.PUT("/:clientId", clientHandler.Update)
		adminClient.DELETE("/:clientId", clientHandler.Delete)
	}

	adminClientSecret := adminClient.Group("/:clientId/secrets")
	{
		adminClientSecret.GET("", clientSecretHandler.FindAll)
		adminClientSecret.POST("", clientSecretHandler.Create)
		adminClientSecret.DELETE("/:clientSecretId", clientSecretHandler.Delete)
	}

	adminUser := admin.Group("/users")
	{
		adminUser.GET("", userHandler.FindAll)
		adminUser.GET("/:userId", userHandler.FindByID)
		adminUser.POST("", userHandler.Create)
		adminUser.PUT("/:userId", userHandler.Update)
		adminUser.DELETE("/:userId", userHandler.Delete)
		adminUser.DELETE("/:userId/generate-password", userHandler.GeneratePassword)
	}

	protectedUser := protected.Group("/users")
	{
		protectedUser.GET("/:userId", oauth2Handler.Profile)
		protectedUser.PUT("/:userId", userHandler.Update)
		protectedUser.DELETE("/:userId", userHandler.Delete)
		protectedUser.POST("/request-change-email", userHandler.RequestChangeEmail)
		protectedUser.POST("/change-email", userHandler.ChangeEmail)

		// Avatar
		protectedUser.POST("/avatar", userAvatarHandler.UploadAvatar)
		protectedUser.DELETE("/avatar", userAvatarHandler.RemoveAvatar)
	}

	publicUser := v1.Group("/users")
	{

		publicUser.GET("/username-available/:username", userHandler.IsUsernameAvailable)
		publicUser.GET("/email-available/:email", userHandler.IsUsernameAvailable)
		publicUser.POST("/request-reset-password", userHandler.RequestResetPassword)
		publicUser.POST("/reset-password", userHandler.ResetPassword)

	}

	// Static
	v1.Static("/avatars", filepath.Join(cfg.UploadsFolder, "avatars"))

	// Internal
	{
		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "ok",
			})
		})
	}
}
