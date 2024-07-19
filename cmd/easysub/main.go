package main

import (
	"easysub/internal/cli"
	"fmt"
)

func main() {
	err := cli.Execute()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
