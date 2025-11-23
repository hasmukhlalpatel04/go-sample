package main

import (
	"bufio"
	"fmt"
	"os"
)

// FileProcessor encapsulates file operations
type FileProcessor struct {
	FilePath string
}

func (fp *FileProcessor) ReadAll() (string, error) {
	data, err := os.ReadFile(fp.FilePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Predicate is a function type that decides whether a line should be kept
type Predicate func(string) bool

// FindAll reads the file line by line and returns all lines
// that satisfy the predicate.
func (fp *FileProcessor) FindAll(predicate Predicate) ([]string, error) {
	file, err := os.Open(fp.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var results []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if predicate(line) {
			results = append(results, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

// ReadFirst reads the file line by line and returns the first line
// that satisfies the predicate.
func (fp *FileProcessor) ReadFirst(predicate Predicate) (string, error) {
	file, err := os.Open(fp.FilePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if predicate(line) {
			return line, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", fmt.Errorf("no matching line found")
}

// AppendToFile appends given lines to the file.
func (fp *FileProcessor) AppendToFile(lines []string) error {
	file, err := os.OpenFile(fp.FilePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range lines {
		if _, err := file.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	return nil
}
