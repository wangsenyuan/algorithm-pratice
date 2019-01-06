package p970

func powerfulIntegers(x int, y int, bound int) []int {

	cal := func(num int) []int {
		if num == 1 {
			return []int{1}
		}
		res := make([]int, 0, 100)
		p := 1

		for p < bound {
			res = append(res, p)
			p *= num
		}

		return res
	}

	xx := cal(x)
	yy := cal(y)

	set := make(map[int]bool)
	res := make([]int, 0, len(xx)+len(yy))

	for _, a := range xx {
		for _, b := range yy {
			if !set[a+b] && a+b <= bound {
				set[a+b] = true
				res = append(res, a+b)
			}
		}
	}
	return res
}
