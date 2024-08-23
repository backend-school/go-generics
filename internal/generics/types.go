package generics

type Foo[T any] struct {
	Bar int 
	Baz T
}

type Poo[T comparable] struct {
	Moo T
}

func InitFoo[T any] (v T, bar int) Foo[T] {
	return Foo[T] {
		Bar: bar,
		Baz: v,
	}
}
