package mcp

import (
	"context"
	"net/http"
)

// Server represents our MCP server
// Similar to a class in C#, but in Go we use structs instead
// This struct encapsulates all the server-related functionality
type Server struct {
	// httpServer is a pointer to the underlying HTTP server
	// Similar to a private field in a C# class
	httpServer *http.Server
}

// NewServer is a constructor function that creates a new Server instance
// Similar to a constructor in C#, but in Go we use a function that returns a pointer
// The 'addr' parameter is the address the server will listen on (e.g., ":8080")
func NewServer(addr string) *Server {
	// Create and return a new Server instance
	// Similar to 'return new Server()' in C#
	return &Server{
		// Initialize the http.Server with the provided address
		// Similar to setting properties in a C# constructor
		httpServer: &http.Server{
			Addr: addr,
		},
	}
}

// Start begins listening for HTTP requests
// Similar to a Start() method in a C# service class
// Returns an error if the server fails to start
func (s *Server) Start() error {
	// ListenAndServe starts the HTTP server and blocks until the server is closed
	// Similar to WebHost.Run() in ASP.NET Core
	return s.httpServer.ListenAndServe()
}

// Stop gracefully shuts down the server
// Similar to a Stop() method in a C# service class
// Takes a context for timeout and cancellation control
// Similar to CancellationToken in C#
func (s *Server) Stop(ctx context.Context) error {
	// Shutdown gracefully stops the server without interrupting any active connections
	// Similar to IHost.StopAsync() in ASP.NET Core
	return s.httpServer.Shutdown(ctx)
}
