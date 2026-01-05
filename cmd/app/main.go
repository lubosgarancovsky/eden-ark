// @title EDEN ARC (IAM) API
// @version 1.0
// @description Identity and Access Management API.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/gin-gonic/gin"
	_ "github.com/lubosgarancovsky/eden-ark/docs"
	"github.com/lubosgarancovsky/eden-ark/internal/config"
	"github.com/lubosgarancovsky/eden-ark/internal/db"
	"github.com/lubosgarancovsky/eden-ark/internal/router"
)

func main() {
	fmt.Println("███████╗██████╗ ███████╗███╗   ██╗       █████╗ ██████╗ ██╗  ██╗\n██╔════╝██╔══██╗██╔════╝████╗  ██║      ██╔══██╗██╔══██╗██║ ██╔╝\n█████╗  ██║  ██║█████╗  ██╔██╗ ██║█████╗███████║██████╔╝█████╔╝ \n██╔══╝  ██║  ██║██╔══╝  ██║╚██╗██║╚════╝██╔══██║██╔══██╗██╔═██╗ \n███████╗██████╔╝███████╗██║ ╚████║      ██║  ██║██║  ██║██║  ██╗\n╚══════╝╚═════╝ ╚══════╝╚═╝  ╚═══╝      ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝\n                                                                ")
	cfg := config.LoadConfig()
	gormDB, err := db.ConnectDB(cfg.DBUrl)
	if err != nil {
		log.Println("DB unavailable, starting without DB:", err)
	}

	r := gin.Default()

	router.SetupRouter(r, gormDB, cfg)

	// Load react client files
	dist := filepath.Join("web", "dist")
	r.Static("/assets", filepath.Join(dist, "assets"))
	r.StaticFile("/favicon.ico", filepath.Join(dist, "favicon.ico"))
	r.NoRoute(func(c *gin.Context) {
		c.File(filepath.Join(dist, "index.html"))
	})

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		if err = r.Run(fmt.Sprintf(":%d", cfg.Port)); err != nil {
			log.Println(err)
			stop()
		}
	}()

	<-ctx.Done() // wait for shutdown signal

	log.Println("shutting down")

	if gormDB != nil {
		sqlDB, _ := gormDB.DB()
		sqlDB.Close()
	}
}
