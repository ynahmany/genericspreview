package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func main() {

	addressBook := []Person{{Name: "dan"}}

	// find using a comparable element
	found := Find(addressBook, Person{Name: "dan"})

	fmt.Println(fmt.Sprintf("We found %s", found.Name))

	// find using a predicate
	found = FindPredicate(addressBook, func(p Person) bool {
		return p.Name == "dan"
	})

	fmt.Println(fmt.Sprintf("We found %s with a predicate!", found.Name))

	//notFound := FindZero(addressBook, Person{Name: "not-dan"})
	//fmt.Println(fmt.Sprintf("We found the zero value %s", notFound))
}

func Find[T comparable](slice []T, want T) *T {
	for _, el := range slice {
		if el == want {
			return &el
		}
	}
	return nil
}


func FindPredicate[T any](slice []T, predicate func(T) bool) *T {

	for _, el := range slice {
		if predicate(el) {
			return &el
		}
	}
	return nil
}

func FindZero[T comparable](slice []T, want T) T {
	for _, el := range slice {
		if el == want {
			return el
		}
	}
	var zero T
	return zero
}