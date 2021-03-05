package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	//tc := readNum(reader)
	var tc int
	fmt.Scanf("%d", &tc)

	var buf bytes.Buffer
	for tc > 0 {
		tc--
		var n, D int
		fmt.Scanf("%d %d", &n, &D)
		d := make([]int, n)
		c := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scanf("%d %d", &d[i], &c[i])
		}
		l, x := solve(n, D, d, c)
		buf.WriteString(fmt.Sprintf("%d %d", l, x))
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

func solve(n int, D int, d []int, C []int) (int, int) {
	// d[i] <= D
	C[n-1] = 0
	find := func(x int) int {
		var cnt int
		var i int
		for i < n {
			cnt++
			dist := D
			pivot := -1
			for i < n && dist > 0 {
				if dist >= d[i] {
					if C[i] <= x {
						pivot = i
					}
					dist -= d[i]
				} else {
					break
				}
				i++
			}
			if pivot < 0 {
				return INF
			}
			i = pivot + 1
		}
		return cnt
	}

	L := find(INF)

	left, right := 0, INF

	for left < right {
		mid := (left + right) >> 1
		if find(mid) == L {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return L, right
}

func solve1(n int, D int, d []int, C []int) (int, int) {
	// d[i] <= D
	P := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		P[i] = P[i-1] + int64(d[i-1])
	}
	DD := int64(D)
	dp := make([]int, n+1)
	que := make([]int, n+1)
	// dp[i] = min days need when chef breaks at hotel i
	find := func(x int) int {
		dp[n] = INF
		var front, end int
		que[end] = 0
		end++
		for i := 1; i <= n; i++ {
			if i < n && C[i-1] > x {
				continue
			}
			for front < end && P[que[front]]+DD < P[i] {
				front++
			}

			if front == end {
				// never able to reach N
				return INF
			}
			dp[i] = dp[que[front]] + 1
			que[end] = i
			end++
		}
		return dp[n]
	}

	L := find(INF)

	left, right := 0, INF

	for left < right {
		mid := (left + right) >> 1
		if find(mid) == L {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return L, right
}
