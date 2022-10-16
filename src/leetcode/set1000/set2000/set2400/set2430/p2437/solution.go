package p2437

const MOD = 1000000007
const H = 30

func modMul(a, b int) int {
	r := int64(a) * int64(b)
	return int(r % MOD)
}

func productQueries(n int, queries [][]int) []int {
	var pws []int

	for i := 0; i < 30; i++ {
		if (n>>i)&1 == 1 {
			pws = append(pws, 1<<i)
		}
	}

	res := make([]int, len(queries))

	for i, cur := range queries {
		l, r := cur[0], cur[1]
		res[i] = 1
		for j := l; j <= r; j++ {
			res[i] = modMul(res[i], pws[j])
		}
	}

	return res
}
