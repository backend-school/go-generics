package main

import "golang.org/x/exp/constraints"


func LessInt(a,b int8) bool {
	return a < b
}


func LessString(a,b string) bool {
	return a < b
}


// T is too generic. How to compare values from non-comparable sets ?
// func BadLess[T any] (a, b T) bool {
// 	return a < b
// }

type Ordered interface {
	// constraints.Integer | constraints.Float | constraints.Complex | ~string
	constraints.Integer | constraints.Float | ~string
}

// Now okay, accept only comparable values
func Less[T Ordered] (a,b T) bool {
	return a < b
}