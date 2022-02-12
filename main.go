package main

import (
	"encoding/json"
	"fmt"
	"github.com/scyanh/velocitylimits/db"
	"github.com/scyanh/velocitylimits/entities"
	"github.com/scyanh/velocitylimits/utils"
	"log"
)

func main() {
	// DB Connection
	db.Init()
	db.GetDb().AutoMigrate(&entities.AttemptRequest{})
	repository := entities.NewLoadInputRepository()

	// Output var
	var textLines []string

	// Read input
	loadedJson := utils.ReadFile(utils.InputPath)
	for _, jsonInput := range loadedJson {
		var input entities.Input
		err := json.Unmarshal([]byte(jsonInput), &input)
		if err != nil {
			log.Printf("Error occured during unmarshaling. Error: %s\n", err.Error())
			continue
		}

		attemptRequest, err := input.Format()
		if err != nil {
			log.Printf("Error occured formatting input. Error: %s\n", err.Error())
			continue
		}

		res := attemptRequest.Filter()

		previousInput := repository.FindInput(*attemptRequest)
		if previousInput != nil {
			log.Printf("Load ID:%s is observed more than once for a customer=%s", previousInput.InputID, previousInput.CustomerID)
			continue
		}

		_, err = repository.Save(*attemptRequest)
		if err != nil {
			log.Printf("Error occured to save input request. Error: %s\n", err.Error())
		}

		line, err := json.Marshal(res)
		if err != nil {
			fmt.Println(err)
			return
		}
		textLines = append(textLines, string(line))
	}

	if err := utils.WriteLines(textLines, utils.OutputPath); err != nil {
		log.Fatalf("Error writeLines: %s", err)
	}

	fmt.Println("done")
}
