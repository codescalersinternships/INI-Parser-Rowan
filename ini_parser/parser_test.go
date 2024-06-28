package main

import (
	// "go/parser"
	"reflect"
	"slices"

	// "strings"
	"testing"
)

func Test_LoadFromFile(t *testing.T) {

}
func Test_LoadFromString(t *testing.T) {
	parser := NewParser()
	t.Run("Nil Error returned", func(t *testing.T) {
		stringToTest := "[people]\nrowan = just a girl\nbob ross = bad mentor\n[entity]\ncodeScalers = company\n[location]\nalex = city"
		err := parser.LoadFromString(stringToTest)
		if err != nil {
			t.Errorf("Got %v, wanted %v", err, nil)
		}
	})
	t.Run("Missing value assignment", func(t *testing.T) {
		stringToTest := "[people]\nrowan = just a girl\nbob ross = bad mentor\n[entity]\ncodeScalers \n[location]\nalex = city"
		err := parser.LoadFromString(stringToTest)
		if err != ErrMissingValueAssignment {
			t.Errorf("Got %v, wanted %v", err, nil)
		}
	})
	t.Run("Section name missing closure parenthesis", func(t *testing.T) {
		stringToTest := "[people\nrowan = just a girl\nbob ross = bad mentor\n[entity]\ncodeScalers = company\n[location]\nalex = city"
		err := parser.LoadFromString(stringToTest)
		if err != ErrSectionNameMissingClosure {
			t.Errorf("Got %v, wanted %v", err, nil)
		}
	})
	t.Run("Wrong Paranthesis Order", func(t *testing.T) {
		stringToTest := "]people[\nrowan = just a girl\nbob ross = bad mentor\n[entity]\ncodeScalers \n[location]\nalex = city"
		err := parser.LoadFromString(stringToTest)
		if err != ErrWrongParanthesisOrder {
			t.Errorf("Got %v, wanted %v", err, nil)
		}
	})
	t.Run("Invalid Section starting", func(t *testing.T) {
		stringToTest := "people]\nrowan = just a girl\nbob ross = bad mentor\n[entity]\ncodeScalers = company\n[location]\nalex = city"
		err := parser.LoadFromString(stringToTest)
		if err != ErrInvalidSectionName {
			t.Errorf("Got %v, wanted %v", err, nil)
		}
	})
}

func Test_GetSectionNames(t *testing.T) {
	t.Run("Several Sections", func(t *testing.T) {
		parser := NewParser()
		stringToTest := "[people]\nrowan = just a girl\nbob ross = bad mentor\n[entity]\ncodeScalers = company\n[location]\nalex = city"
		err := parser.LoadFromString(stringToTest)
		if err != nil {
			t.Error()
		}
		got := parser.GetSectionNames()
		expected := []string{"people", "entity", "location"}
		if !slices.Equal(got, expected) {
			t.Errorf("Expected %v , Got %v", expected, got)
		}
	})
	t.Run("One section", func(t *testing.T) {
		parser := NewParser()
		stringToTest := "[people]\nrowan = just a girl"
		err := parser.LoadFromString(stringToTest)
		if err != nil {
			t.Error()
		}
		got := parser.GetSectionNames()
		expected := []string{"people"}
		if !slices.Equal(got, expected) {
			t.Errorf("Expected %v , Got %v", expected, got)
		}
	})
	t.Run("Some sections with empty maps", func(t *testing.T) {
		parser := NewParser()
		stringToTest := "[people]\n[entity]\n[location]\nalex = city"
		err := parser.LoadFromString(stringToTest)
		if err != nil {
			t.Error()
		}
		got := parser.GetSectionNames()
		expected := []string{"people", "entity", "location"}
		if !slices.Equal(got, expected) {
			t.Errorf("Expected %v , Got %v", expected, got)
		}
	})
}

func Test_GetSections(t *testing.T) {
	t.Run("sections with maps", func(t *testing.T) {
		parser := NewParser()
		stringToTest := "[people]\nrowan = just a girl\nbob ross = good mentor\n[entity]\ncodeScalers = company\n[location]\nalex = city"
		err := parser.LoadFromString(stringToTest)
		if err != nil {
			t.Error()
		}
		got := parser.GetSections()
		expected := map[string]map[string]string{
			"people": {
				"rowan":    "just a girl",
				"bob ross": "good mentor",
			},
			"entity": {
				"codeScalers": "company",
			},
			"location": {
				"alex": "city",
			},
		}
		if reflect.DeepEqual(got, expected) {
			t.Errorf("Expected %v , Got %v", expected, got)
		}
	})
	// t.Run("", func(t *testing.T) {

	// })
	// t.Run("", func(t *testing.T) {

	// })

}

func Test_Get(t *testing.T) {
	parser := NewParser()
	stringToTest := "[people]\nrowan = just a girl\nbob ross = good mentor\n[entity]\ncodeScalers = company\n[location]\nalex = city"
	err := parser.LoadFromString(stringToTest)
	if err != nil {
		t.Error()
	}

	t.Run("Get value found", func(t *testing.T) {
		got := parser.Get("people", "rowan")
		expected := "just a girl"
		if got != expected {
			t.Errorf("Expected %v , Got %v", expected, got)
		}
	})

	t.Run("Get value with key not found", func(t *testing.T) {
		got := parser.Get("people", "rawan")
		var expected string
		if got != expected {
			t.Errorf("Expected %v , Got %v", expected, got)
		}
	})
	t.Run("Get value with section not found", func(t *testing.T) {
		got := parser.Get("cat", "luna")
		var expected string
		if got != expected {
			t.Errorf("Expected %v , Got %v", expected, got)
		}
	})
}

func Test_Set(t *testing.T) {
	parser := NewParser()
	stringToTest := "[people]\nrowan = just a girl\nbob ross = good mentor\n[entity]\ncodeScalers = company\n[location]\nalex = city"
	err := parser.LoadFromString(stringToTest)
	if err != nil {
		t.Error()
	}

	t.Run("Set an already present key to new value", func(t *testing.T) {
		parser.Set("people", "bob ross", "bad mentor")
		got := parser.Get("people", "bob ross")
		expected := "bad mentor"
		if got != expected {
			t.Errorf("Expected %v , Got %v", expected, got)
		}
	})

	t.Run("Set a new key", func(t *testing.T) {
		parser.Set("people", "steve", "someone")
		got := parser.Get("people", "steve")
		expected := "someone"
		if got != expected {
			t.Errorf("Expected %v , Got %v", expected, got)
		}
	})
	t.Run("Set a new section with a new key", func(t *testing.T) {
		parser.Set("precious things", "luna", "cat")
		got := parser.Get("precious things", "luna")
		expected := "cat"
		if got != expected {
			t.Errorf("Expected %v , Got %v", expected, got)
		}
	})
}

// func Test_ToString(t *testing.T){
// 	parser := NewParser()
// 	expected := "[people]\nrowan = just a girl\nbob ross = good mentor\n[entity]\ncodeScalers = company\n[location]\nalex = city"

// 	t.Run("ToString same as input", func(t *testing.T) {
// 		stringToTest := "[people]\nrowan = just a girl\nbob ross = good mentor\n[entity]\ncodeScalers = company\n[location]\nalex = city"
// 		parser.LoadFromString(stringToTest)
// 		got := parser.ToString()
// 		if got != expected {
// 			t.Errorf("Expected\n %v ,\n Got:\n %v", expected, got)
// 		}
// 	})
// 	t.Run("Input has redundant spaces", func(t *testing.T) {
// 		stringToTest := "[people]\nrowan   =   just a girl \nbob ross =  good mentor\n\n[entity]\ncodeScalers = company\n[location]\nalex = city"
// 		parser.LoadFromString(stringToTest)
// 		got := parser.ToString()
// 		if got != expected {
// 			t.Errorf("Expected %v , Got %v", expected, got)
// 		}
// 	})
// }
