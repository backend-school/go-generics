package generics

import "golang.org/x/exp/constraints"

func AddInt(a,b int8) int8 {
	return a + b
}

func AddString(a,b int8) int8 {
	return a + b
}

// Type T is too generic. You can't 
// func BadAdd[T any] (a,b T) T {
// 	return a + b
// }

// Metrical interface to allow ops in a metrical space
type Metrical interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

// Ok, T is constrained to metrical spaces
func Add[T Metrical] (a,b T) T {
	return a + b
}