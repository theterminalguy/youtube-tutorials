package main

import (
	"fmt"
	"os"

	"example.com/scaffold/internal/repository"
)

func main() {
	// get the first argument
	name := os.Args[1]
	if err := repository.GenerateBoilerplate(name); err != nil {
		panic(err)
	}
	fmt.Printf("[created] %v\n", fmt.Sprintf("internal/repository/%v_repository.go", name))
}
