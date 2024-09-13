package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	P := make([][]int, n)
	for i := 0; i < n; i++ {
		P[i] = readNNums(reader, 2)
	}
	res := solve(n, P)
	fmt.Println(res)
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

type pair struct {
	x int
	y int
}

func solve(n int, P [][]int) int {
	to := make([]int, n*2+1)
	// pos表示在背面，还是在正面
	pos := make([]int, n*2+1)
	ps := make([]pair, n)
	for i := range P {
		ps[i] = pair{P[i][0], P[i][1]}
		pos[ps[i].y] = 1
		to[ps[i].x] = ps[i].y
		to[ps[i].y] = ps[i].x
	}

	a := make([]int, n)
	al, ar := 0, n-1
	b := make([]int, n)
	bl, br := 0, n-1
	vis := make([]bool, n*2+2)
	mn, mx := 1, n*2

	var ans int

	for {
		for vis[mn] {
			mn++
		}
		if mn > n*2 {
			break
		}

		cnt := [2]int{}
		t := mn
		for {
			last := 0
			for ; mn <= t; mn++ {
				if vis[mn] {
					continue
				}
				cnt[pos[mn]]++

				vis[mn] = true
				a[al] = mn
				al++

				last = to[mn]
				vis[last] = true
				b[bl] = last
				bl++
			}

			if last == 0 {
				break
			}

			t = last
			last = 0
			for ; mx >= t; mx-- {
				if vis[mx] {
					continue
				}
				cnt[pos[mx]]++

				vis[mx] = true
				a[ar] = mx
				ar--

				last = to[mx]
				vis[last] = true
				b[br] = last
				br--
			}

			if last == 0 {
				break
			}
			t = last
		}
		ans += min(cnt[0], cnt[1])
	}
	slices.Reverse(b)
	if !slices.IsSorted(a) || !slices.IsSorted(b) {
		return -1
	}

	return ans
}

func solve1(n int, P [][]int) int {
	f := make([]int, n+1)
	c := make([]int, n+1)
	for i := 0; i < n; i++ {
		a := P[i][0]
		b := P[i][1]
		var cost int
		if a > b {
			a, b = b, a
			cost = 1
		}
		if a > n {
			return -1
		}
		f[a] = b
		c[a] = cost
	}

	suf_max := make([]int, n+2)
	suf_max[n+1] = -1
	for i := n; i > 0; i-- {
		suf_max[i] = max(suf_max[i+1], f[i])
	}

	prev_min := INF
	var cost0, cost1 int
	var ans int
	seq0 := NewQueue(n)
	seq1 := NewQueue(n)

	for i := 1; i <= n; i++ {
		prev_min = min(prev_min, f[i])
		if seq0.IsEmpty() || f[seq0.Back()] > f[i] {
			seq0.Push(i)
			cost0 += c[i]
		} else if seq1.IsEmpty() || f[seq1.Back()] > f[i] {
			seq1.Push(i)
			cost1 += c[i]
		} else {
			return -1
		}

		if prev_min > suf_max[i+1] {
			s0 := seq0.Size()
			s1 := seq1.Size()
			ans += min(cost0+s1-cost1, cost1+s0-cost0)
			cost0 = 0
			cost1 = 0
			seq0.Reset()
			seq1.Reset()
		}
	}
	return ans
}

type Queue struct {
	arr   []int
	front int
	end   int
}

func NewQueue(n int) *Queue {
	q := new(Queue)
	q.arr = make([]int, n)
	q.front = 0
	q.end = 0
	return q
}

func (q *Queue) Size() int {
	return q.end - q.front
}

func (q *Queue) IsEmpty() bool {
	return q.front == q.end
}

func (q *Queue) Back() int {
	return q.arr[q.end-1]
}

func (q *Queue) Push(v int) {
	q.arr[q.end] = v
	q.end++
}

func (q *Queue) Reset() {
	q.front = 0
	q.end = 0
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
