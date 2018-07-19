package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	f, err := os.Open("./D-large-practice.in")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	var tc int
	fmt.Fscanf(f, "%d", &tc)

	for i := 1; i <= tc; i++ {
		var H, W int64
		var R int
		fmt.Fscanf(f, "%d %d %d", &H, &W, &R)
		rocks := make([][]int64, R)
		for j := 0; j < R; j++ {
			rocks[j] = make([]int64, 2)
			fmt.Fscanf(f, "%d %d", &rocks[j][0], &rocks[j][1])
		}
		res := solve(H, W, R, rocks)
		fmt.Printf("Case #%d: %d\n", i, res)
	}
}

const MOD = 10007

var cc [][]int64

func init() {
	cc = make([][]int64, MOD)
	cc[0] = make([]int64, MOD)
	cc[0][0] = 1
	for i := 1; i < MOD; i++ {
		cc[i] = make([]int64, MOD)
		cc[i][0] = 1
		cc[i][i] = 1
		for j := 1; j < i; j++ {
			cc[i][j] = (cc[i-1][j-1] + cc[i-1][j])
			if cc[i][j] >= MOD {
				cc[i][j] -= MOD
			}
		}
	}
}

func solve(H, W int64, R int, rocks [][]int64) int {
	var res int64

	for state := 0; state < 1<<uint(R); state++ {
		tmp := calculate(state, rocks, H, W)
		cnt := countBits(state, R)
		if cnt%2 == 0 {
			res += tmp
			if res >= MOD {
				res -= MOD
			}
		} else {
			res -= tmp
			if res < 0 {
				res += MOD
			}
		}
	}

	return int(res)
}

func countBits(state int, size int) int {
	var cnt int
	for i := 0; i < size; i++ {
		if state&(1<<uint(i)) > 0 {
			cnt++
		}
	}
	return cnt
}

func calculate(state int, rocks [][]int64, H, W int64) int64 {
	cnt := countBits(state, len(rocks))
	list := make(Pairs, cnt+2)
	list[0] = Pair{1, 1}
	list[1] = Pair{H, W}
	cnt = 2
	for i := 0; i < len(rocks); i++ {
		if state&(1<<uint(i)) > 0 {
			list[cnt] = Pair{rocks[i][0], rocks[i][1]}
			cnt++
		}
	}
	sort.Sort(list)
	res := int64(1)

	for i := 0; i < cnt-1; i++ {
		a, b := list[i], list[i+1]
		// a.first <= b.first
		if a.second > b.second {
			return 0
		}
		res = (res * calculate1(b.first-a.first, b.second-a.second)) % MOD
	}

	return res
}

func calculate1(a, b int64) int64 {
	x := (2*b - a) / 3
	y := (2*a - b) / 3
	if x+2*y != a || y+2*x != b || x < 0 || y < 0 {
		return 0
	}

	return calculate2(x+y, x)
}

func calculate2(n, k int64) int64 {
	if n < k {
		return 0
	}
	if k == 0 || n == k {
		return 1
	}
	res := int64(1)

	for n > 0 && k > 0 {
		x := n % MOD
		y := k % MOD
		res = (res * choose(int(x), int(y))) % MOD
		n /= MOD
		k /= MOD
	}
	return res
}

func choose(x, y int) int64 {
	if y > x {
		return 0
	}
	return cc[x][y]
}

func solveSmall(H, W, R int, rocks [][]int) int {
	dp := make([][]int, H+1)
	for i := 0; i <= H; i++ {
		dp[i] = make([]int, W+1)
	}

	evil := make([][]bool, H+1)
	for i := 0; i <= H; i++ {
		evil[i] = make([]bool, W+1)
	}

	for _, rock := range rocks {
		i, j := rock[0], rock[1]
		evil[i][j] = true
	}

	dp[1][1] = 1

	for i := 1; i <= H; i++ {
		for j := 1; j <= W; j++ {
			if evil[i][j] {
				continue
			}
			if i-1 >= 1 && j-2 >= 1 {
				dp[i][j] += dp[i-1][j-2]
				if dp[i][j] >= MOD {
					dp[i][j] -= MOD
				}
			}
			if i-2 >= 1 && j-1 >= 1 {
				dp[i][j] += dp[i-2][j-1]
				if dp[i][j] >= MOD {
					dp[i][j] -= MOD
				}
			}
		}
	}

	return dp[H][W]
}

type Pair struct {
	first, second int64
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first || (ps[i].first == ps[j].first && ps[i].second < ps[j].second)
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
