package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	home := readNNums(reader, 2)

	n := readNum(reader)

	items := make([][]int, n)
	for i := 0; i < n; i++ {
		items[i] = readNNums(reader, 2)
	}

	tot, path := solve(home, items)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", tot))

	for _, cur := range path {
		buf.WriteString(fmt.Sprintf("%d ", cur))
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

const INF = 1 << 29

func solve(home []int, items [][]int) (int, []int) {
	n := 1 + len(items)

	dist := make([][]int, n)

	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i != j {
				if i == 0 {
					dist[i][j] = distance(home, items[j-1])
				} else if j > 0 {
					dist[i][j] = distance(items[i-1], items[j-1])
				} else {
					dist[i][j] = distance(items[i-1], home)
				}
			}
		}
	}

	N := 1 << (n - 1)

	dp := make([]int, N)
	parent := make([]int, N)
	for i := 0; i < N; i++ {
		dp[i] = INF
	}
	dp[0] = 0

	for state := 0; state < N; state++ {
		if dp[state] >= INF {
			continue
		}
		for i := 1; i < n; i++ {
			if state&(1<<(i-1)) == 0 {
				// i is not picked yet
				for j := i; j < n; j++ {
					if state&(1<<(j-1)) == 0 {
						if dp[state|(1<<(i-1))|(1<<(j-1))] > dp[state]+dist[i][0]+dist[j][0]+dist[i][j] {
							dp[state|(1<<(i-1))|(1<<(j-1))] = dp[state] + dist[i][0] + dist[j][0] + dist[i][j]
							parent[state|(1<<(i-1))|(1<<(j-1))] = state
						}
					}
				}
				// 这个break很重要。
				// Now think in this way. While building the answer for 'i', take any 'on' bit.
				// There are two cases 1. It is picked alone. 2. It is picked with one other object.
				// That's it. This covers all the cases for 'i'.
				// So we can compute for these two cases and then break i.e. only one bit is required. Hope it is clear!
				break
			}
		}
	}
	best := dp[N-1]
	res := make([]int, 0, 2*n)
	res = append(res, 0)

	cur := N - 1

	for cur > 0 {

		prev := parent[cur]

		tmp := prev ^ cur

		for i := 1; i < n; i++ {
			if tmp&(1<<(i-1)) > 0 {
				res = append(res, i)
			}
		}
		res = append(res, 0)
		cur = prev
	}

	return best, res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func distance(a, b []int) int {
	dx := b[0] - a[0]
	dy := b[1] - a[1]
	return dx*dx + dy*dy
}
