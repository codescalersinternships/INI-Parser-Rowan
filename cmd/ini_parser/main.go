package main

import (
// "fmt"
// "strings"
// "regexp"
)

func main() {
	// var m map[int]int
	// m = make(map[int]int)
	// m[1] = 3
	// str := "  Hello World  \n"
	// trimmed := strings.Trim(str, " ")
	// trimmed = trimmed[:len(trimmed) -1]
	// trimmed = strings.Replace(trimmed,"[","", 1)
	// trimmed = strings.Replace(trimmed,"]","", 1)
	// fmt.Print(trimmed)
	// fmt.Println("hello")
	parserLogic(loadFromFile("cmd/ini_parser/file.ini"))
	parserLogic(loadFromString("[hello]\nrowan = just a girl\ntest = ini\n[works]\nluna = cat"))
	// loadFromString("rowan \nsays \nhi \nnom")
}
