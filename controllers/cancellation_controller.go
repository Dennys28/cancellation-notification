﻿package controllers

import (
	"cancellation-notification/config"
	"cancellation-notification/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCancellation(c *gin.Context) {
	var request models.Cancellation

	// Validar datos JSON
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Insertar la cancelación en la base de datos
	query := "INSERT INTO cancellations (reservation_id, email, reason, created_at) VALUES (?, ?, ?, NOW())"
	_, err := config.DB.Exec(query, request.ReservationID, request.Email, request.Reason)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al registrar la cancelación"})
		return
	}

	// Simular el envío de notificación
	fmt.Printf("Enviando notificación de cancelación a %s para la reserva %d. Motivo: %s\n",
		request.Email, request.ReservationID, request.Reason)

	c.JSON(http.StatusOK, gin.H{"message": "Notificación de cancelación enviada"})
}
