# MCP Linode Server

A Go-based Management Control Protocol (MCP) server for Linode, designed to help learn both Go programming and MCP server concepts.

## Project Context

This project is a learning exercise to:
1. Understand MCP server concepts
2. Learn Go programming (with comparisons to C#)
3. Implement Linode API integration

## Current State

The project currently has a basic HTTP server setup with:
- Basic server implementation
- Test coverage
- Environment configuration
- Linode API integration (basic instance listing)

## Project Structure

```
mcp-linode/
├── cmd/
│   └── server/
│       └── main.go        # Application entry point
├── internal/
│   └── mcp/
│       ├── server.go      # Core MCP server implementation
│       └── server_test.go # Tests for the MCP server
├── .env.example           # Template for environment variables
└── .gitignore            # Git ignore rules
```

## Setup Instructions

1. **Install Go**
   - Download and install Go from [golang.org](https://golang.org/dl/)
   - Verify installation: `go version`

2. **Clone the Repository**
   ```bash
   git clone https://github.com/yourusername/mcp-linode.git
   cd mcp-linode
   ```

3. **Environment Setup**
   - Copy `.env.example` to `.env`
   - Add your Linode API token to `.env`:
     ```
     LINODE_API_TOKEN=your_token_here
     ```

4. **Install Dependencies**
   ```bash
   go mod download
   ```

## Running the Server

1. **Start the Server**
   ```bash
   go run cmd/server/main.go
   ```

2. **Run Tests**
   ```bash
   go test ./internal/mcp/...
   ```

## Key Concepts Implemented

1. **HTTP Server**
   - Basic routing
   - Request/response handling
   - Graceful shutdown

2. **Testing**
   - Server lifecycle testing
   - HTTP response verification
   - Concurrent testing with goroutines

3. **Environment Management**
   - Secure API token handling
   - Configuration through environment variables

## C# to Go Comparisons

Key differences from C# that are implemented in this project:

1. **Error Handling**
   - Explicit error returns instead of exceptions
   - `if err != nil` pattern
   - No try-catch blocks

2. **Concurrency**
   - Goroutines instead of Tasks
   - Context for cancellation
   - `go` keyword for concurrent execution

3. **Object Orientation**
   - Structs instead of classes
   - Methods defined outside structs
   - Composition over inheritance

## Next Steps

1. Add more MCP server features
2. Implement additional Linode API endpoints
3. Add middleware for logging and authentication
4. Implement proper error handling and recovery

## Development Notes

- The server runs on port 8080 by default
- Tests use port 8081 to avoid conflicts
- Environment variables are loaded from `.env` file
- API tokens should never be committed to version control
