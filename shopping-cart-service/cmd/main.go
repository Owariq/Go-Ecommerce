package main

import (
	"log"

	"github.com/Owariq/Go-Ecommerce/shopping-cart-service/internal/app"
)

func main() {

	a, err := app.NewApp()
	if err != nil {
		log.Fatalf("failed to init app: %v", err)
	}
	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %v", err)
	}
}
