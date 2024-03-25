package abstracts

import (
	"encoding/csv"
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

func ReadDataCsv() {
	file, err := os.Open("examples/data.csv")
	if err != nil {
		fmt.Print("Error while opening the file")
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Print("Failed to read the records")
	}

	for i := 0; i < len(records); i++ {
		if i < 1 {
			continue
		}
		currentRecord := records[i]

		fmt.Println(currentRecord[0], currentRecord[1])
	}
}
