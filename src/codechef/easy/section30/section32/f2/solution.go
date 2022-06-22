package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		first := readNNums(reader, 4)
		second := readString(reader)
		res := solve(first[0], first[1], first[2], first[3], second)
		buf.WriteString(fmt.Sprintln(fmt.Sprintf("%d", res)))
	}
	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

const MOD = 7051954

const MAX_D = 502

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func solve(N, L, A, B int, mcq string) int {
	var left, right int
	for i := 0; i < N; i++ {
		if mcq[i] == '+' {
			right++
		} else {
			left--
		}
	}

	right = min(L-A, right)
	left = max(-A, left)
	if B-A > right || B-A < left {
		return 0
	}

	size := right - left + 1
	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]int, N+1)
	}
	dp[-left][0] = 1

	plus := NewQueue(N)
	minus := NewQueue(N)

	plus.Push(0)
	minus.Push(0)

	for i := 1; i <= N; i++ {
		if mcq[i-1] == '+' {
			for plus.Size() > 0 {
				j := plus.Pop()
				for k := 0; k < size-1; k++ {
					dp[k+1][i] = modAdd(dp[k+1][i], dp[k][j])
				}
			}
		} else {
			for minus.Size() > 0 {
				j := minus.Pop()
				for k := 1; k < size; k++ {
					dp[k-1][i] = modAdd(dp[k-1][i], dp[k][j])
				}
			}
		}
		plus.Push(i)
		minus.Push(i)
	}

	var ans int

	for i := 0; i <= N; i++ {
		ans = modAdd(ans, dp[B-A-left][i])
	}
	return ans
}

type Queue struct {
	arr []int
	pos int
}

func NewQueue(n int) *Queue {
	q := new(Queue)
	q.arr = make([]int, n)
	q.pos = 0
	return q
}

func (q Queue) Size() int {
	return q.pos
}

func (q *Queue) Pop() int {
	res := q.arr[q.pos-1]
	q.pos--
	return res
}

func (q *Queue) Push(v int) {
	q.arr[q.pos] = v
	q.pos++
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
