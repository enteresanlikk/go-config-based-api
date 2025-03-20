package main

import (
	"fmt"
	"log"

	"go-config-based-api/internal/config"
	"go-config-based-api/internal/handlers"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func main() {
	// Initialize config loader
	config.GetInstance()

	// Create router
	r := router.New()

	// Register routes
	r.GET("/configs", handlers.GetAllConfigs)
	r.GET("/configs/{id}", handlers.GetConfigByID)

	// Create server
	server := &fasthttp.Server{
		Handler: r.Handler,
		Name:    "ConfigAPI",
	}

	// Start server
	port := 8081
	fmt.Printf("Server is running on http://localhost:%d\n", port)
	if err := server.ListenAndServe(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalf("Error in ListenAndServe: %v", err)
	}
}
