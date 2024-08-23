package generics_test

import (
	"testing"

	"github.com/drev74/backshool-go-generics/internal/generics"
	"github.com/stretchr/testify/assert"
)

func Test_gen_types(t *testing.T) {

		t.Run("int", func(t *testing.T) {
			param := 5

			exp := generics.Foo[int]{
				Bar: 5,
				Baz: param,
			}

			act := generics.InitFoo(5, 5)
			assert.EqualValues(t, exp, act)
	})

		t.Run("string", func(t *testing.T) {
			param := "phi"

			exp := generics.Foo[string]{
				Bar: 5,
				Baz: param,
			}

			act := generics.InitFoo(param,5)
			assert.EqualValues(t, exp, act)
	})

		t.Run("struct", func(t *testing.T) {
			param := generics.Poo[string] {
				Moo: "moo",
			}

			exp := generics.Foo[generics.Poo[string]]{
				Bar: 5,
				Baz: param,
			}

			act := generics.InitFoo(param,5)
			assert.EqualValues(t, exp, act)
	})
}