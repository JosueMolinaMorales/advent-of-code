package io

import (
	"os"
	"strings"
)

func ReadFileAsLines(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), "\n"), nil
}

func ReadFileAsString(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func ReadFileAs2DArray(path string) ([][]string, error) {
	data, err := ReadFileAsLines(path)
	if err != nil {
		return nil, err
	}

	arr := make([][]string, len(data))
	for i, line := range data {
		row := make([]string, len(line))
		for j, ch := range line {
			row[j] = string(ch)
		}
		arr[i] = row
	}

	return arr, nil
}
