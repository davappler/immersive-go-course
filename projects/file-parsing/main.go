package main

import (
	"file-parsing/abstracts"
	"os"
)

func main() {
	var arg string
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}
	if arg == "json.txt" {
		abstracts.ReadJsonTxt()
	} else if arg == "repeated-json.txt" {
		abstracts.ReadRepeatedJsonTxt()
	} else if arg == "data.csv" {
		abstracts.ReadDataCsv()
	}
}
