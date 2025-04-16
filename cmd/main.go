package main

import (
	"log"

	"github.com/rishavkumar7/ecommerce/config"
	"github.com/rishavkumar7/ecommerce/internals/apis/http/routes"
	databases "github.com/rishavkumar7/ecommerce/internals/infrastructure/databases/mongodb"
)

func main() {
	cfg, err := config.SetConfig()
	if err != nil {
		log.Fatalf("Error in loading the config: %v", err)
	}
	client, err := databases.NewDb(cfg)
	if err != nil {
		log.Fatalf("Error in connecting to the database: %v", err)
	}
	router := routes.SetUpRouter(client)
	router.Run(":" + cfg.PORT)
}
