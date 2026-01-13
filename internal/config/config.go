package config

import (
	"log"

	"github.com/lubosgarancovsky/go-kit"
)

type Config struct {
	Port             int    `field:"PORT" default:"9091"`
	PublicURL        string `field:"PUBLIC_URL"`
	DBUrl            string `field:"DB_URL"`
	PrivateKeyPath   string `field:"PRIVATE_KEY_PATH"`
	PublicKeyPath    string `field:"PUBLIC_KEY_PATH"`
	Issuer           string `field:"ISSUER"`
	AccessExp        int    `field:"ACCESS_EXP"`
	RefreshExp       int    `field:"REFRESH_EXP"`
	SessionExp       int    `field:"SESSION_EXP"`
	SMTPHost         string `field:"SMTP_HOST"`
	SMTPPort         int    `field:"SMTP_PORT"`
	SMTPFrom         string `field:"SMTP_FROM"`
	RecoveryTokenExp int    `field:"RECOVERY_TOKEN_EXP"`
	TemplatesFolder  string `field:"TEMPLATES_FOLDER"`
	UploadsFolder    string `field:"UPLOADS_FOLDER"`
}

func LoadConfig() *Config {
	var appConfig Config
	if err := go_kit.LoadEnv(&appConfig); err != nil {
		log.Fatal("Failed to load config from .env file", err)
	}

	return &appConfig
}
