package config

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	appPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	appPath = strings.ReplaceAll(appPath, "\\bin", "")
	appPath = strings.ReplaceAll(appPath, "\\main", "")
	appPath = strings.ReplaceAll(appPath, "/bin", "")
	appPath = strings.ReplaceAll(appPath, "/main", "")
	appPath = strings.ReplaceAll(appPath, "\\", "/")

	os.Setenv("APP_PATH", appPath+"/")

	err = godotenv.Load(appPath + ".env")
	if err != nil {
		err = godotenv.Load(appPath + "/.env")
		if err != nil {
			err = godotenv.Load(appPath + "\\.env")
			if err != nil {
				log.Fatal("Error loading .env file path: ", appPath)
			}
		}
	}
}
