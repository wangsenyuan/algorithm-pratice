package p1079

var C [][]int

func init() {
	C = make([][]int, 8)

	for i := 0; i < 8; i++ {
		C[i] = make([]int, 8)
	}

	C[0][0] = 1

	for i := 1; i < 8; i++ {
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i; j++ {
			C[i][j] = C[i-1][j-1] + C[i-1][j]
		}
	}
}

func numTilePossibilities(tiles string) int {
	cnt := make(map[int]int)
	for i := 0; i < len(tiles); i++ {
		cnt[int(tiles[i]-'A')]++
	}

	id := make([]int, 26)
	var p int
	for k := range cnt {
		id[k] = p
		p++
	}

	cnt2 := make([]int, p)
	for k, v := range cnt {
		cnt2[id[k]] = v
	}

	calc := func(use []int) int {
		var sum int
		for _, x := range use {
			sum += x
		}
		res := 1
		for i := 0; i < p; i++ {
			res *= C[sum][use[i]]
			sum -= use[i]
		}
		return res
	}

	//n := len(tiles)

	var dfs func(i int, use []int) int

	dfs = func(i int, use []int) int {
		if i == p {
			return calc(use)
		}
		var ans int
		for x := 0; x <= cnt2[i]; x++ {
			use[i] = x
			ans += dfs(i+1, use)
		}
		return ans
	}
	return dfs(0, make([]int, p)) - 1
}
