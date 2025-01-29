package config

import (
	"log"
	"github.com/joho/godotenv"
)

// loadEnv carga las variables del archivo .env
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}
}
