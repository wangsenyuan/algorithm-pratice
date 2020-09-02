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
		n, m := readTwoNums(reader)
		A := readNNums(reader, n)
		fmt.Println(solve(n, m, A))
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

func solve(n int, m int, A []int) int {
	if n == 1 {
		return 0
	}
	var v = -1
	for i := 0; i < n; i++ {
		if A[i] == -1 {
			continue
		}
		if v < 0 {
			v = (A[i] + i) & 1
		} else if v != (A[i]+i)&1 {
			return -1
		}
	}

	if v < 0 {
		return (n - 2) / (m - 1)
	}

	var curX = -1
	var curY = -1

	for i := 0; i < n; i++ {
		if A[i] < 0 {
			continue
		}
		if curX < 0 {
			curX = i
			curY = A[i]
			continue
		}
		if i-curX < abs(curY-A[i]) {
			return -1
		}
	}

	arr := make([]Pair, 0, n)

	for i := 0; i < n; i++ {
		if A[i] > 0 {
			arr = append(arr, Pair{i, A[i]})
		}
	}

	for i := 1; i < len(arr); i++ {
		if abs(arr[i].second-arr[i-1].second) > abs(arr[i].first-arr[i-1].first) {
			return -1
		}
	}
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, 2)
	}
	if arr[0].second-1 < arr[0].first {
		dp[0][1] = (arr[0].first - arr[0].second + 1 + m - 2) / (m - 1)
	}
	if m-arr[0].second < arr[0].first {
		dp[0][0] = (arr[0].first - (m - arr[0].second) + m - 2) / (m - 1)
	}
	s := len(arr)

	var p int
	for i := 0; i < s-1; i++ {
		q := 1 - p
		dp[p][0] = min(dp[p][0], dp[p][1]+1)
		dp[p][1] = min(dp[p][1], dp[p][0]+1)
		dp[q][0] = INF
		dp[q][1] = INF

		dir := GOING_UP
		cost := play(m, arr[i].second, arr[i].first, arr[i+1].second, arr[i+1].first, &dir)
		if dir == GOING_UP {
			dp[q][1] = min(dp[q][1], dp[p][1]+cost)
		} else {
			dp[q][0] = min(dp[q][0], dp[p][1]+cost)
		}

		dir = GOING_DOWN
		cost = play(m, arr[i].second, arr[i].first, arr[i+1].second, arr[i+1].first, &dir)

		if dir == GOING_UP {
			dp[q][1] = min(dp[q][1], dp[p][0]+cost)
		} else {
			dp[q][0] = min(dp[q][0], dp[p][0]+cost)
		}
		p = q
	}
	dp[p][0] = min(dp[p][0], dp[p][1]+1)
	dp[p][1] = min(dp[p][1], dp[p][0]+1)
	res := INF
	if arr[len(arr)-1].first == n-1 {
		res = min(dp[p][0], dp[p][1])
	} else {
		if n-1-arr[s-1].first > m-arr[s-1].second {
			res = dp[p][1] + (n-1-arr[s-1].first-(m-arr[s-1].second)+(m-2))/(m-1)
		} else {
			res = dp[p][1]
		}

		if n-1-arr[s-1].first > arr[s-1].second-1 {
			res = min(res, dp[p][0]+(n-1-arr[s-1].first-(arr[s-1].second-1)+(m-2))/(m-1))
		} else {
			res = min(res, dp[p][0])
		}
	}

	if res >= INF {
		return -1
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Pair struct {
	first, second int
}

const GOING_UP = 1
const GOING_DOWN = -1
const INF = 1 << 20

func play(M int, startY, startX, endY, endX int, dir *int) int {
	var cost int
	cur := startY
	idx := startX

	for idx < endX {
		if *dir == GOING_UP && cur-endY == endX-idx {
			cost++
			*dir = GOING_DOWN
			break
		}
		if *dir == GOING_DOWN && endY-cur == endX-idx {
			cost++
			*dir = GOING_UP
			break
		}

		if *dir == GOING_UP {
			if cur == M {
				cost++
				*dir = GOING_DOWN
				cur = M - 1
			} else {
				cur++
			}
		} else {
			if cur == 1 {
				cost++
				*dir = GOING_UP
				cur = 2
			} else {
				cur--
			}
		}
		idx++
	}

	return cost
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
