package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func fillNNums(bs []byte, n int, res []int) {
	x := 0
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
}
func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		N := readNum(reader)
		A := readNNums(reader, N)
		a, b := solve(N, A)
		fmt.Printf("%d %.4f\n", a, b)
	}
}

const MAX_X = 101

var set []bool

func init() {
	set = make([]bool, MAX_X)

	for x := 2; x < MAX_X; x++ {
		if set[x] {
			continue
		}
		for y := x * x; y < MAX_X; y += x {
			set[y] = true
		}
	}
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func solve(N int, A []int) (int, float64) {
	sort.Ints(A)

	res1 := pickOptimaly(N, A)
	res2 := pickRandomly(N, A)
	return res1, res2
}

func pickOptimaly(n int, A []int) int {
	mem := make([][]int, MAX_X)
	for i := 0; i < MAX_X; i++ {
		mem[i] = make([]int, n+1)
	}

	var dfs func(g int, taken int) bool

	dfs = func(g int, taken int) bool {
		if mem[g][taken] != 0 {
			return mem[g][taken] > 0
		}

		var win bool

		if taken < n {
			if g == 1 {
				win = true
			} else {
				var tot int
				for i := 0; i < n; i++ {
					if gcd(A[i], g) == g {
						tot++
					}
				}
				if taken < tot && !dfs(g, taken+1) {
					win = true
				}

				for i := 0; i < n && !win; i++ {
					x := gcd(g, A[i])
					if x > 1 && x != g {
						if !dfs(x, taken+1) {
							win = true
						}
					}
				}
			}
		}
		if win {
			mem[g][taken] = 1
		} else {
			mem[g][taken] = -1
		}

		return win
	}

	if dfs(0, 0) {
		return 1
	}
	return 0
}

func pickRandomly(n int, A []int) float64 {
	mem := make([][]float64, MAX_X)
	for i := 0; i < MAX_X; i++ {
		mem[i] = make([]float64, n+1)
		for j := 0; j <= n; j++ {
			mem[i][j] = -1
		}
	}

	var dfs func(g int, taken int) float64

	dfs = func(g int, taken int) float64 {
		if mem[g][taken] >= 0 {
			return mem[g][taken]
		}
		var res float64

		if g == 1 {
			res = 1
		} else {
			var tot int
			for i := 0; i < n; i++ {
				if gcd(A[i], g) == g {
					tot++
				}
			}
			if taken < tot {
				res += (1.0 - dfs(g, taken+1)) * float64(tot-taken) / float64(n-taken)
			}
			for i := 0; i < n; i++ {
				if gcd(A[i], g) == g {
					continue
				}
				res += (1.0 - dfs(gcd(g, A[i]), taken+1)) / float64(n-taken)

			}
		}

		if res < 1e-8 {
			res = 0
		}

		mem[g][taken] = res

		return res
	}

	return dfs(0, 0)
}
