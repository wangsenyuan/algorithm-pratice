package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	n := readNum(scanner)
	A := readNNums(scanner, n)
	fmt.Println(solve(n, A))
}

func solve(n int, A []int) int {
	left := findLeftBest(n, A)
	right := findRightBest(n, A)
	var best int
	for i := 0; i < n-1; i++ {
		if left[i]+right[i+1] > best {
			best = left[i] + right[i+1]
		}
	}
	return best
}

func findRightBest(n int, A []int) []int {
	trie := &Trie{children: make([]*Trie, 2)}
	trie.Add(30, 0)

	res := make([]int, n)
	var cum int

	for i := n - 1; i >= 0; i-- {
		cum = cum ^ A[i]
		res[i] = trie.Xor(30, cum, 0)
		if i < n-1 && res[i+1] > res[i] {
			res[i] = res[i+1]
		}
		trie.Add(30, cum)
	}
	return res
}

func findLeftBest(n int, A []int) []int {
	trie := &Trie{children: make([]*Trie, 2)}
	trie.Add(30, 0)

	res := make([]int, n)
	var cum int
	for i := 0; i < n; i++ {
		cum = cum ^ A[i]
		res[i] = trie.Xor(30, cum, 0)
		if i > 0 && res[i-1] > res[i] {
			res[i] = res[i-1]
		}
		trie.Add(30, cum)
	}
	return res
}

type Trie struct {
	children []*Trie
}

func (t Trie) Xor(i int, y int, ans int) int {
	if i < 0 {
		return ans
	}

	firstDigit := (y >> uint(i)) & 1
	if t.children[1-firstDigit] != nil {
		return t.children[1-firstDigit].Xor(i-1, y, ans|(1<<uint(i)))
	} else if t.children[firstDigit] != nil {
		return t.children[firstDigit].Xor(i-1, y, ans)
	}
	return ans
}

func (t *Trie) Add(i int, y int) {
	if i < 0 {
		return
	}
	firstDigit := (y >> uint(i)) & 1
	if t.children[firstDigit] == nil {
		t.children[firstDigit] = &Trie{children: make([]*Trie, 2)}
	}
	t.children[firstDigit].Add(i-1, y)
}
