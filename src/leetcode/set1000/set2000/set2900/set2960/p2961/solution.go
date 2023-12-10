package p2961

func getGoodIndices(variables [][]int, target int) []int {
	var res []int
	for i, cur := range variables {
		a, b, c, m := cur[0], cur[1], cur[2], cur[3]
		tmp := pow(a, b, 10)
		tmp = pow(tmp, c, m)
		if tmp == target {
			res = append(res, i)
		}
	}

	return res
}

func pow(a, b, mod int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = r * a % mod
		}
		a = a * a % mod
		b >>= 1
	}
	return r
}
