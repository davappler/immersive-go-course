package abstracts

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Player struct {
	Name      string `json:"name"`
	HighScore int64  `json:"high_score"`
}

func ReadJsonTxt() {
	jsonFile, err := os.Open("examples/json.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Success: File opened")
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	// fmt.Printf("Type: %T", byteValue) => []uint8

	var players []Player
	json.Unmarshal(byteValue, &players)

	for i := 0; i < len(players); i++ {
		fmt.Printf("Name %s, Score %d \n", players[i].Name, players[i].HighScore)
	}
}

func ReadRepeatedJsonTxt() {
	// Read the JSON file
	data, err := os.ReadFile("examples/repeated-json.txt")
	if err != nil {
		panic(err)
	}

	// Filter out comment lines
	var jsonData []string
	for _, line := range strings.Split(string(data), "\n") {
		if !strings.HasPrefix(line, "#") {
			jsonData = append(jsonData, line)
		}
	}
	// Wrap the data in an array without trailing comma
	var builder strings.Builder
	builder.WriteRune('[')
	for i, line := range jsonData {
		builder.WriteString(line)
		if i < len(jsonData)-2 { // Only add comma if not the last element
			builder.WriteRune(',')
		}
	}
	builder.WriteRune(']')
	wrappedData := builder.String()

	var players []Player
	json.Unmarshal([]byte(wrappedData), &players)

	for i := 0; i < len(players); i++ {
		fmt.Printf("Name %s, Score %d \n", players[i].Name, players[i].HighScore)
	}

}
