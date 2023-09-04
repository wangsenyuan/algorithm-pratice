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

	t := readNum(reader)

	for t > 0 {
		t--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
	}

	fmt.Print(buf.String())
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

func solve(C []int) bool {
	// 要能够出现 1, 1, 2, 3, 5, .... 这样一个序列
	// 假设当前需要x，那么要么存在x，要么存在一个远比x大的数
	// 假设 x = f(i), 则 cur - x >= f(i + 2)
	// sum(C) = sum(F[...i])
	k := len(C)
	if k == 1 {
		return C[0] == 1
	}

	var sum int64
	for i := 0; i < k; i++ {
		sum += int64(C[i])
	}

	var f []int
	f = append(f, []int{1, 1}...)
	var fsum int64 = 2
	for fsum < sum {
		cur := f[len(f)-2] + f[len(f)-1]
		fsum += int64(cur)
		f = append(f, cur)
	}

	if sum != fsum {
		return false
	}

	pq := make(PQ, 0, len(C))

	for i, c := range C {
		it := new(Item)
		it.id = i
		it.priority = c
		heap.Push(&pq, it)
	}

	n := len(f)

	buf := make([]*Item, 3)

	res := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		expect := f[i]
		var j int
		for j < len(buf) {
			if pq.Len() == 0 {
				return false
			}
			cur := heap.Pop(&pq).(*Item)
			if cur.priority < expect {
				return false
			}
			if i == n-1 || cur.id != res[i+1] {
				res[i] = cur.id
				cur.priority -= expect
				if cur.priority > 0 {
					buf[j] = cur
					j++
				}
				break
			}
			if j == 2 {
				// already 2 items larger than expect
				return false
			}
			buf[j] = cur
			j++
		}
		for j > 0 {
			heap.Push(&pq, buf[j-1])
			j--
		}
	}

	return true
}

type Item struct {
	id       int
	priority int
}

type PQ []*Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	i := x.(*Item)
	*pq = append(*pq, i)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	res := old[n-1]
	*pq = old[:n-1]
	return res
}
