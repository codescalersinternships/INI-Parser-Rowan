# INI-Parser-Rowan
An API that parses INI files and provides some methods to get info about this file

## Features
- `LoadFromString`
- `LoadFromFile`
- `GetSectionNames` list of all section names
- `GetSections` serialize convert into a dictionary/map ` { section_name: {key1: val1, key2, val2} ...}`
- `Get(section_name, key)` gets the value of key `key` in section `section_name`
- `Set(section_name, key, value)` sets a `key` in section `section_name` to value `value`
- `ToString`
- `SaveToFile`

## Module Files
- `parser.go`: has all API functions implementation
- `parser_test.go`: includes all tests for each functionality of the API to ensure 
- temp.ini: temporary file for the `LoadFromFile` and `SaveToFile` tests

## Step-By-Step Guide
To use the library/API
1. To create an an object of type Parser:
    ```
    parser := NewParser()
    ```
2. Load your ini format using either file or string, this returns an error which user can choose to inspect.
    ```
    err := parser.LoadFromString(string)
    ```
    ```
    err := parser.LoadFromString(ilePath)
    ```
3. Perform the functionalities you want
    ```
    valueReceived := parser.Get(sectionName, keyName)
    ```
    ```
    sections := parser.GetSectionNames() //returns a slice of section names
    ```
    ```
    fullMap := parser.GetSections()
    ```
    ```
    parser.Set(sectionName, keyName, ValueToSet)
    ```
    ```
    stringToPrint := parser.ToString()
    ```
    ```
    err := parser.SaveToFile(filePath)
    ```
## Errors To Inspect
- `ErrCouldNotOpen` happens when file cannot be opened and provides file name


- `ErrMissingValueAssignment` happens when a key isn't followed by an = statement


- `ErrSectionNameMissingClosure` happens when section name is missing the ']' paranthesis


- `ErrWrongParanthesisOrder` happens when section name starts by the wrong paranthesis ']'


- `ErrInvalidSectionName` happens when section is written in a wrong form --> ex: sectionName]


- `ErrCouldNotWriteToFile` happens when file cannot be written to
