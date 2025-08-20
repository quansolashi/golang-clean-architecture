package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	wrapper "github.com/rs/cors/wrapper/gin"
)

func NewOptions() cors.Options {
	return cors.Options{
		AllowedOrigins: []string{
			"http://*",
			"https://*",
		},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"User-Agent",
			"X-Forwarded-For",
			"X-Forwarded-Proto",
			"X-Real-Ip",
		},
		AllowCredentials:   true,
		MaxAge:             1440, // 60m * 24h
		OptionsPassthrough: false,
		Debug:              false,
	}
}

func NewGinMiddleware() gin.HandlerFunc {
	options := NewOptions()
	return wrapper.New(options)
}
