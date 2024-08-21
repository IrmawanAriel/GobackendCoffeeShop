package main

import (
	"IrmawanAriel/goBackendCoffeeShop/internal/routers"
	"IrmawanAriel/goBackendCoffeeShop/migrations/seed"
	"IrmawanAriel/goBackendCoffeeShop/pkg"
	"log"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db, err := pkg.Posql()
	if err != nil {
		log.Fatal(err)
	}
	if err := seed.SeedUsers(db); err != nil {
		log.Fatalln("Seeding failed:", err)
	}
	if err := seed.SeedProducts(db); err != nil {
		log.Fatalln("Seeding failed:", err)
	}

	router := routers.New(db)
	server := pkg.Server(router)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
