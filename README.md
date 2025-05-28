# probe-mcp-go

A Go-based MCP (Multi-Tool Code Processing) server for searching and extracting code using ElasticSearch and AST-grep, with support for code block extraction and advanced code queries.

## Features

- **Search code** in your repository using ElasticSearch queries.
- **Query code** using ast-grep structural pattern matching.
- **Extract code blocks** from files by line number or symbol name.
- Supports filtering, context lines, test code inclusion, and output formatting.

## Getting Started

### Prerequisites

- Go 1.23+ installed
- [probe-go-sdk](https://github.com/edmondfrank/probe-go-sdk) and [mcp-go](https://github.com/mark3labs/mcp-go) dependencies

### Build

```bash
cd cli
go build -o probe-mcp-server main.go
```

### Run

```bash
./probe-mcp-server
```

The server will start and listen for stdio requests.

## Usage

The server registers the following tools:

- `search_code`: Search code using ElasticSearch queries.
- `query_code`: Search code using ast-grep patterns.
- `extract_code`: Extract code blocks by line or symbol.

Each tool exposes a set of arguments for flexible code search and extraction. See the code for details on each tool's options.

## Development

- Main server logic is in `main.go` and `cli/main.go`.
- Tool handlers are in `search_code.go`, `query_code.go`, and `extract_code.go`.
- Uses the [mcp-go](https://github.com/mark3labs/mcp-go) server framework.

## License

MIT License. See [LICENSE](LICENSE) for details.
