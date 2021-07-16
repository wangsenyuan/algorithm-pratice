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
		n, x := readTwoNums(reader)
		C := readNNums(reader, n)
		T := readNNums(reader, n-1)
		res := solve(n, x, C, T)
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

const INF = int64(1) << 50

func solve(N int, X int, C []int, T []int) int64 {
	// it is better to spend all X
	// can finish in h hours
	// dp[i][j] = at position i, last branch at j, what min cost
	for N > 0 && C[N-1] == 0 {
		N--
	}

	if N == 0 {
		return 0
	}

	arr := make([]Pair, N)

	minCost := func(p int, h int64) int64 {
		var sum int64
		var res int64
		for i := p - 1; i >= 0; i-- {
			cur := arr[i]
			if cur.second > h {
				return INF
			}
			sum += cur.first
			dd := h - cur.second + 1
			res = max(res, (sum+dd-1)/dd)
		}
		return res
	}

	dp := make([]int64, N+1)
	dp[0] = 0
	// h will be >= all
	check := func(h int64) bool {
		// reset dp
		for i := 1; i <= N; i++ {
			dp[i] = INF
			// let find out dp[i][i]
			// sum is the number of person in queue
			var p int
			var t int64
			for j := i; j > 0 && h > t; j-- {
				// we need to find min cost to finish sum in h hours
				// suppose setting up a branch with k cost at i,
				// if this C[j] persons arrive at i, and still need to queue up, then the time is took + (sum + k - 1) / k
				// else if the queue empty, then the time needs is took + (C[j] + k - 1) / k
				if C[j-1] > 0 {
					arr[p] = Pair{int64(C[j-1]), t + 1}
					p++
				}

				dp[i] = min(dp[i], dp[j-1]+minCost(p, h))

				if j > 1 {
					t += int64(T[j-2])
				}
			}
			if dp[i] > int64(X) {
				return false
			}
		}

		return true
	}

	var l, r int64 = 1, INF

	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

type Pair struct {
	first, second int64
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
