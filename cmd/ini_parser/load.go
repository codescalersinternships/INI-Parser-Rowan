package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var ini []string

func loadFromFile(fileName string) {
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
		ini = append(ini, scanner.Text())
	}
	for _, line := range ini {
		fmt.Println(line)
	}

}

func loadFromString(str string) {
	ini = strings.Split(str, "\n")
	for _, line := range ini {
		fmt.Println(line)
	}
}
