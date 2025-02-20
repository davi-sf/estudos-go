package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func New(inputFilePath, outputFilePath string) FileManager {
	return FileManager{
		InputFilePath:  inputFilePath,
		OutputFilePath: outputFilePath,
	}
}

func (fm FileManager) ReadLines() ([]string, error) {

	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, errors.New("Failed to open the file")
	}

	var lines []string

	sc := bufio.NewScanner(file)
	sc.Scan()

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	err = sc.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("Failed to read line in the file")
	}

	file.Close()
	return lines, nil

}

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("Error creating file")
	}

	enconder := json.NewEncoder(file)
	err = enconder.Encode(data)

	if err != nil {
		file.Close()
		errors.New("failed to convert data to JSON")
	}

	file.Close()
	return nil

}
