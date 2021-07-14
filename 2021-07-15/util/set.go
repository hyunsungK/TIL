package util 

import (
	"fmt"
)


type Set struct {
	Element []string
}

func NewSet(elements ...string) *Set {
	fmt.Println(elements)
	return &Set{elements}
}

func Union(first *Set, second * Set) *Set {
	return &Set{ append(first.Element, second.Element...)}
}

// func Intersection(first *Set, second * Set)  {
// 	for index, element := range first.Element {
// 		fmt.Println(index, element)
// 		for index2, 
// 	}
// }