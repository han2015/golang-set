package mapset

import (
	"fmt"
)

type YourType struct {
	Name string
}

func ExampleIterator() {
	set := NewSetFromSlice([]interface{}{
		&YourType{Name: "Alise"},
		&YourType{Name: "Bob"},
		&YourType{Name: "John"},
		&YourType{Name: "Nick"},
	})

	var found *YourType
	it := set.Iterator()

	for elem := range it.C {
		if elem.Key.(*YourType).Name == "John" {
			found = elem.Key.(*YourType)
			it.Stop()
		}
	}

	fmt.Printf("Found %+v\n", found)

	// Output: Found &{Name:John}
}
