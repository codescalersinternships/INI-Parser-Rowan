package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Parser Structure: dictionary, sections
// It's basically a map of a map, map of sections where each section maps to keys : values
// map[section] --> returns a map where each key --> value
type Parser struct {
	dictionary map[string]map[string]string
	sections   []string
}

var parsedMap Parser

// LoadFromFile loads ini file
// Saves all lines locally into an array of strings
func LoadFromFile(fileName string) {
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
	parserLogic(iniLines)

}

// LoadFromString loads ini script from a string
// Saves all lines locally into an array of strings
func LoadFromString(str string) {
	iniLines := strings.Split(str, "\n")
	parserLogic(iniLines)
}

func parserLogic(iniLines []string) {
	var sections []string
	parsedMap = Parser{make(map[string]map[string]string), sections}

	var section, key, value string

	for _, line := range iniLines {
		if len(line) == 0 {
			continue
		} else if line[0] == '[' {
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
	}
}

// GetSectionNames provides section names of parsed file/string
func GetSectionNames() []string {
	return parsedMap.sections
}

// GetSections provides the dictionary/map structure
func GetSections() map[string]map[string]string {
	return parsedMap.dictionary
}

// Get function takes 2 parameters: section and its key
// Provides equivalent value
// If section or key aren't found, it returns the zero values
func Get(sectionName, key string) string {
	return parsedMap.dictionary[sectionName][key]
}

// Set Function takes 3 parameters: section, key and value
// If section isn't already present, it makes the map first
func Set(sectionName, key, value string) {
	if parsedMap.dictionary[sectionName] == nil {
		parsedMap.dictionary[sectionName] = make(map[string]string)
	}
	parsedMap.dictionary[sectionName][key] = value
}

// ToString function returns a string structure of the ini file
// Ignores redundant spaces
func ToString() string {
	var stringVersion string
	for _, section := range parsedMap.sections {
		stringVersion += "[" + section + "]"
		for key, value := range parsedMap.dictionary[section] {
			stringVersion += "\n"
			stringVersion += key + " = " + value
		}
		stringVersion += "\n"
	}
	return stringVersion
}

// SaveToFile saves the whole ini map to the given file path
func SaveToFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Cannot open file!")
		os.Exit(1)
	}
	defer file.Close()

	file.WriteString(ToString())
}
