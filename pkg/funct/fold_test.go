package funct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFold(t *testing.T) {
	type te struct {
		old []int
		f   func(int, int) int
		new int
	}
	tests := []te{
		{
			old: []int{2, 4, 6},
			f:   func(a, b int) int { return a + b },
			new: 13,
		},
		{
			old: []int{2, 4, 6},
			f:   func(a, b int) int { return a * b },
			new: 48,
		},
	}

	for _, test := range tests {
		got, want := Fold(1, test.old, test.f), test.new
		assert.Equal(t, got, want, "got %d, want %d", got, want)
	}
}
