package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n, m, B := readThreeNums(reader)
		G := make([][]int, n)
		for i := 0; i < n; i++ {
			G[i] = readNNums(reader, m)
		}
		res := solve(n, m, B, G)
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

const MOD = 1000000007

func solve(n, m int, B int, G [][]int) int {
	// n * m <= 200, min(n, m) <= 15
	if n < m {
		G = transpose(G)
		n, m = m, n
	}

	nums := make([]int, n)
	var X []int
	var Y []int
	if n < 40 && (1<<uint(n+1)/2) < B {
		X = make([]int, (1<<uint(n/2))+1)
		Y = make([]int, (1<<uint(n-n/2))+1)
	}

	meetInMiddle := func(nums []int) int {
		h := n / 2
		a := allSums(nums[:h], X, B)
		b := allSums(nums[h:], Y, B)
		l := b
		r := b

		var res int
		for i := 0; i < a; i++ {
			for l > 0 && X[i]+Y[l-1] >= B {
				l--
			}
			// X[i] + Y[l-1] < B
			for r > 0 && X[i]+Y[r-1] > B {
				r--
			}
			res += r - l
			if res >= MOD {
				res -= MOD
			}
		}

		return res
	}

	mem := make([]map[int]int, n)

	var dfs func(i int, target int) int

	dfs = func(i int, target int) int {
		if target == 0 {
			return 1
		}
		if i == n || target < nums[i] {
			return 0
		}

		if v, found := mem[i][target]; found {
			return v
		}

		a := dfs(i+1, target)
		b := dfs(i+1, target-nums[i])

		a += b
		if a >= MOD {
			a -= MOD
		}

		mem[i][target] = a
		return a
	}

	//dp := make([]int, B+1)
	//knapsack := func(nums []int) int {
	//	for i := 0; i <= B; i++ {
	//		dp[i] = 0
	//	}
	//	dp[0] = 1
	//
	//	for _, num := range nums {
	//		for y := B; y >= num; y-- {
	//			dp[y] += dp[y-num]
	//			if dp[y] >= MOD {
	//				dp[y] -= MOD
	//			}
	//		}
	//	}
	//
	//	return dp[B]
	//}

	find := func(nums []int) int {
		if n < 40 && (1<<uint(n+1)/2) < B {
			return meetInMiddle(nums)
		}

		sort.Ints(nums)

		for i := 0; i < n; i++ {
			mem[i] = make(map[int]int)
		}

		return dfs(0, B)
	}

	// m <= n
	M := 1 << uint(m)
	// M <= 2 ** 15 ~~ 32000
	var res int
	for mask := 1; mask < M; mask++ {
		// take mask columns
		for i := 0; i < n; i++ {
			nums[i] = 0
			for j := 0; j < m; j++ {
				if mask&(1<<uint(j)) > 0 {
					nums[i] += G[i][j]
				}
			}
		}
		res += find(nums)
		if res >= MOD {
			res -= MOD
		}
	}
	return res
}

func allSums(arr []int, res []int, B int) int {
	n := len(arr)
	N := 1 << uint(n)
	var i int
	for mask := 0; mask < N; mask++ {
		var tmp int
		for j := 0; j < n; j++ {
			if mask&(1<<uint(j)) > 0 {
				tmp += arr[j]
				if tmp > B {
					break
				}
			}
		}
		if tmp <= B {
			res[i] = tmp
			i++
		}
	}
	sort.Ints(res[:i])
	return i
}

func transpose(G [][]int) [][]int {
	n := len(G)
	m := len(G[0])
	R := make([][]int, m)
	for i := 0; i < m; i++ {
		R[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			R[j][i] = G[i][j]
		}
	}
	return R
}
