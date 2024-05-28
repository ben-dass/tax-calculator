package filemanager

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm *FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file [%v]", err)
	}

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, fmt.Errorf("failed to read lines in file [%v]", err)
	}

	file.Close()
	return lines, nil
}

func (fm *FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return fmt.Errorf("failed to create file [%v]", err)
	}

	encoder := json.NewEncoder(file)

	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return fmt.Errorf("failed to convert data to JSON [%v]", err)
	}

	file.Close()
	return nil
}

func New(inputPath, outputPath string) *FileManager {
	return &FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
