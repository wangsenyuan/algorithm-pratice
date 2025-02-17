package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	ask := func(a, b int) int {
		fmt.Printf("? %d %d\n", a, b)
		return readNum(reader)
	}

	tc := readNum(reader)
	for range tc {
		n := readNum(reader)
		res := solve(n, ask)
		var buf bytes.Buffer
		buf.WriteString("!")
		for _, x := range res {
			buf.WriteString(fmt.Sprintf(" %d", x))
		}
		buf.WriteByte('\n')
		fmt.Print(buf.String())
	}
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

func solve(n int, ask func(a int, b int) int) []int {
	res := make([]int, n)
	r := 2

	for r < n {
		ans := ask(1, r)
		if ans == 0 {
			break
		}
		// ans == 1, r是某个path的开端点
		res[r] = 0
		r++
	}
	if r == n {
		return res[1:]
	}
	if r == 2 {
		// 只有一条path
		for r := 2; r < n; r++ {
			res[r] = r - 1
		}
		return res[1:]
	}
	// 1,2.3....r-1
	link := NewDLink(r - 1)
	// r 的服节点是1
	for i := 1; i < r; i++ {
		res[i] = 0
		link.Set(i)
	}
	res[r] = link.Get()
	link.Set(r)
	r++
	for r < n {
		x := link.Get()
		ans := ask(r, x)
		if ans == 1 {
			// need to remove current
			link.Erase()
		} else {
			res[r] = x
			// 更新当前的为r
			link.Set(r)
			r++
		}
	}

	return res[1:]
}

type DLink struct {
	prev []int
	next []int
	val  []int
	pos  int
}

func NewDLink(n int) *DLink {
	prev := make([]int, n)
	next := make([]int, n)
	val := make([]int, n)
	for i := 0; i < n; i++ {
		prev[i] = (i - 1 + n) % n
		next[i] = (i + 1) % n
	}
	return &DLink{prev, next, val, 0}
}

func (link *DLink) Set(v int) {
	link.val[link.pos] = v
	link.pos = link.next[link.pos]
}

func (link *DLink) Get() int {
	res := link.val[link.pos]
	return res
}

func (link *DLink) Erase() {
	pos := link.pos
	next := link.next[pos]
	prev := link.prev[pos]

	if prev != -1 {
		link.next[prev] = next
	}
	if next != -1 {
		link.prev[next] = prev
	}

	link.next[pos] = -1
	link.prev[pos] = -1
	link.pos = next
}
