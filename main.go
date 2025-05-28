package main

import (
	"context"
	"fmt"

	"github.com/0xUmang/dictionaryapi/cmd"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Errorf("Error loading .env file. Please check if .env file is present with the API key")
	}

	ctx := context.Background()
	cmd.Execute(ctx)
}
