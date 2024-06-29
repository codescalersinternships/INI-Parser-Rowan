package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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

// ErrCouldNotOpen happens when file cannot be opened and provides file name
var ErrCouldNotOpen error

// ErrMissingValueAssignment happens when a key isn't followed by an = statement
var ErrMissingValueAssignment error

// ErrSectionNameMissingClosure happens when section name is missing the ] paranthesis
var ErrSectionNameMissingClosure error

// ErrWrongParanthesisOrder happens when section name starts by the wrong paranthesis ']'
var ErrWrongParanthesisOrder error

// ErrInvalidSectionName happens when section is written in a wrong form --> ex: sectionName]
var ErrInvalidSectionName error

// ErrCouldNotWriteToFile happens when file cannot be written to
var ErrCouldNotWriteToFile error

// LoadFromFile loads ini file
// Saves all lines locally into an array of strings
func (parsedMap *Parser) LoadFromFile(fileName string) error {
	var iniLines []string
	var input io.Reader

	file, err := os.Open(fileName)
	if err != nil {
		ErrCouldNotOpen = fmt.Errorf("cannot open file: %s", fileName)
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
					if parsedMap.dictionary[section] == nil {
						parsedMap.dictionary[section] = make(map[string]string)
					}
				}
			}
			if !closingParaFound {
				ErrSectionNameMissingClosure = fmt.Errorf("section [%s] is missing closure paranthesis ']'", line[1:])
				return ErrSectionNameMissingClosure
			}
			if parsedMap.dictionary == nil {
				parsedMap.dictionary = make(map[string]map[string]string)
			}

		} else if line[0] == ']' {
			ErrWrongParanthesisOrder = fmt.Errorf("wrong section paranthesis order, the section provided, %s, cannot start by ]", line)
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
					parsedMap.dictionary[section][key] = value
					break
				} else if ch == ']' {
					ErrInvalidSectionName = fmt.Errorf("section name provided, %s, can't start with anything other than '['", line[:j+1])
					return ErrInvalidSectionName
				}
			}
			if !equalFound {
				ErrMissingValueAssignment = fmt.Errorf("key provided, %s , is not assigned to a value, no '=' found", line)
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
	var sortedSections []string
	for section := range parsedMap.dictionary {
		sortedSections = append(sortedSections, section)
	}
	sort.Strings(sortedSections)
	for _, section := range sortedSections {
		stringVersion += "[" + section + "]"

		var keys []string
		for key := range parsedMap.dictionary[section] {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			stringVersion += "\n"
			stringVersion += key + " = " + parsedMap.dictionary[section][key]
		}
		stringVersion += "\n"
	}
	stringVersion = strings.Trim(stringVersion, "\n")
	return stringVersion
}

// SaveToFile saves the whole ini map to the given file path
func (parsedMap *Parser) SaveToFile(fileName string) error {
	err := os.WriteFile(fileName, []byte(parsedMap.ToString()), 0644)
	if err != nil {
		ErrCouldNotWriteToFile = fmt.Errorf("cannot write/save to file: %s", fileName)
		return ErrCouldNotWriteToFile
	}
	return nil
}
