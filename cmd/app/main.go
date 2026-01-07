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
	r.StaticFile("/logo_16x16.svg", filepath.Join(dist, "logo_16x16.svg"))
	r.StaticFile("/logo_32x32.svg", filepath.Join(dist, "logo_32x32.svg"))
	r.StaticFile("/logo_48x48.svg", filepath.Join(dist, "logo_48x48.svg"))
	r.StaticFile("/logo_64x64.svg", filepath.Join(dist, "logo_64x64.svg"))
	r.StaticFile("/logo_180x180.svg", filepath.Join(dist, "logo_180x180.svg"))
	r.StaticFile("/logo_192x192.svg", filepath.Join(dist, "logo_192x192.svg"))
	r.StaticFile("/logo_512x512.svg", filepath.Join(dist, "logo_512x512.svg"))
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

	<-ctx.Done() // wait for a shutdown signal

	log.Println("shutting down")

	if gormDB != nil {
		sqlDB, _ := gormDB.DB()
		sqlDB.Close()
	}
}
