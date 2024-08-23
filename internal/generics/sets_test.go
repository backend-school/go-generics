package generics_test

import (
	"testing"

	"github.com/drev74/backshool-go-generics/pkg/functional"
	"github.com/stretchr/testify/assert"
)

func Test_gen_sets(t *testing.T) {

		pack := []int{0,1,2}
		packMap := make (map[int]int, 0)

		packMap[0] = 0
		packMap[1] = 1
		packMap[2] = 2

		t.Run("slice int contains", func(t *testing.T) {
		 res := functional.Contains(pack, 4)
		 assert.False(t, res)
	})	

		t.Run("slice int fold", func(t *testing.T) {
		 res := functional.Fold(pack, 0, func (a, b int) int{
			return a+b
		 })

		 assert.Equal(t, res, 3)
	})	
}