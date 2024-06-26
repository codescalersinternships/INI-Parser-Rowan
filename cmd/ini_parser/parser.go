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

var parsedMap Parser

func loadFromFile(fileName string) []string {
	var iniLines []string
	var input io.Reader

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
	return iniLines

}

func loadFromString(str string) []string {
	iniLines := strings.Split(str, "\n")
	return iniLines
}

func parserLogic(iniLines []string) {
	var sections []string
	parsedMap = Parser{make(map[string]map[string]string), sections}

	var section, key, value string

	for _, line := range iniLines {
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
		// fmt.Println(parsedMap.sections)
		// fmt.Println(parsedMap.dictionary)
	}
	fmt.Println("Sections are: ", parsedMap.sections)
	fmt.Println("Dictionary is: ", parsedMap.dictionary)
}

// func printParsedMap(){
// 	fmt.Println(parsedMap)
// }

func getSectionNames() []string {
	return parsedMap.sections
}

func getSections() map[string]map[string]string {
	return parsedMap.dictionary
}

func get(sectionName, key string) string {
	return parsedMap.dictionary[sectionName][key]
}

func set(sectionName, key, value string) {
	// found := false
	if parsedMap.dictionary[sectionName] == nil {
		// for _, section := range parsedMap.sections {
		// 	if section == sectionName {
		// 		found = true
		// 		break
		// 	}
		// }
		// if found {

		// }
		parsedMap.dictionary[sectionName] = make(map[string]string)
	}
	parsedMap.dictionary[sectionName][key] = value
}
func toString() string {
	var stringVersion string
	for _, section := range parsedMap.sections {
		stringVersion += "[" + section + "]"
		for key, value := range parsedMap.dictionary[section] {
			stringVersion += "\n"
			stringVersion += key + " = " + value
		}
	}
	return stringVersion
}
