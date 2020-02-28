package main

import (
	"bufio"
	"fmt"
	"math"
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
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		N, Z := readTwoNums(scanner)

		A := readNNums(scanner, N)

		fmt.Println(solve(N, Z, A))
	}
}

func solve(N, Z int, A []int) int64 {
	sort.Ints(A)
	reverse(A)

	dp := make([][]int, Z+1)

	for i := 0; i <= Z; i++ {
		dp[i] = make([]int, N+1)
	}

	for i := 1; i <= N; i++ {
		dp[1][i] = i * A[i-1]
	}

	Q := make([]int, N+1)

	for i := 2; i <= Z; i++ {
		var tail int
		Q[tail] = i - 1

		for j := i; j <= N; j++ {
			for tail >= 1 && int64(dp[i-1][Q[tail]]-dp[i-1][Q[tail-1]]) >= int64(A[j-1])*int64(Q[tail]-Q[tail-1]) {
				tail--
			}
			dp[i][j] = dp[i-1][Q[tail]] + (j-Q[tail])*A[j]
			for tail >= 0 && dp[i-1][Q[tail]] >= dp[i-1][j] {
				tail--
			}

			for tail >= 1 &&
				int64(dp[i-1][j]-dp[i-1][Q[tail]])*int64(Q[tail]-Q[tail-1]) <= int64(dp[i-1][Q[tail]]-dp[i-1][Q[tail-1]])*int64(j-Q[tail]) {
				tail--
			}

			tail++
			Q[tail] = j
		}
	}

	var ans int64 = math.MaxInt64

	for i := Z; i <= N; i++ {
		ans = min(ans, int64(dp[Z][i]))
	}

	return ans
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
func solve1(N, Z int, A []int) int64 {
	sort.Ints(A)

	hulls := make([]*ConvexHull, Z+1)

	for i := 0; i <= Z; i++ {
		hull := make(ConvexHull, 0, N)
		hulls[i] = &hull
	}

	dp := make([][]int64, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int64, Z+1)
		for j := 0; j <= Z; j++ {
			dp[i][j] = math.MaxInt64
		}
	}

	for i := N - 1; i >= 0; i-- {
		a := int64(A[i])
		dp[i][1] = a * int64(N-i)

		for j := 2; j <= Z; j++ {
			dp[i][j] = hulls[j-1].Query(a)
			dp[i][j] -= int64(i) * a
		}

		for j := 1; j <= Z; j++ {
			hulls[j].Insert(NewLine(int64(i), dp[i][j]))
		}
	}

	var res int64 = math.MaxInt64

	for i := 0; i < N; i++ {
		res = min(res, dp[i][Z])
	}

	return res
}

type ConvexHull []Line

func (hull *ConvexHull) Insert(line Line) {
	n := len(*hull)
	for len(*hull) > 1 {
		a := (*hull)[n-2]
		b := (*hull)[n-1]
		if isBad(a, b, line) {
			n--
		} else {
			break
		}
	}
	*hull = (*hull)[:n]
	*hull = append(*hull, line)
}

func (hull ConvexHull) Query(x int64) int64 {
	var low int
	var high = len(hull) - 1

	for high-low > 2 {
		mid := (low + high) >> 1
		v1 := hull[mid].Eval(x)
		v2 := hull[mid+1].Eval(x)

		if v2 < v1 {
			low = mid
		} else {
			high = mid
		}
	}

	var res int64 = math.MaxInt64

	for i := low; i <= high; i++ {
		res = min(res, hull[i].Eval(x))
	}

	return res
}

// a * x + b
type Line struct {
	a int64
	b int64
}

func NewLine(a, b int64) Line {
	return Line{a, b}
}

func (line Line) Eval(x int64) int64 {
	return line.a*x + line.b
}

func isBad(x, y, z Line) bool {
	v1 := x.b - z.b
	v2 := z.a - x.a
	v3 := x.b - y.b
	v4 := y.a - x.a

	return v1*v4 <= v2*v3
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
