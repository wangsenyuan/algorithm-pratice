package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, X, Y := readThreeNums(reader)
		A := readNNums(reader, n)
		res := solve(A, X, Y)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
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

const INF = 1 << 30

func solve(A []int, X, Y int) int {
	n := len(A)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, Y+1)
		for j := 0; j <= Y; j++ {
			dp[i][j] = INF
		}
	}

	dp[0][0] = 0
	// dp[i][j] = min count of elements not picked to get sum j in first i elements
	for i := 0; i < n; i++ {
		for j := 0; j < A[i] && j <= Y; j++ {
			dp[i+1][j] = dp[i][j] + 1
		}
		for j := A[i]; j <= Y; j++ {
			dp[i+1][j] = min(dp[i][j]+1, dp[i][j-A[i]])
		}
	}

	best := INF

	for j1 := X; j1 <= Y; j1++ {
		best = min(best, dp[n][j1])
	}

	fp := make([]int, Y+1)
	for i := 0; i < len(fp); i++ {
		fp[i] = INF
	}

	fp[0] = 0

	que := make([]Pair, Y+1)

	for i := n - 1; i >= 0; i-- {
		for j := Y; j >= A[i]; j-- {
			fp[j] = min(fp[j], fp[j-A[i]]+1)
		}

		var front, end int
		// find min value in the window [Y-j1, X-j1]
		// j2 = Y - j1
		// when j1 + 1, then j2 - 1
		j2 := Y
		for j2 >= X {
			for front < end && que[end-1].second >= fp[j2] {
				end--
			}
			que[end] = Pair{j2, fp[j2]}
			end++
			j2--
		}
		for j1 := 0; j1 <= Y; j1++ {
			for front < end && que[front].first+j1 > Y {
				front++
			}
			if front < end && que[front].first+j1 >= X {
				best = min(best, max(que[front].second, dp[i][j1]))
			}

			if j2 >= 0 {
				for front < end && que[end-1].second >= fp[j2] {
					end--
				}
				que[end] = Pair{j2, fp[j2]}
				end++
			}
			j2--
		}
	}

	if best >= INF {
		return -1
	}

	return best
}

type Pair struct {
	first, second int
}

func solve1(A []int, X, Y int) int {
	n := len(A)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, Y+1)
		for j := 0; j <= Y; j++ {
			dp[i][j] = INF
		}
	}

	dp[0][0] = 0
	// dp[i][j] = min count of elements not picked to get sum j in first i elements
	for i := 0; i < n; i++ {
		for j := 0; j < A[i] && j <= Y; j++ {
			dp[i+1][j] = dp[i][j] + 1
		}
		for j := A[i]; j <= Y; j++ {
			dp[i+1][j] = min(dp[i][j]+1, dp[i][j-A[i]])
		}
	}

	L := Y + 1

	fp := make([]int, 2*L)
	for i := 0; i < len(fp); i++ {
		fp[i] = INF
	}

	update := func(arr []int, p int, v int) {
		p += L
		arr[p] = min(arr[p], v)
		for p > 1 {
			arr[p>>1] = min(arr[p], arr[p^1])
			p >>= 1
		}
	}
	// fp[n][0] = 0
	update(fp, 0, 0)

	get := func(arr []int, p int) int {
		p += L
		return arr[p]
	}

	getRange := func(arr []int, l, r int) int {
		res := INF
		l += L
		r += L
		for l < r {
			if l&1 == 1 {
				res = min(res, arr[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = min(res, arr[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}
	best := INF

	for j1 := X; j1 <= Y; j1++ {
		best = min(best, dp[n][j1])
	}

	for i := n - 1; i >= 0; i-- {
		for j := Y; j >= A[i]; j-- {
			// fp[i][j] = min(fp[i+1][j], fp[i+1][j-A[i]]+1)
			update(fp, j, get(fp, j-A[i])+1)
		}
		for j1 := 0; j1 <= Y; j1++ {
			if dp[i][j1] < INF {
				// j2 >= max(0, X - j1) and j2 <= Y - j1
				// we need to find min value of fp[i][s...e]
				tmp := getRange(fp, max(0, X-j1), Y-j1+1)
				best = min(best, max(dp[i][j1], tmp))
			}
		}
	}

	if best >= INF {
		return -1
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
