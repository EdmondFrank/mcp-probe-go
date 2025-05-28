package mcpprobego

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/edmondfrank/probe-go-sdk"
)

const SearchCodeToolName = "search_code"

func NewSearchCodeTool() mcp.Tool {
	return mcp.NewTool(SearchCodeToolName,
		mcp.WithDescription("Search code in the repository using ElasticSearch. Use this tool first for any code-related questions."),
		mcp.WithString("path",
			mcp.Required(),
			mcp.Description("Absolute path to the directory to search in (e.g., \"/Users/username/projects/myproject\")."),
		),
		mcp.WithString("query",
			mcp.Required(),
			mcp.Description("Elastic search query. Supports logical operators (AND, OR, NOT), and grouping with parentheses. Examples: \"config\", \"(term1 OR term2) AND term3\". Use quotes for exact matches, like function or type names."),
		),
		mcp.WithBoolean("filesOnly",
			mcp.Description("Skip AST parsing and just output unique files"),
		),
		mcp.WithArray("ignore",
			mcp.Description("Custom patterns to ignore (in addition to .gitignore and common patterns)"),
		),
		mcp.WithBoolean("excludeFilenames",
			mcp.Description("Exclude filenames from being used for matching"),
		),
		mcp.WithBoolean("allowTests",
			mcp.Description("Allow test files and test code blocks in results (disabled by default)"),
		),
		mcp.WithNumber("maxResults",
			mcp.Description("Maximum number of results to return"),
		),
		mcp.WithNumber("maxTokens",
			mcp.Description("Maximum number of tokens to return"),
		),
		mcp.WithString("session",
			mcp.Description("Session identifier for caching. Set to \"new\" if unknown, or want to reset cache. Re-use session ID returned from previous searches"),
		),
	)
}

func SearchCodeHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	path, err := req.RequireString("path")
	if err != nil {
		return mcp.NewToolResultError("Missing required argument: path"), nil
	}
	query, err := req.RequireString("query")
	if err != nil {
		return mcp.NewToolResultError("Missing required argument: query"), nil
	}
	filesOnly := req.GetBool("filesOnly", false)
	ignore := req.GetStringSlice("ignore", nil)
	excludeFilenames := req.GetBool("excludeFilenames", false)
	maxResults := req.GetInt("maxResults", 0)
	maxTokens := req.GetInt("maxTokens", 0)
	allowTests := req.GetBool("allowTests", false)
	session := req.GetString("session", "")

	client := probe.NewProbeClient("")
	opts := probe.SearchOptions{
		Path:             path,
		Query:            query,
		FilesOnly:        filesOnly,
		Ignore:           ignore,
		ExcludeFilenames: excludeFilenames,
		MaxResults:       maxResults,
		MaxTokens:        maxTokens,
		AllowTests:       allowTests,
		Session:          session,
		JSON:             true,
	}
	result, err := client.Search(opts)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Probe search error: %v", err)), nil
	}
	b, err := json.Marshal(result)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal result: %v", err)), nil
	}
	return mcp.NewToolResultText(string(b)), nil
}
