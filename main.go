package main

import (
	"fmt"
	"log"

	"github.com/dudemous17/gator/internal/config"
)

func main() {
	// Read the config file
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	// Set the current user and update the file
	err = cfg.SetUser("dudemous")
	if err != nil {
		log.Fatalf("error setting user: %v", err)
	}
	fmt.Println("Config updated successfully.")

	// Read again and print results
	updatedCfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading updated config: %v", err)
	}

	fmt.Printf("Current Config Struct: %+v\n", updatedCfg)
}
