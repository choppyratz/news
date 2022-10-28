package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func GetConfig() error {
	mainPath, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("could not get mainPath: %w", err)
	}
	err = godotenv.Load(mainPath + "/pkg/config/config.env")
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}
	return nil
}
