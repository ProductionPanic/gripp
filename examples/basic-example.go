package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ProductionPanic/gripp"
	"github.com/ProductionPanic/gripp/examples/utils"
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

	// fetching projects
	projects, err := client.Projects().
		Filter("archived", false).
		Filter("name", "like", "%minitrekkers%").
		Get()
	if err != nil {
		panic(err)
	}

	for _, project := range projects {
		fmt.Println("Project:")
		fmt.Println("Name:", project.Name)
		fmt.Println("Id:", project.ID)

		fmt.Printf("Employees:\n")
		for _, employee := range project.Employees {
			fmt.Printf("- %s: %s\n", employee.ID, employee.SearchName)
		}

		fmt.Println("Project lines:")
		for _, line := range project.ProjectLines {
			fmt.Printf("- %d: %s\n", line.ID, line.Searchname)
		}
		utils.Br()
	}
}
