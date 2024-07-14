package main

import "easysub/internal/cli"

func main() {
	err := cli.Execute()
	if err != nil {
		return
	}
}
