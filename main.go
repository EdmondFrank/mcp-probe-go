package mcpprobego

import (
	"github.com/mark3labs/mcp-go/server"
)

// NewProbeMCPServer returns a configured MCP server instance.
// The caller is responsible for starting the server (e.g., with server.ServeStdio).
func NewProbeMCPServer(name, version string) *server.MCPServer {
	s := server.NewMCPServer(
		name,
		version,
		server.WithToolCapabilities(false),
	)
	s.AddTool(NewSearchCodeTool(), SearchCodeHandler)
	s.AddTool(NewQueryCodeTool(), QueryCodeHandler)
	s.AddTool(NewExtractCodeTool(), ExtractCodeHandler)
	return s
}
