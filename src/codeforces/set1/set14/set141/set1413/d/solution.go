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

	n := readNum(reader)
	events := make([]string, 2*n)

	for i := 0; i < 2*n; i++ {
		events[i] = readString(reader)
	}

	res := solve(events)

	if len(res) > 0 {
		var buf bytes.Buffer
		buf.WriteString("YES\n")
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		fmt.Println(buf.String())
	} else {
		fmt.Println("NO")
	}
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

func solve(events []string) []int {
	n := len(events) / 2
	// 给item进行赋值，如果能够成功，就是yes
	// 假设现在pop 出价格为x的，那么所有在队列中的价格都至少是 x + 1
	// 倒过来考虑
	// 如果现在- x，说明需要x在之前被放进去了，现在把它放进去（但这个时机是还需要确定的）
	// 然后- y, 继续放进去
	// 现在遇到一个 +, 这时候就需要考虑到底这里要放入的是x，还是y了
	// 如果这时放入的是y，
	// 然后遇到 - z， 且 z > x， 但是 z < y, 因为x还在队列中，所以不可能先得到z，再得到z
	// 所以应该是在此时放入x更好
	pq := make(IntHeap, 0, n)

	ans := make([]int, n)
	it := n - 1

	for i := 2*n - 1; i >= 0; i-- {
		cur := events[i]
		if cur[0] == '-' {
			var x int
			readInt([]byte(cur), 2, &x)
			if len(pq) > 0 && pq[0] < x {
				return nil
			}
			heap.Push(&pq, x)
		} else {
			// +
			if len(pq) == 0 {
				return nil
			}
			ans[it] = heap.Pop(&pq).(int)
			it--
		}
	}

	if len(pq) > 0 {
		return nil
	}

	return ans
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
