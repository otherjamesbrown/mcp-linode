package linode

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestListInstances(t *testing.T) {
	// Load the .env file
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	// Get the Linode API token from environment variable
	apiToken := os.Getenv("LINODE_API_TOKEN")
	if apiToken == "" {
		t.Fatal("LINODE_API_TOKEN environment variable is not set")
	}

	// Create a new client
	client := NewClient(apiToken)

	// Create a context
	ctx := context.Background()

	// List instances
	instances, err := client.ListInstances(ctx)
	if err != nil {
		t.Fatalf("Failed to list instances: %v", err)
	}

	// Log the number of instances found
	t.Logf("Found %d Linode instances", len(instances))

	// Basic validation
	if len(instances) == 0 {
		t.Log("No instances found - this might be expected if you don't have any Linodes")
	}

	// If we have instances, log some basic information about them
	for _, instance := range instances {
		t.Logf("Instance: %s (ID: %d, Status: %s)", instance.Label, instance.ID, instance.Status)
	}
}
