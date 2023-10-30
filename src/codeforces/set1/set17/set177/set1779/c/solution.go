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

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		a := readNNums(reader, n)
		res := solve(a, m)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(a []int, m int) int {
	// a[0].... + a[m-1]
	// 要最小
	// let pref[i] = sum(0...i)
	// pref[m] - pref[m-1]
	// pref[m+1] - pref[m-1]
	// ..
	// sum(a[m...i]) >= 0
	// sum(a[j...m-1]) <= 0 if j <= m - 1
	m--
	var ans int
	n := len(a)
	var sum int
	pq := make(IntHeap, 0, n-m+1)
	for i := m + 1; i < n; i++ {
		// sum < 0, 需要把其中最小的数(绝对值最大的负数, 变成正数)
		heap.Push(&pq, a[i])
		sum += a[i]
		for sum < 0 {
			x := heap.Pop(&pq).(int)
			sum -= 2 * x
			ans++
		}
		// sum >= 0
	}
	// 后面部分是对的，但前面部分就不对了
	// pref[i] >= pref[m]
	// pref[i] - pref[m] >= 0
	// pref[m] - pref[i] <= 0
	sum = a[m]
	pq = make(IntHeap, 0, m+1)
	heap.Push(&pq, -a[m])
	for i := m - 1; i >= 0; i-- {
		for sum > 0 {
			x := heap.Pop(&pq).(int)
			x *= -1
			sum -= 2 * x
			ans++
		}
		sum += a[i]
		heap.Push(&pq, -a[i])
	}
	return ans
}

type IntHeap []int

func (pq IntHeap) Len() int {
	return len(pq)
}

func (pq IntHeap) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq IntHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *IntHeap) Push(x interface{}) {
	num := x.(int)
	*pq = append(*pq, num)
}

func (pq *IntHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	res := old[n-1]
	*pq = old[:n-1]
	return res
}
