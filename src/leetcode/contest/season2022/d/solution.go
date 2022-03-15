package d

import "sort"

const MOD = 1000000007
const X = 1000

func coopDevelop(skills [][]int) int {
	// n <= 100000
	// len(skills[i]) <= 4
	// skills[i][j] <= 1000
	// n * n TLE
	sort.Slice(skills, func(i, j int) bool {
		return len(skills[i]) < len(skills[j])
	})

	mem := make(map[Key]int)

	arr := make([]int, 4)

	var res int

	for i := 0; i < len(skills); i++ {
		cur := skills[i]
		m := len(cur)
		var cnt int
		for state := 1; state < 1<<m; state++ {
			var p int
			for j := 0; j < m; j++ {
				if (state>>j)&1 == 1 {
					arr[p] = cur[j]
					p++
				}
			}
			cnt += mem[NewKey(arr[:p])]
		}
		res += i - cnt

		if res >= MOD {
			res -= MOD
		}

		mem[NewKey(cur)]++
	}

	return res
}

const H1 = 27
const H2 = 31

type Key struct {
	key1 int64
	key2 int64
}

func NewKey(arr []int) Key {
	key1 := getKey(arr, H1)
	key2 := getKey(arr, H2)
	return Key{key1, key2}
}

func getKey(arr []int, base int64) int64 {
	var res int64
	for _, num := range arr {
		res = res*base%MOD + int64(num)
		res %= MOD
	}
	return res
}
