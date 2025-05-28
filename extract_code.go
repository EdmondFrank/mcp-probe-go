package mcpprobego

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/edmondfrank/probe-go-sdk"
)

const ExtractCodeToolName = "extract_code"

func NewExtractCodeTool() mcp.Tool {
	return mcp.NewTool(ExtractCodeToolName,
		mcp.WithDescription("Extract code blocks from files based on line number, or symbol name. Fetch full file when line number is not provided."),
		mcp.WithArray("files",
			mcp.Required(),
			mcp.Description("Files and lines or symbols to extract from: /path/to/file.rs:10, /path/to/file.rs#func_name Path should be absolute."),
		),
		mcp.WithBoolean("allowTests",
			mcp.Description("Allow test files and test code blocks in results (disabled by default)"),
		),
		mcp.WithNumber("contextLines",
			mcp.Description("Number of context lines to include before and after the extracted block when AST parsing fails to find a suitable node"),
		),
		mcp.WithString("format",
			mcp.Description("Output format for the extracted code"),
		),
	)
}

func ExtractCodeHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	files, err := req.RequireStringSlice("files")
	if err != nil {
		return mcp.NewToolResultError("Missing required argument: files"), nil
	}
	allowTests := req.GetBool("allowTests", false)
	contextLines := req.GetInt("contextLines", 0)
	format := req.GetString("format", "plain")

	client := probe.NewProbeClient("")

	// If files is empty or contains a single empty string, treat as no files
	var fileList []string
	if len(files) > 0 && !(len(files) == 1 && files[0] == "") {
		fileList = files
	}

	opts := probe.ExtractOptions{
		Files:        fileList,
		AllowTests:   allowTests,
		ContextLines: contextLines,
		Format:       format,
		JSON:         true,
	}

	result, err := client.Extract(opts)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Probe extract error: %v", err)), nil
	}
	b, err := json.Marshal(result)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal result: %v", err)), nil
	}
	return mcp.NewToolResultText(string(b)), nil
}
