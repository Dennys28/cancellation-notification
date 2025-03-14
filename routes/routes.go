﻿package routes

import (
	"cancellation-notification/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Ruta de cancelación
	r.POST("/cancellation", controllers.CreateCancellation)

	return r
}
