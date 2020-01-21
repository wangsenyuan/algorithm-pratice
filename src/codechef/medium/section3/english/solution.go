package main

import (
	"bufio"
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
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n := readNum(scanner)
		words := make([]string, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			words[i] = scanner.Text()
		}
		fmt.Println(solve(n, words))
	}
}

const MAX_L = 1e5

func solve(n int, words []string) int64 {
	trie := NewTrie(MAX_L)

	bak := make([]int, MAX_L)

	for i := 0; i < n; i++ {
		word := words[i]
		rword := reverse(word)

		for j := 0; j < len(word); j++ {
			x := int(word[j] - 'a')
			y := int(rword[j] - 'a')
			bak[j] = x*26 + y
		}

		trie.Insert(bak[:len(word)])
	}

	return trie.GetAns()
}

func reverse(s string) string {
	bs := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}

	return string(bs)
}

type Trie struct {
	count []int
	child [][]int
	nodes int
}

func NewTrie(n int) Trie {
	count := make([]int, n+1)
	child := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		child[i] = make([]int, 676)
	}
	return Trie{count, child, 1}
}

func (trie *Trie) Insert(word []int) {
	count := trie.count
	child := trie.child
	nodes := trie.nodes

	var u int

	for i := 0; i < len(word); i++ {
		if child[u][word[i]] == 0 {
			child[u][word[i]] = nodes
			nodes++
		}
		u = child[u][word[i]]
	}
	count[u]++
	trie.nodes = nodes
}

func (trie *Trie) GetAns() int64 {
	count := trie.count
	child := trie.child

	var loop func(u int, lp int) int64

	loop = func(u int, lp int) int64 {
		var ans int64

		for i := 0; i < 676; i++ {
			if child[u][i] > 0 {
				ans += loop(child[u][i], lp+1)
				count[u] += count[child[u][i]]
			}
		}

		for count[u] >= 2 {
			count[u] -= 2
			ans += int64(lp) * int64(lp)
		}

		return ans
	}

	return loop(0, 0)
}
