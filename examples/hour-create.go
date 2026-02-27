package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ProductionPanic/gripp"
	"github.com/joho/godotenv"
)

const projectLine = 5723
const employeeID = 97403

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

	result, err := client.Hours().Create(gripp.HourCreateData{
		Offerprojectline: projectLine,
		Employee:         employeeID,
		Description:      "Worked on the project",
		Amount:           .5,
	})
	if err != nil {
		panic(err)
	}
	if result.Success {
		fmt.Printf("Hour created successfully with id %d\n", result.RecordId)
	} else {
		fmt.Printf("Hour creation failed \n")
	}
}
