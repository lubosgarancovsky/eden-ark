// @title EDEN ARC (IAM) API
// @version 1.0
// @description Identity and Access Management API.
package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/gin-gonic/gin"
	_ "github.com/lubosgarancovsky/eden-ark/docs"
	"github.com/lubosgarancovsky/eden-ark/internal/config"
	"github.com/lubosgarancovsky/eden-ark/internal/db"
	"github.com/lubosgarancovsky/eden-ark/internal/router"
)

func main() {
	cfg := config.LoadConfig()
	gormDB := db.ConnectDB(cfg.DBUrl)

	r := gin.Default()

	router.SetupRouter(r, gormDB, cfg)

	// Load react client files
	dist := filepath.Join("web", "dist")
	r.Static("/assets", filepath.Join(dist, "assets"))
	r.StaticFile("/favicon.ico", filepath.Join(dist, "favicon.ico"))
	r.NoRoute(func(c *gin.Context) {
		c.File(filepath.Join(dist, "index.html"))
	})

	err := r.Run(fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatal(err)
	}
}
