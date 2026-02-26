package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ProductionPanic/gripp"
	"github.com/joho/godotenv"
)

const projectLine = 5578

// an example of how i want the api to work
func main() {
	_ = godotenv.Load()
	f, err := os.OpenFile("debug.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	log.SetOutput(f)

	client, err := gripp.NewClient(gripp.Config{
		ApiKey: os.Getenv("API_KEY"),
	})
	if err != nil {
		panic(err)
	}

	// fetching hours
	hours, err := client.Hours().
		ByProjectLineID(projectLine).
		Get()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hours for project line %d:\n", projectLine)
	total := 0.0
	for _, hour := range hours {
		fmt.Printf("- %d: %s - %f\n", hour.ID, hour.Description, hour.Amount)
		total += hour.Amount
	}
	fmt.Printf("Total hours: %f\n", total)
}
