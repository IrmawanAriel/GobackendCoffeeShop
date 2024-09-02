package main

import (
	"IrmawanAriel/goBackendCoffeeShop/internal/routers"
	"IrmawanAriel/goBackendCoffeeShop/pkg"
	"log"

	"github.com/gin-contrib/cors"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db, err := pkg.Posql()
	if err != nil {
		log.Fatal(err)
	}

	router := routers.New(db)

	// CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                             // Izinkan semua origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},  // Metode HTTP yang diizinkan
		AllowHeaders:     []string{"Content-Type", "Authorization"}, // Header yang diizinkan
		AllowCredentials: true,
	}))

	server := pkg.Server(router)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
