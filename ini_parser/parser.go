package main

import (
	"bufio"
	"errors"
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
}

// NewParser creates an instance of our Parser, which is iniitially an empty map of maps
func NewParser() Parser {
	return Parser{make(map[string]map[string]string)}
}

// var parsedMap Parser

// ErrCouldNotOpen happens when file cannot be opened
var ErrCouldNotOpen = errors.New("cannot open file error") //add filename + change new to fmt.Errorf

// ErrMissingValueAssignment happens when a key isn't followed by an = statement
var ErrMissingValueAssignment = errors.New("key is not assigned to a value, no '=' found")

// ErrSectionNameMissingClosure happens when section name is missing the ] paranthesis
var ErrSectionNameMissingClosure = errors.New("section is missing closure paranthesis ]")

// ErrWrongParanthesisOrder happens when section name starts by thw wrong paranthesis ']'
var ErrWrongParanthesisOrder = errors.New("WrongParanthesisOrder section paranthesis order, section name cannot start by ]")

// ErrInvalidSectionName happens when section is written in a wrong form --> ex: sectionName]
var ErrInvalidSectionName = errors.New("section name can't start with anything other than [")

// LoadFromFile loads ini file
// Saves all lines locally into an array of strings
func (parsedMap *Parser) LoadFromFile(fileName string) error {
	var iniLines []string
	var input io.Reader

	file, err := os.Open(fileName)
	if err != nil {
		return ErrCouldNotOpen
	}
	defer file.Close()

	input = file
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		iniLines = append(iniLines, scanner.Text())
	}
	return parsedMap.parserLogic(iniLines)
}

// LoadFromString loads ini script from a string
// Saves all lines locally into an array of strings
func (parsedMap *Parser) LoadFromString(str string) error {
	iniLines := strings.Split(str, "\n")
	return parsedMap.parserLogic(iniLines)
}

func (parsedMap *Parser) parserLogic(iniLines []string) error {
	// parsedMap.dictionary = make(map[string]map[string]string)

	var section, key, value string

	for _, line := range iniLines {
		if len(line) == 0 {
			continue
		} else if line[0] == '[' {
			closingParaFound := false
			for j, ch := range line {
				if ch == ']' {
					section = line[1:j]
					section = strings.Trim(section, " ")
					closingParaFound = true
				}
				// if closingParaFound && !(line[j] == ' ' || line[j] == '\n' || line[j] == '\t') {
				// 	return ErrTextInSectionLine
				// }
			}
			if !closingParaFound {
				return ErrSectionNameMissingClosure
			}
			if parsedMap.dictionary == nil {
				parsedMap.dictionary = make(map[string]map[string]string)
			}

		} else if line[0] == ']' {
			return ErrWrongParanthesisOrder
		} else if line[0] == ';' || line[0] == ' ' || line[0] == '\n' || line[0] == '\t' {
			continue
		} else {
			equalFound := false
			for j, ch := range line {
				if ch == '=' {
					equalFound = true
					key = line[0:j]
					value = line[j+1:]
					key = strings.Trim(key, " ")
					value = strings.Trim(value, " ")
					if parsedMap.dictionary[section] == nil {
						parsedMap.dictionary[section] = make(map[string]string)
					}
					parsedMap.dictionary[section][key] = value
					break
				} else if ch == ']' {
					return ErrInvalidSectionName
				}
			}
			if !equalFound {
				return ErrMissingValueAssignment
			}
		}
	}
	return nil
}

// GetSectionNames provides section names of parsed file/string
func (parsedMap *Parser) GetSectionNames() []string {
	var sectionNames []string
	for section := range parsedMap.dictionary {
		sectionNames = append(sectionNames, section)
	}
	return sectionNames
}

// GetSections provides the dictionary/map structure
func (parsedMap *Parser) GetSections() map[string]map[string]string {
	return parsedMap.dictionary
}

// Get function takes 2 parameters: section and its key
// Provides equivalent value
// If section or key aren't found, it returns the zero values
func (parsedMap *Parser) Get(sectionName, key string) string {
	return parsedMap.dictionary[sectionName][key]
}

// Set Function takes 3 parameters: section, key and value
// If section isn't already present, it makes the map first
func (parsedMap *Parser) Set(sectionName, key, value string) {
	if parsedMap.dictionary[sectionName] == nil {
		parsedMap.dictionary[sectionName] = make(map[string]string)
	}
	parsedMap.dictionary[sectionName][key] = value
}

// ToString function returns a string structure of the ini file
// Ignores redundant spaces
func (parsedMap *Parser) ToString() string {
	var stringVersion string
	for section := range parsedMap.dictionary {
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
func (parsedMap *Parser) SaveToFile(fileName string) {
	err := os.WriteFile(fileName, []byte(parsedMap.ToString()), 0644)
	if err != nil {
		fmt.Println("Error while writing to file")
		os.Exit(1)
	}
}
