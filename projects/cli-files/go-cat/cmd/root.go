package cmd

import (
	"fmt"
	"os"
)

func Execute() {
	argsLength := len(os.Args)
	if argsLength < 2 {
		return
	}

	for _, arg := range os.Args[1:] {
		catFile(arg)
	}
}

func catFile(path string) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Printf("Error: '%s' no such file or directory", path)
		return
	}

	if fileInfo.IsDir() {
		fmt.Printf("go-cat: %s is a directory", path)
		return
	}

	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Print("Error: Failed to read the file")
		return
	}

	stringContent := string(content)
	fmt.Print(stringContent)
}
