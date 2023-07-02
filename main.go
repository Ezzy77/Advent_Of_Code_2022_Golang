package main

import (
	"fmt"
	"os"
)

func main() {

	path := "sample.txt"
	content := readFile(path)

	fmt.Println(content)
}

func readFile(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("error reading file")
	}

	var content []string

	for _, v := range data {
		content = append(content, string(v))
	}

	return content
}
