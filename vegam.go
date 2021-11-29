package vegam

func TestFunc(a, b int) int {
	return a + b
}

func SubtractInt(a, b int) int {
	return a - b
}

func MultiplyInt(a, b int) int {
	return a * b
}

func MultiplyAndAccumulate(a, b, c int) int {
	product := a * b
	sum := product + c

	return sum
}
