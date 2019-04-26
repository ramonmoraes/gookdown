package gookdown

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

// Compile is the main function which will compile every file listed in the input file
// and export to the output file
func Compile(inputPath string, outputPath string) {
	summary, err := ioutil.ReadFile(inputPath)
	if err != nil {
		fmt.Println("Could not find summary at", inputPath)
		log.Fatal(err)
	}

	lines := getLinesFromString(string(summary))
	var filePaths []string
	for _, line := range lines {
		path, err := getPathFromReference(line)
		if err != nil {
			fmt.Println("Could not find path for", line)
		}
		if path != "" {
			filePaths = append(filePaths, path)
		}
	}

	compiledContent := getCompiledContent(filePaths)
	err = ioutil.WriteFile(outputPath, compiledContent, 0644)
	if err != nil {
		fmt.Println("Could not write file at", outputPath)
		log.Fatal(err)
	}
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

func getCompiledContent(filePaths []string) []byte {
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

	return content
}
