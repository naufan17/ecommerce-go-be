package utils

import (
	"github.com/gin-gonic/gin"
)

var (
	ErrNotFound            = gin.H{"error": "Not found"}
	ErrBadRequest          = gin.H{"error": "Bad request"}
	ErrInternalServerError = gin.H{"error": "Internal server error"}
	ErrUnauthorized        = gin.H{"error": "Unauthorized"}
	ErrForbidden           = gin.H{"error": "Forbidden"}
	ErrConflict            = gin.H{"error": "Conflict"}
	ErrInvalidCredentials  = gin.H{"error": "Invalid credentials"}
)
