package main

import (
	"fmt"
	"os"

	"github.com/edmondfrank/probe-mcp-go"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	version := "1.0.0"
	if v := os.Getenv("MCP_SERVER_VERSION"); v != "" {
		version = v
	}

	// Create a new MCP server with all tools registered
	s := mcpprobego.NewProbeMCPServer(
		"@buger/probe-mcp-go",
		version,
	)

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		os.Exit(1)
	}
}
