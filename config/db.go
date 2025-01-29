package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
		// Definir manualmente las variables de entorno
	os.Setenv("DB_USER", "admin")
	os.Setenv("DB_PASSWORD", "Hola1244")
	os.Setenv("DB_HOST", "18.212.223.216")
	os.Setenv("DB_NAME", "reservation_db")

	// Obtener las variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// Construir la cadena de conexión
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbHost, dbName)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Error al verificar conexión con la base de datos:", err)
	}

	fmt.Println("Conexión exitosa con la base de datos")
}
