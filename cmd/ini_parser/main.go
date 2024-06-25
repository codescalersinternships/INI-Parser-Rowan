package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")
	loadFromFile("cmd/ini_parser/file.ini")
	loadFromString("rowan \n says \n hi \n nom")
}
