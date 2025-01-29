﻿package main

import (
	"cancellation-notification/config"
	"cancellation-notification/routes"
)

func main() {
	// Iniciar la conexión a la base de datos
	config.InitDB()

	// Configurar rutas y ejecutar el servidor
	r := routes.SetupRouter()
	r.Run(":8080")
}
