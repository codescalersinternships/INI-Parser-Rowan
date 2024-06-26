package main

import "fmt"

func main() {
	parserLogic(loadFromFile("cmd/ini_parser/file.ini"))
	fmt.Println(getSectionNames())
	fmt.Println(getSections())
	fmt.Println(get("rowan", "hi"))
	fmt.Println(get("hello", "rowan"))
	fmt.Println(get("hello", "mohamed"))
	set("hello", "uni", "corn")
	set("new", "uni", "corn")
	fmt.Println(parsedMap.dictionary)
	fmt.Println(toString())
	parserLogic(loadFromString("[hello]\nrowan = just a girl\ntest = ini\n[works]\nluna = cat"))
}
