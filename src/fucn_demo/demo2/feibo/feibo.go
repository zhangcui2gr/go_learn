package feibo

func func_feibo(n int) int {

	if n <= 2 {
		return 1
	}
	a, b := 1, 2
	for i := 2; i < n; i++ {
		c := a + b
		a = b
		b = c
	}

	return b
}
