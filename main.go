package main

import (
	"fmt"
	"log"

	"github.com/Mayank3299/BlogAggregator/internal/config"
)

func main() {
	file, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config %+v\n", file)

	err = file.SetUser("mayank")
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	file, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config again %+v\n", file)
}
