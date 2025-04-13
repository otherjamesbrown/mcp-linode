package mcp

import (
	"context"
	"net/http"
	"testing"
	"time"
)

// TestServerStartStop tests the basic functionality of starting and stopping the server
// Similar to a test method in C# using NUnit or xUnit
func TestServerStartStop(t *testing.T) {
	// Create a new server instance
	// Similar to creating a new instance of a service in C# tests
	server := NewServer(":8081")

	// Create a context with a 5-second timeout
	// Similar to using CancellationTokenSource in C#
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// Ensure we cancel the context when the test is done
	// Similar to using 'using' statement in C# for cleanup
	defer cancel()

	// Start the server in a goroutine (similar to Task.Run() in C#)
	// This allows the test to continue while the server runs
	go func() {
		// Start the server and check for errors
		// Similar to await service.Start() in C#
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			t.Errorf("Server failed to start: %v", err)
		}
	}()

	// Give the server a moment to start up
	// Similar to Task.Delay() in C#
	time.Sleep(100 * time.Millisecond)

	// Test if the server is responding to HTTP requests
	// Similar to using HttpClient in C# tests
	resp, err := http.Get("http://localhost:8081")
	if err != nil {
		t.Errorf("Failed to connect to server: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	// Stop the server gracefully
	// Similar to await service.Stop() in C#
	if err := server.Stop(ctx); err != nil {
		t.Errorf("Failed to stop server: %v", err)
	}
}
