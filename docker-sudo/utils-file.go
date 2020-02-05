package main

import (
	"bufio"
	"io/ioutil"
	"os"
)

func getListInFile(path string) ([]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func getListInFolder(path string) ([]string, error) {

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var lines []string
	for _, file := range files {
		if file.IsDir() {
			lines = append(lines, file.Name())
		}
	}
	return lines, nil
}
