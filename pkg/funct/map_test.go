package funct

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	type te struct {
		name string
		old  []int
		f    func(int) string
		new  []string
	}
	tests := []te{
		{
			name: "empty",
			old:  []int{2, 4, 6},
			f:    func(i int) string { return strconv.Itoa(i) },
			new:  []string{"2", "4", "6"},
		},
		{
			name: "empty",
			old:  []int{2, 4, 6},
			f:    func(i int) string { return strconv.Itoa(i * 2) },
			new:  []string{"4", "8", "12"},
		},
	}

	for _, test := range tests {
		got, want := Map(test.old, test.f), test.new
		assert.Equal(t, want, got, test.name)
	}
}
