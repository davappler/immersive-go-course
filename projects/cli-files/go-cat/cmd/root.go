package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Execute() {
	argsLength := len(os.Args)
	if argsLength < 2 {
		return
	}

	path := os.Args[1]
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Printf("Error: '%s' no such file or directory", path)
		return
	}

	if fileInfo.IsDir() {
		fmt.Printf("go-cat: %s is a directory", path)
		return
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print("Error: Failed to read the file")
		return
	}

	stringContent := string(content)
	fmt.Print(stringContent)

}

