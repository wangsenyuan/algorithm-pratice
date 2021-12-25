package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		S, _ := reader.ReadString('\n')
		T, _ := reader.ReadString('\n')
		res := solve(n, m, S, T)
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

const INF = math.MaxInt32

func solve(n, m int, S, T string) int {
	if n > m {
		return -1
	}
	q := NewQueue(n + 3)
	if n == m {
		// no add-on
		for i := 0; i < n; i++ {
			if S[i] != T[i] {
				q.PushBack(i)
			}
		}
		if q.Size()%2 == 1 {
			return -1
		}
		return helper(*q)
	}
	res := INF
	for i := 0; i+n-1 < m; i++ {
		q.Reset()
		for j := 0; j < n; j++ {
			if S[j] != T[i+j] {
				q.PushBack(i + j)
			}
		}
		if q.Size()%2 == 1 {
			if i > 0 {
				q.PushFront(i - 1)
				res = min(res, helper(*q))
				q.PopFront()
			}
			if i+n-1 < m-1 {
				q.PushBack(i + n)
				res = min(res, helper(*q))
			}
		} else {
			res = min(res, helper(*q))
			if i > 0 && i+n-1 < m-1 {
				q.PushFront(i - 1)
				q.PushBack(i + n)
				res = min(res, helper(*q))
			}
		}
	}

	return res + m - n
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Queue struct {
	arr        []int
	front, end int
}

func NewQueue(n int) *Queue {
	q := new(Queue)
	q.arr = make([]int, n+1)
	q.front = 0
	q.end = 0
	return q
}

func (q *Queue) PushBack(v int) {
	size := len(q.arr)
	q.arr[q.end] = v
	q.end = (q.end + 1) % size
}

func (q *Queue) PushFront(v int) {
	size := len(q.arr)
	prev := (q.front - 1 + size) % size
	q.arr[prev] = v
	q.front = prev
}

func (q *Queue) PopBack() int {
	size := len(q.arr)
	prev := (q.end - 1 + size) % size
	ret := q.arr[prev]
	q.end = prev
	return ret
}

func (q *Queue) PopFront() int {
	size := len(q.arr)
	ret := q.arr[q.front]
	q.front = (q.front + 1) % size
	return ret
}

func (q *Queue) IsEmpty() bool {
	return q.front == q.end
}

func (q *Queue) Size() int {
	if q.end > q.front {
		return q.end - q.front
	}
	return len(q.arr) - q.front + q.end
}

func (q *Queue) Reset() {
	q.front = 0
	q.end = 0
}

func helper(que Queue) int {
	// que size must be even
	var steps int

	for !que.IsEmpty() {
		a := que.PopFront()
		b := que.PopFront()
		steps += b - a
	}
	return steps
}
