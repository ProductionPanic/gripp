package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ProductionPanic/gripp"
	"github.com/joho/godotenv"
)

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

	// get all active projects for a certain employee
	projects, err := client.Projects().
		Archived(false).
		Search("searchname", "minitrekker").
		Get()

	if err != nil {
		panic(err)
	}

	for _, project := range projects {
		fmt.Printf("- %d: %s - %s\n", project.ID, project.Company.SearchName, project.Name)
	}
}
