package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	n := readNum(r)

	A := readNInt64s(r, n)

	res := solve(A)

	fmt.Println(res)
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')
	if len(s) == 0 || len(s) == 1 && s[0] == '\n' {
		return readNInt64s(reader, n)
	}
	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

const H = 60

func solve(A []int64) int64 {
	n := len(A)
	suf := make([]int64, n+1)
	for i := n - 1; i >= 0; i-- {
		suf[i] = suf[i+1] ^ A[i]
	}
	best := suf[0]

	t := new(Trie)
	Add(t, 0)
	var pref int64
	for i := 0; i < n; i++ {
		pref ^= A[i]
		Add(t, pref)
		x := suf[i+1]
		cur := t
		var tmp int64
		for j := H - 1; j >= 0; j-- {
			a := (x >> j) & 1
			if cur.children[1^a] != nil {
				tmp |= int64(1 << j)
				cur = cur.children[1^a]
			} else {
				cur = cur.children[a]
			}
		}
		if best < tmp {
			best = tmp
		}
	}
	return best
}

type Trie struct {
	children [2]*Trie
}

func Add(t *Trie, num int64) {
	cur := t
	for i := H - 1; i >= 0; i-- {
		x := (num >> i) & 1
		if cur.children[x] == nil {
			cur.children[x] = new(Trie)
		}
		cur = cur.children[x]
	}
}
