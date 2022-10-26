package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func GetConfig() {
	mainPath, err := os.Getwd()
	if err != nil {
		fmt.Errorf("could not get mainPath: %w", err)
	}
	err = godotenv.Load(mainPath + "/pkg/config/config.env")
	if err != nil {
		fmt.Errorf("error loading .env file: %w", err)
	}

}
