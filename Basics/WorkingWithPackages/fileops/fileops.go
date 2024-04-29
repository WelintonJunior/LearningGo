package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	// Nil quer dizer que não tem erro, se for diferente de nil tem erro
	if err != nil {
		return 0, errors.New("Failed to find file")
	}

	dataText := string(data)
	value, err := strconv.ParseFloat(dataText, 64)

	if err != nil {
		return 0, errors.New("Failed to parse stored value")
	}

	return value, nil
}

func WriteFloatToFile(data float64, fileName string) {
	dataText := fmt.Sprint(data)
	os.WriteFile(fileName, []byte(dataText), 0644)
}
