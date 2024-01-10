package main

import (
	"context"
	"log"
	"os"
)

func main() {
	index_f, err := os.Create("index.html")
	if err != nil {
		log.Fatalf("Failed to create index file: %v\n", err)
	}
	err = index().Render(context.Background(), index_f)
	if err != nil {
		log.Fatalf("Failed to write to index file: %v\n", err)
	}
	portfolio_f, err := os.Create("portfolio.html")
	if err != nil {
		log.Fatalf("Failed to create portfolio file: %v\n", err)
	}
	err = portfolio().Render(context.Background(), portfolio_f)
	if err != nil {
		log.Fatalf("Failed to write to portfolio file: %v\n", err)
	}
	log.Printf("All done!")
}
