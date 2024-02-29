package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	words := make([]string, n)

	for i := 0; i < n; i++ {
		words[i] = readString(reader)
	}

	res := solve(words)

	fmt.Println(res)
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

func solve(words []string) int {
	t := NewTrie()
	var sum int
	for _, word := range words {
		t.Add(reverse(word), 0)
		sum += len(word)
	}
	var res int

	n := len(words)

	for _, word := range words {
		node := t
		ln := len(word)
		res += sum + ln*n
		for i := 0; i < ln && node != nil; i++ {
			x := int(word[i] - 'a')
			next := node.children[x]
			if next != nil {
				res -= int(next.cnt * 2)
			}
			node = next
		}
	}

	return res
}

func reverse(s string) string {
	buf := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return string(buf)
}

type Trie struct {
	children []*Trie
	cnt      uint32
}

func NewTrie() *Trie {
	res := new(Trie)
	res.children = make([]*Trie, 26)
	return res
}

func (t *Trie) Add(word string, i int) {
	t.cnt++

	if i == len(word) {
		return
	}

	x := int(word[i] - 'a')
	if t.children[x] == nil {
		t.children[x] = NewTrie()

	}
	t.children[x].Add(word, i+1)
}
