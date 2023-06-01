package funct

func Fold[T interface{}, Y interface{}](initial Y, arr []T, f func(Y, T) Y) Y {
	for _, v := range arr {
		initial = f(initial, v)
	}
	return initial
}
