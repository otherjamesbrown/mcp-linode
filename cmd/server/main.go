package main

import (
	"fmt"
	"log"
	"net/http"
)

// main is the entry point of our application, similar to Program.Main() in C#
func main() {
	// In Go, we use HandleFunc to register a route handler, similar to ASP.NET's routing
	// The first parameter is the path pattern ("/" means root path)
	// The second parameter is a function that handles the request
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// http.ResponseWriter is similar to HttpResponse in C#
		// http.Request is similar to HttpRequest in C#

		// Write a response to the client
		// Similar to Response.Write() in C#
		fmt.Fprintf(w, "MCP Server is running!")
	})

	// Define the port the server will listen on
	// Similar to setting up a port in C# Web.config or appsettings.json
	port := ":8080"

	// Print a startup message to the console
	// Similar to Console.WriteLine() in C#
	fmt.Printf("Starting MCP server on port %s\n", port)

	// Start the HTTP server and listen for incoming requests
	// Similar to WebHost.Run() in ASP.NET Core
	// If there's an error starting the server, log it and exit
	// log.Fatalf is similar to throwing an exception in C#
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
