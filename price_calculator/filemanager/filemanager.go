package filemanager

import (
	"bufio"
	"errors"
	"os"
)

func ReadLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, errors.New("error opening file")
	}

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		return nil, errors.New("error reading file")
	}

	defer file.Close()
	return lines, nil
}
