package funct

func Map[T interface{}, U interface{}](old []T, f func(T) U) []U {
	new := make([]U, len(old))
	for i, v := range old {
		new[i] = f(v)
	}
	return new
}
