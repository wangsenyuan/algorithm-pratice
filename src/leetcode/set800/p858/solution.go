package p858

func mirrorReflection(p int, q int) int {
	for p&1 == 0 && q&1 == 0 {
		p >>= 1
		q >>= 1
	}

	return 1 - p%2 + q%2
}
