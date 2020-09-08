package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		firstLine := readNNums(reader, 4)
		N := firstLine[0]
		K := firstLine[1]
		D := firstLine[2]
		m := firstLine[3]
		A := readNNums(reader, N)
		fmt.Println(solve(N, K, D, m, A))
	}
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

const MAX_N = 5010
const INF = 1 << 60
const N_INF = -INF

var pre [2 * MAX_N]int64
var F [MAX_N][MAX_N]int64
var Q [MAX_N]Pair

func reset(n, k int) {
	for i := 0; i <= n; i++ {
		for j := 0; j <= k; j++ {
			F[i][j] = N_INF
		}
	}
}

func solve(N, K, D, m int, A []int) int64 {
	reset(N, K)
	for i := 1; i <= 2*N; i++ {
		pre[i] = pre[i-1]
		if i <= N {
			pre[i] += int64(A[i-1])
		}
	}
	M := int64(m)
	for i := 1; i <= N; i++ {
		F[i][1] = pre[i-1] + (pre[i+D]-pre[i])*M
	}

	for j := 2; j <= K; j++ {
		var tmp int64 = N_INF
		for i := 1; i <= N; i++ {
			if i-D-1 >= 0 {
				tmp = max(tmp, F[i-D-1][j-1]-pre[i-1])
			}
			F[i][j] = max(F[i][j], tmp+pre[i-1]+(pre[i+D]-pre[i])*M)
		}

		var qh, qt int
		for i := 1; i <= N; i++ {
			for qh < qt {
				if Q[qh].first+D < i {
					qh++
				} else {
					break
				}
			}
			if qh < qt {
				F[i][j] = max(F[i][j], Q[qh].second+pre[i-1]*M+(pre[i+D]-pre[i])*M)
			}

			tmp = F[i][j-1] - pre[i+D]*M
			for qh < qt {
				if Q[qt-1].second <= tmp {
					qt--
				} else {
					break
				}
			}
			Q[qt] = Pair{i, tmp}
			qt++
		}
	}

	var res int64

	for i := 1; i <= N; i++ {
		tmp := F[i][K]
		if i+D < N {
			tmp += pre[N] - pre[i+D]
		}
		res = max(res, tmp)
	}

	return res
}

type Pair struct {
	first  int
	second int64
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
