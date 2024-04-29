package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
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

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	a := readNNums(reader, n)
	p := readNNums(reader, n)
	res := solve(a, p)
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	buf.WriteByte('\n')
	fmt.Print(buf.String())
}

const D = 30

func solve(A []int, P []int) []int {
	n := len(A)

	tr := NewTrie()

	for _, num := range P {
		tr.Add(num)
	}

	res := make([]int, n)

	for i, num := range A {
		cur := 0
		for j := D - 1; j >= 0; j-- {
			x := (num >> j) & 1
			if tr.next[cur][x] != 0 && tr.cnt[tr.next[cur][x]] > 0 {
				cur = tr.next[cur][x]
				// keep 0
			} else {
				// has to be 1
				res[i] |= 1 << j
				// 只能走另外一个路径
				cur = tr.next[cur][x^1]
			}
			tr.cnt[cur]--
		}
	}

	return res
}

type Trie struct {
	next [][2]int
	cnt  []int
	id   int
}

func NewTrie() *Trie {
	tr := new(Trie)
	tr.next = make([][2]int, 1000)
	tr.cnt = make([]int, 1000)
	tr.id = 1
	return tr
}

func (tr *Trie) expand() int {
	if tr.id == len(tr.next) {
		tmp := make([][2]int, 1000)
		tr.next = append(tr.next, tmp...)
		tmp2 := make([]int, 1000)
		tr.cnt = append(tr.cnt, tmp2...)
	}
	tr.id++
	return tr.id - 1
}

func (tr *Trie) Add(num int) {
	// from root 0
	var cur int
	for i := D - 1; i >= 0; i-- {
		x := (num >> i) & 1
		tmp := tr.next[cur][x]
		if tmp == 0 {
			tmp = tr.expand()
			tr.next[cur][x] = tmp
		}
		cur = tmp
		tr.cnt[cur]++
	}
}
