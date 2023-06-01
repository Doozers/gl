package permutations

func genCombWithRepAndOrder[T interface{}](n int, elements []T, currentCombination []T, combinations *[][]T) {
	if n == 0 {
		*combinations = append(*combinations, append([]T(nil), currentCombination...))
	} else {
		for _, element := range elements {
			currentCombination = append(currentCombination, element)
			genCombWithRepAndOrder(n-1, elements, currentCombination, combinations)
			currentCombination = currentCombination[:len(currentCombination)-1]
		}
	}
}

type CombOpts struct {
	WithRep   bool
	WithOrder bool
}

func Comb[T interface{}](n int, elements []T, opts CombOpts) [][]T {
	var combinations [][]T
	if opts.WithRep && opts.WithOrder {
		genCombWithRepAndOrder(n, elements, []T{}, &combinations)
	} else {
		panic("Not implemented")
	}
	return combinations
}
