package main

import "fmt"

func main() {
	LoadFromFile("ini_parser/file.ini")
	fmt.Println(GetSectionNames())
	fmt.Println(GetSections())
	fmt.Println(Get("rowan", "hi"))
	fmt.Println(Get("hello", "rowan"))
	fmt.Println(Get("hello", "mohamed"))
	Set("hello", "uni", "corn")
	Set("new", "uni", "corn")
	// fmt.Println(parsedMap.dictionary)
	fmt.Println(ToString())
	LoadFromString("[hello]\nrowan = just a girl\ntest = ini\n[works]\nluna = cat")
}
