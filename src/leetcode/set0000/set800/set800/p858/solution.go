package p858

func mirrorReflection(p int, q int) int {
	f := lcm(p, q)

	x, y := f/p, f/q

	if y%2 == 0 {
		return 2
	}

	if x%2 == 1 {
		return 1
	}

	return 0
}

func lcm(a, b int) int {
	x := a * b

	for b > 0 {
		a, b = b, a%b
	}

	return x / a
}
