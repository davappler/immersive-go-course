// cmd/root.go
package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func Execute() {
	argsLength := len(os.Args)
	fmt.Print(os.Args)
	if argsLength > 2 {
		fmt.Println("Error: Too many arguments, please enter only one file name")
		return
	}
	path := ""
	// When we run code with go-ls . => it will have two args [go-ls .]
	if argsLength == 2 {
		path = os.Args[1]
	}

	if path == "" {
		path = "."
	}

	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error occured", err)
		return
	}
	for _, dir := range files {
		c := color.New(color.Reset)
		if dir.IsDir() {
			c = color.New(color.FgCyan).Add(color.Bold)
		}
		c.Printf("  %s", dir.Name())
	}
}
