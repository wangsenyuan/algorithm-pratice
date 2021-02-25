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
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(n, A)
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

func solve(n int, A []int) int {
	var mx int
	for i := 0; i < n; i++ {
		mx = max(mx, A[i])
	}

	mx++

	cnt := make([]int, mx)
	dp := make([]int, mx)

	for i := 0; i < n; i++ {
		cnt[A[i]]++
	}
	var best int
	for i := 1; i < mx; i++ {
		dp[i] += cnt[i]
		for j := 2 * i; j < mx; j += i {
			dp[j] = max(dp[j], dp[i])
		}
		best = max(best, dp[i])
	}

	return n - best
}

func solve1(n int, A []int) int {
	sort.Ints(A)
	dp := make([]int, A[n-1]+1)

	var best = 1

	for i := 0; i < n; i++ {
		cur := A[i]
		if i == 0 || A[i] != A[i-1] {
			// find divisors of cur
			for x := 1; x <= cur/x; x++ {
				if cur%x == 0 {
					y := cur / x
					dp[cur] = max(dp[cur], max(dp[y], dp[x])+1)
				}
			}
		} else {
			dp[cur]++
		}
		// if we keep cur
		best = max(best, dp[cur])
	}
	return n - best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
