// @title EDEN ARC (IAM) API
// @version 1.0
// @description Identity and Access Management API.
package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-arc/internal/config"
	"github.com/lubosgarancovsky/eden-arc/internal/db"
	"github.com/lubosgarancovsky/eden-arc/internal/router"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/lubosgarancovsky/eden-arc/docs"
)

func main() {
	cfg := config.LoadConfig()
	gormDB := db.ConnectDB(cfg.DBUrl)

	r := gin.Default()
	r.GET("/api/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.SetupRouter(r, gormDB, cfg)

	err := r.Run(fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatal(err)
	}
}
