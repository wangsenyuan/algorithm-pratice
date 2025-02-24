package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer

	tc := readNum(reader)

	for range tc {
		res := process(reader)
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d ", x))
		}
		buf.WriteByte('\n')
	}

	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	a := readNNums(reader, n)
	return solve(a)
}

const mod = 1_000_000_000 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func mul(a, b int) int {
	return a * b % mod
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = mul(r, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return r
}

type pair struct {
	first  int
	second int
}

func solve(a []int) []int {
	n := len(a)
	res := make([]int, n)
	pq := make(PQ, 0, n)

	var sum int

	for i, x := range a {
		sum = add(sum, x)
		var cnt int
		for x%2 == 0 {
			cnt++
			x /= 2
		}

		for pq.Len() > 0 {
			u := pq[0]
			gt := cnt > 30 || x*(1<<cnt) >= u.first
			if !gt {
				break
			}
			u = heap.Pop(&pq).(pair)
			sum = sub(sum, mul(pow(2, cnt), x))
			sum = add(sum, mul(pow(2, cnt+u.second), x))
			sum = sub(sum, mul(pow(2, u.second), u.first))
			sum = add(sum, u.first)
			cnt += u.second
		}

		if cnt != 0 {
			heap.Push(&pq, pair{x, cnt})
		}
		res[i] = sum
	}

	return res
}

type PQ []pair

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i int, j int) bool {
	return pq[i].first < pq[j].first || pq[i].first == pq[j].first && pq[i].second > pq[j].second
}

func (pq PQ) Swap(i int, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x any) {
	p := x.(pair)
	*pq = append(*pq, p)
}

func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	res := old[n-1]
	*pq = old[:n-1]
	return res
}
