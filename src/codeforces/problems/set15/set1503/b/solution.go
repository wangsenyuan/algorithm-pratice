package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	a := readNum(reader)

	ask := func(b int, i int, j int, times int) int {
		fmt.Printf("%d %d %d\n", b, i, j)
		if times == n*n {
			return 0
		}
		return readNum(reader)
	}

	solve(n, a, ask)

}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(n int, a int, ask func(b int, i int, j int, times int) int) {
	// if n is even
	// 第一步在1，1处输出b(not a), 如果反馈不是b，可以在紧邻位置输出a
	// 对于位置（i，j), 它被top 和 left控制
	// 如果 top == left, 则没有问题
	que := make([][]Pair, 2)
	for i := 0; i < 2; i++ {
		que[i] = make([]Pair, 0, n*n/2)
	}
	grid := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		grid[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			k := (i + j) % 2
			que[k] = append(que[k], Pair{i, j})
		}
	}
	times := 1
	for len(que[0]) > 0 || len(que[1]) > 0 {
		if a == 1 {
			// put 2
			if len(que[0]) > 0 {
				b := 2
				cur := que[0][len(que[0])-1]
				a = ask(b, cur.first, cur.second, times)
				que[0] = que[0][:len(que[0])-1]
			} else {
				b := 3
				cur := que[1][len(que[1])-1]
				a = ask(b, cur.first, cur.second, times)
				que[1] = que[1][:len(que[1])-1]
			}
		} else if a == 2 {
			if len(que[1]) > 0 {
				b := 1
				cur := que[1][len(que[1])-1]
				a = ask(b, cur.first, cur.second, times)
				que[1] = que[1][:len(que[1])-1]
			} else {
				b := 3
				cur := que[0][len(que[0])-1]
				a = ask(b, cur.first, cur.second, times)
				que[0] = que[0][:len(que[0])-1]
			}
		} else {
			if len(que[0]) > 0 {
				b := 2
				cur := que[0][len(que[0])-1]
				a = ask(b, cur.first, cur.second, times)
				que[0] = que[0][:len(que[0])-1]
			} else {
				b := 1
				cur := que[1][len(que[1])-1]
				a = ask(b, cur.first, cur.second, times)
				que[1] = que[1][:len(que[1])-1]
			}
		}
		times++
	}

}

type Pair struct {
	first  int
	second int
}
