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

	// fetching employees
	employees, err := client.Employee().
		Filter("active", true).
		Get()
	if err != nil {
		panic(err)
	}
	fmt.Println("Employees:")
	for _, employee := range employees {
		fmt.Println("Name:", employee.Searchname)
		fmt.Println("Id:", employee.ID)

		utils.Hr()
	}
}
