package main

import (
	"context"
	"fmt"
	"log"
	"os"

	sg "github.com/StackGuardian/sg-sdk-go/stackguardian"
)

func main() {
	// Create a configuration
	config := sg.DefaultConfig()
	config.APIKey = "apikey " + os.Getenv("SG_API_TOKEN")

	// You can customize the configuration
	// config.BaseURL = "https://api.app.stackguardian.io"
	// config.MaxRetries = 5

	// Create a new client
	client, err := sg.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Example: Read organization details
	ctx := context.Background()
	org, err := client.Organizations.ReadOrganization(ctx, "demo-org")
	if err != nil {
		if sg.IsNotFoundError(err) {
			log.Println("Organization not found")
		} else if sg.IsUnauthorizedError(err) {
			log.Println("Invalid API key or insufficient permissions")
		} else {
			log.Fatalf("Failed to read organization: %v", err)
		}
		return
	}

	fmt.Printf("Organization: %+v\n", org)
}
