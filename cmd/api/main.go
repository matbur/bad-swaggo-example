package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"swaggo/internal/users"
)

// @title Insights API
// @version 1.0
// @description API that manages all the insights data

// @securityDefinitions.apikey Token
// @in header
// @name Authorization

// @query.collection.format multi

// @BasePath  /
func main() {
	router := chi.NewRouter()
	users.RegisterRoutes(router)

	server := &http.Server{
		Addr:    fmt.Sprintf(":8080"),
		Handler: router,
	}

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		panic(err)
	}
}
