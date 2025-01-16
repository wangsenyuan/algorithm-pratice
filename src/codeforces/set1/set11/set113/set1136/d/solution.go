package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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

func process(reader *bufio.Reader) int {
	n, m := readTwoNums(reader)
	p := readNNums(reader, n)
	likes := make([][]int, m)
	for i := range m {
		likes[i] = readNNums(reader, 2)
	}
	return solve(p, likes)
}

func solve(p []int, likes [][]int) int {
	n := len(p)
	pos := make([]int, n)
	for i := 0; i < n; i++ {
		p[i]--
		pos[p[i]] = i
	}
	at := make([][]int, n)
	for _, like := range likes {
		u, v := like[0], like[1]
		u--
		v--
		at[u] = append(at[u], v)
	}
	var move int
	for i := n - 2; i >= 0; i-- {
		u := p[i]
		var cnt int
		for _, v := range at[u] {
			if pos[v] < n && pos[v] > i {
				cnt++
			}
		}
		if cnt+move == n-i-1 {
			move++
			pos[u] = n
		}
	}

	return move
}
