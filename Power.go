package main

func Power(m, n int) int {
	if n > 0 {
		return m * Power(m, n-1)
	}
	return 1
}

func PowerMod(m, n int) int {
	if n == 0 {
		return 1
	} else if n%2 == 0 {
		return PowerMod(m*m, n/2)
	} else {
		return m * PowerMod(m*m, (n-1)/2)
	}
}
