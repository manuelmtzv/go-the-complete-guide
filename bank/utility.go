package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(filename string) (float64, error) {
	data, err := os.ReadFile("balance.txt")
	if err != nil {
		return 0, errors.New("failed to read balance")
	}
	parsedBalance, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return 0, errors.New("failed to parse balance")
	}
	return parsedBalance, nil
}

func WriteFloatToFile(filename string, value float64) error {
	return os.WriteFile(filename, []byte(fmt.Sprintf("%f", value)), 0644)
}
