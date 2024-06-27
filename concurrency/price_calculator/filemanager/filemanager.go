package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func New(inputPath, outputPath string) (*FileManager, error) {
	if inputPath == "" || outputPath == "" {
		return nil, errors.New("input and output paths should not be empty")
	}

	return &FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}, nil
}

func (filemanager *FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(filemanager.InputFilePath)

	if err != nil {
		return nil, errors.New("error opening file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		return nil, errors.New("error reading file")
	}

	return lines, nil
}

func (filemanager *FileManager) WriteJson(data any) error {
	file, err := os.Create(filemanager.OutputFilePath)

	if err != nil {
		return errors.New("error creating file")
	}

	defer file.Close()

	time.Sleep(3 * time.Second)

	jsonEncoder := json.NewEncoder(file)
	err = jsonEncoder.Encode(data)

	if err != nil {
		return errors.New("error writing to file")
	}

	return nil
}
