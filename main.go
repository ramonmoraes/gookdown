package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func main() {
	summary := `
		# 1. Introdução
		- [Introdução](./files/foo.md)
		# 2. Objetos de Estudo
		- [grafos](./files/foo.md)
	`
	lines := getLinesFromString(summary)
	var paths []string
	for _, line := range lines {
		path, err := getPathFromReference(line)
		if err != nil {
			fmt.Println("Could not find path for", line)
		}
		if path != "" {
			paths = append(paths, path)
		}
	}
	compilePaths(paths)
}

func getLinesFromString(input string) []string {
	lines := strings.Split(input, "\n")
	var trimmedLines []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		trimmedLines = append(trimmedLines, line)
	}
	return trimmedLines
}

func getPathFromReference(line string) (string, error) {
	re, _ := regexp.Compile(".*\\[(.*)\\]\\((.*)\\).*")
	filePathGroup := 2

	find := re.Find([]byte(line))
	if len(find) < filePathGroup {
		return "", errors.New("Could not find submatch")
	}

	submatchs := re.FindSubmatch([]byte(line))
	return string(submatchs[filePathGroup]), nil
}

func compilePaths(filePaths []string) {
	var content []byte
	contentSeparator := []byte("\n\n")

	for i, path := range filePaths {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal("File in path not fount", path)
		}

		if i != 0 {
			content = append(content, contentSeparator...)
		}

		content = append(content, data...)
	}

	outputPath := "./output.md"
	err := ioutil.WriteFile(outputPath, content, 0644)
	if err != nil {
		log.Fatal("Could not write file at", outputPath)
	}
}
