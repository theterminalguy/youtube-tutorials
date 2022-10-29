# codegen-example

This repo demonstrates how to implement a simple code generator with Go. 

### Usage

1. Clone or download this subdirectory
2. Install the necessary dependencies which is only "github.com/google/uuid" with `go mod download`
3. Run `go run main.go <entity_name>` to generate the code for the entity
4. Check `internal/repository/` for the generated code