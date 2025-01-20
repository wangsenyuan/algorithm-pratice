package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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

func process(reader *bufio.Reader) int {
	n, p, k := readThreeNums(reader)
	a := readNNums(reader, n)
	s := make([][]int, n)
	for i := range n {
		s[i] = readNNums(reader, p)
	}
	return solve(p, k, a, s)
}

type pair struct {
	first  int
	second int
}

const inf = 1 << 60

func solve(p int, k int, a []int, s [][]int) int {
	// p <= 7
	n := len(a)
	arr := make([]pair, n)
	for i := range n {
		arr[i] = pair{a[i], i}
	}
	slices.SortFunc(arr, func(x, y pair) int {
		return y.first - x.first
	})

	dp := make([]int, 1<<p)
	fp := make([]int, 1<<p)
	for i := range 1 << p {
		dp[i] = -inf
		fp[i] = -inf
	}

	dp[0] = 0

	for i, cur := range arr {
		x := cur.second
		for state := 0; state < 1<<p; state++ {
			cnt := bits.OnesCount(uint(state))
			z := i - cnt
			fp[state] = dp[state]
			if z < k {
				// x在选完state表示的player后，可以作为观众
				fp[state] = max(fp[state], dp[state]+a[x])
			}
			for j := 0; j < p; j++ {
				if (state>>j)&1 == 1 {
					fp[state] = max(fp[state], dp[state^(1<<j)]+s[x][j])
				}
			}
		}
		copy(dp, fp)
	}
	return dp[1<<p-1]
}
