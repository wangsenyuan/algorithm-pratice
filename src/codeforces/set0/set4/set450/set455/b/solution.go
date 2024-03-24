package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, k := readTwoNums(reader)
	words := make([]string, n)
	for i := 0; i < n; i++ {
		words[i] = readString(reader)
	}
	res := solve(k, words)
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
func solve(k int, words []string) string {
	var next [][]int

	next = append(next, make([]int, 26))

	add := func(word string) {
		var cur int
		for i := 0; i < len(word); i++ {
			x := int(word[i] - 'a')
			if next[cur][x] == 0 {
				next = append(next, make([]int, 26))
				next[cur][x] = len(next) - 1
			}
			cur = next[cur][x]
		}
	}

	for _, word := range words {
		add(word)
	}

	win := make([]bool, len(next))
	lose := make([]bool, len(next))

	var dfs func(u int)
	dfs = func(u int) {
		// 假设它必输
		// 哪些节点是肯定可以输的呢？
		// 它要存在一个后继为赢的节点
		// 且该节点，没有走到后继为赢的子节点
		// 如果v是一个当前用户可以必输的节点
		// 那么u如果要必输，它就不能进入v
		// 也就是说，u要必输，它的子节点，不能全部必输
		leaf := true
		lose[u] = false
		win[u] = false
		for i := 0; i < 26; i++ {
			v := next[u][i]
			if v > 0 {
				leaf = false
				dfs(v)
				if !win[v] {
					// 只要有一个胜利后继
					win[u] = true
				}
				if !lose[v] {
					lose[u] = true
				}
			}
		}
		if leaf {
			lose[u] = true
		}

	}

	dfs(0)

	if win[0] && lose[0] {
		return "First"
	}
	if win[0] {
		if k&1 == 1 {
			return "First"
		}
	}

	return "Second"
}
