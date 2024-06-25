package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Parser Structure: dictionary, sections
// It's basically a map of a map
// map[section] --> returns a map where each key --> value
type Parser struct {
	dictionary map[string]map[string]string
	sections   []string
}

func loadFromFile(fileName string) []string {
	var iniLines []string
	var input io.Reader
	fmt.Println("hi")

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Cannot open file!")
		os.Exit(1)
	}
	defer file.Close()

	input = file
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		iniLines = append(iniLines, scanner.Text())
	}
	// for _, line := range iniLines {
	// 	fmt.Println(line)
	// }
	return iniLines

}

func loadFromString(str string) []string {
	iniLines := strings.Split(str, "\n")
	// for _, line := range iniLines {
	// 	fmt.Println(line)
	// }
	return iniLines
}

func parserLogic(iniLines []string) {
	var sections []string
	parsedMap := Parser{make(map[string]map[string]string), sections}

	var section, key, value string

	for _, line := range iniLines {
		// fmt.Println(line)
		if line[0] == '[' {
			for j, ch := range line {
				if ch == ']' {
					section = line[1:j]
					parsedMap.sections = append(parsedMap.sections, section)
				}
			}
			if parsedMap.dictionary == nil {
				parsedMap.dictionary = make(map[string]map[string]string)
			}

		} else if line[0] == ';' || line[0] == ' ' || line[0] == '\n' || line[0] == '\t' {
			continue
		} else {
			for j, ch := range line {
				if ch == '=' {
					key = line[0:j]
					value = line[j+1:]
					key = strings.Trim(key, " ")
					value = strings.Trim(value, " ")
					if parsedMap.dictionary[section] == nil {
						parsedMap.dictionary[section] = make(map[string]string)
						parsedMap.dictionary[section][key] = value
					}
					break
				}
			}
		}
		fmt.Println(parsedMap.sections)
		fmt.Println(parsedMap.dictionary)
	}
}
