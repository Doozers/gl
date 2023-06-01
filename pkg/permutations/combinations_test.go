package permutations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenCombWithRep(t *testing.T) {
	type te struct {
		n              int
		elements       []int
		expectedLength int
	}
	tests := []te{
		{
			n:              2,
			elements:       []int{1, 2, 3},
			expectedLength: 9,
		},
		{
			n:              3,
			elements:       []int{1, 2, 3},
			expectedLength: 27,
		},
		{
			n:              4,
			elements:       []int{1, 2, 3},
			expectedLength: 81,
		},
	}

	for _, test := range tests {
		res := Comb(test.n, test.elements, CombOpts{WithRep: true, WithOrder: true})
		got, want := len(res), test.expectedLength
		assert.Equal(t, got, want, "got %d, want %d", got, want)
	}
}
