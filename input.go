package main

import "os"

func fileReader(fileName string) ([]byte, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return content, nil
}
