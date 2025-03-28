package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"slices"
	"strings"
)

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
	return strings.TrimSpace(s)
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

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(process(reader))
}

func process(reader *bufio.Reader) string {
	n := readNum(reader)
	names := make([]string, n)
	for i := 0; i < n; i++ {
		names[i] = readString(reader)
	}
	surnames := make([]string, n)
	for i := 0; i < n; i++ {
		surnames[i] = readString(reader)
	}
	return solve(names, surnames)
}

func solve(names []string, surnames []string) string {
	p1 := make([]PQ, 26)
	p2 := make([]PQ, 26)

	for _, name := range names {
		x := int(name[0] - 'A')
		heap.Push(&p1[x], name)
	}

	for _, surname := range surnames {
		x := int(surname[0] - 'A')
		heap.Push(&p2[x], surname)
	}

	var res []string

	for i := range 26 {
		for p1[i].Len() > 0 {
			cur := heap.Pop(&p1[i]).(string)
			if p1[i].Len() >= p2[i].Len() {
				// name比surname更多
				var y int
				for y < i && p2[y].Len() == 0 {
					y++
				}
				if y < i {
					res = append(res, fmt.Sprintf("%s %s", cur, heap.Pop(&p2[y]).(string)))
					continue
				}
				//y == i
				if p2[i].Len() > 0 {
					// 优先匹配它
					res = append(res, fmt.Sprintf("%s %s", cur, heap.Pop(&p2[i]).(string)))
					continue
				}
				y = i + 1
				for y < 26 && p2[y].Len() <= p1[y].Len() {
					y++
				}
				// 肯定能匹配到
				res = append(res, fmt.Sprintf("%s %s", cur, heap.Pop(&p2[y]).(string)))
			} else {
				res = append(res, fmt.Sprintf("%s %s", cur, heap.Pop(&p2[i]).(string)))
			}
		}
	}

	slices.SortFunc(res, func(a, b string) int {
		x := a + ", " + b
		y := b + ", " + a
		if x < y {
			return -1
		}
		if x > y {
			return 1
		}
		return 0
	})

	return strings.Join(res, ", ")
}

type PQ []string

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x any) {
	*pq = append(*pq, x.(string))
}

func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
