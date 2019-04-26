package gookdown

import (
	"testing"
)

func TestGetLinesFromString(t *testing.T) {
	input := `two
	lines`
	lines := getLinesFromString(input)
	if len(lines) != 2 {
		t.Error("Input should be breaked in two lines")
	}
	if lines[1] != "lines" {
		t.Error("Second line should be trimmed")
	}
}

func TestGetPathFromReference(t *testing.T) {
	input := `- [Intro](./readme/intro.md) ,`
	expected, err := getPathFromReference(input)
	if err != nil {
		t.Error("Should get paths without errors")
	}
	if expected != "./readme/intro.md" {
		t.Error("Expected path is wrong")
	}
}
