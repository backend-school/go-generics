package main

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

type Metrical interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

func Add[T Metrical] (a,b T) T {
	return a + b
}