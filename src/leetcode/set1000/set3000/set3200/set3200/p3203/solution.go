package p3203

func maxHeightOfTriangle(red int, blue int) int {

	f := func(a, b int) int {
		// a is first, and b is second
		var h int
		val := []int{a, b}
		it := 0
		expect := 1
		for val[it] >= expect {
			h++
			val[it] -= expect
			it ^= 1
			expect++
		}

		return h
	}

	return max(f(red, blue), f(blue, red))
}
