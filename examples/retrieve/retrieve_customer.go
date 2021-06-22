package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tuneuptechnology/tuneuptechnology-go"
)

func main() {
	client := tuneuptechnology.New(os.Getenv("API_EMAIL"), os.Getenv("API_KEY"))

	customer := client.RetrieveCustomer(23)

	prettyJSON, err := json.MarshalIndent(customer, "", "    ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating JSON:", err)
	}
	fmt.Printf("%s\n", string(prettyJSON))
}
