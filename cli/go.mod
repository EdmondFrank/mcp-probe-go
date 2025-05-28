module github.com/edmondfrank/mcp-go-probe

replace github.com/edmondfrank/probe-go-sdk => ../../

replace github.com/edmondfrank/probe-mcp-go => ../

go 1.23.4

require (
	github.com/edmondfrank/probe-mcp-go v0.0.0-00010101000000-000000000000
	github.com/mark3labs/mcp-go v0.30.1
)

require (
	github.com/edmondfrank/probe-go-sdk v0.0.0-00010101000000-000000000000 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/spf13/cast v1.7.1 // indirect
	github.com/yosida95/uritemplate/v3 v3.0.2 // indirect
)
