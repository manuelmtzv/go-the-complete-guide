package utility

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Saver interface {
	Save() error
}

type Outputtable interface {
	Saver
	Display() string
}

func SaveToJson(fileName string, data any) error {
	json, err := json.Marshal(data)

	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, json, 0644)

	if err != nil {
		return err
	}

	return nil
}

func ComposeFileName(title string, extension string) string {
	fileName := strings.ReplaceAll(title, " ", "_") + extension
	fileName = strings.ToLower(fileName)
	return fileName
}

func SaveData(source Saver) error {
	err := source.Save()

	if err != nil {
		fmt.Println("An error occurred while saving data.")
		return err
	}

	fmt.Println("Data saved successfully!")
	return nil
}

func OutputData(source Outputtable) error {
	fmt.Println(source.Display())
	err := source.Save()

	if err != nil {
		fmt.Println("An error occurred while saving data.")
		return err
	}

	return nil
}
