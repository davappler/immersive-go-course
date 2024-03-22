// cmd/root.go
package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func Execute() {
	files, err := os.ReadDir(".")
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
