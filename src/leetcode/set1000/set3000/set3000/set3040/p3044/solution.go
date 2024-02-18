package p3044

const N = 1000010

var lpf [N]int

func init() {
	var primes []int
	for i := 2; i < N; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for j := 0; j < len(primes); j++ {
			if i*primes[j] >= N {
				break
			}
			lpf[i*primes[j]] = primes[j]
			if i%primes[j] == 0 {
				break
			}
		}
	}
}

func mostFrequentPrime(mat [][]int) int {
	// m <= 6
	// 1000000
	n := len(mat)
	m := len(mat[0])

	freq := make(map[int]int)

	var dfs func(x int, y int, dx int, dy int, num int)
	dfs = func(x int, y int, dx int, dy int, num int) {
		if num > 10 && lpf[num] == num {
			freq[num]++
		}
		x += dx
		y += dy
		if x >= 0 && x < n && y >= 0 && y < m {
			dfs(x, y, dx, dy, num*10+mat[x][y])
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					if dx == 0 && dy == 0 {
						continue
					}
					dfs(i, j, dx, dy, mat[i][j])
				}
			}
		}
	}
	if len(freq) == 0 {
		return -1
	}
	var mv int
	for _, v := range freq {
		mv = max(mv, v)
	}
	var ans int
	for k, v := range freq {
		if v == mv {
			ans = max(ans, k)
		}
	}
	return ans
}
