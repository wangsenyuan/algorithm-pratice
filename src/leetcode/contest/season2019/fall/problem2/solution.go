package problem2

func fraction(cont []int) []int {
	// a[n-2] + 1/(a[n-1])
	// = (a[n-2] * a[n-1] + 1) / a[n-1]
	n := len(cont)
	if n == 1 {
		return []int{cont[0], 1}
	}

	var loop func(i int) []int64

	loop = func(i int) []int64 {
		if i == n-1 {
			return []int64{1, int64(cont[i])}
		}
		ret := loop(i + 1)
		b, c := ret[0], ret[1]
		a := int64(cont[i])

		x := a*c + b
		y := c

		g := gcd(x, y)
		return []int64{y / g, x / g}
	}
	ret := loop(0)

	return []int{int(ret[1]), int(ret[0])}
}

func gcd(x, y int64) int64 {
	for y > 0 {
		x, y = y, x%y
	}
	return x
}
