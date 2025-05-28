package mcpprobego

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/edmondfrank/probe-go-sdk"
)
const QueryCodeToolName = "query_code"

func NewQueryCodeTool() mcp.Tool {
	return mcp.NewTool(QueryCodeToolName,
		mcp.WithDescription("Search code using ast-grep structural pattern matching. Use this tool to find specific code structures like functions, classes, or methods."),
		mcp.WithString("path",
			mcp.Required(),
			mcp.Description("Absolute path to the directory to search in (e.g., \"/Users/username/projects/myproject\")."),
		),
		mcp.WithString("pattern",
			mcp.Required(),
			mcp.Description("The ast-grep pattern to search for. Examples: \"fn $NAME($$$PARAMS) $$$BODY\" for Rust functions, \"def $NAME($$$PARAMS): $$$BODY\" for Python functions."),
		),
		mcp.WithString("language",
			mcp.Description("The programming language to search in. If not specified, the tool will try to infer the language from file extensions. Supported languages: rust, javascript, typescript, python, go, c, cpp, java, ruby, php, swift, csharp."),
		),
		mcp.WithString("ignore",
			mcp.Description("Custom patterns to ignore (in addition to common patterns)"),
		),
		mcp.WithBoolean("allowTests",
			mcp.Description("Allow test files and test code blocks in results (disabled by default)"),
		),
		mcp.WithNumber("maxResults",
			mcp.Description("Maximum number of results to return"),
		),
		mcp.WithString("format",
			mcp.Description("Output format for the query results"),
		),
	)
}

func QueryCodeHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	path, err := req.RequireString("path")
	if err != nil {
		return mcp.NewToolResultError("Missing required argument: path"), nil
	}
	pattern, err := req.RequireString("pattern")
	if err != nil {
		return mcp.NewToolResultError("Missing required argument: pattern"), nil
	}
	language := req.GetString("language", "")
	ignore := req.GetStringSlice("ignore", nil)
	allowTests := req.GetBool("allowTests", false)
	maxResults := req.GetInt("maxResults", 0)
	format := req.GetString("format", "")

	client := probe.NewProbeClient("")
	opts := probe.QueryOptions{
		Path:       path,
		Pattern:    pattern,
		Language:   language,
		Ignore:     ignore,
		AllowTests: allowTests,
		MaxResults: maxResults,
		Format:     format,
		JSON:       true,
	}
	result, err := client.Query(opts)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Probe query error: %v", err)), nil
	}
	b, err := json.Marshal(result)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal result: %v", err)), nil
	}
	return mcp.NewToolResultText(string(b)), nil
}
