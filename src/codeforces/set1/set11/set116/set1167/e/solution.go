package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, x := readTwoNums(reader)
	a := readNNums(reader, n)
	res := solve(n, x, a)

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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

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

func solve(n int, x int, a []int) int {
	first := make([]int, x+1)
	last := make([]int, x+1)

	for i, num := range a {
		if first[num] == 0 {
			first[num] = i + 1
		}
		last[num] = i + 1
	}

	// g[u]等于保留1...u时，且数组是递增的，的最大下标
	g := make([]int, x+2)
	g[0] = 0
	l := 1
	for ; l <= x; l++ {
		g[l] = g[l-1]
		if first[l] == 0 {
			// 这个数不存在
			continue
		}
		if first[l] < g[l] {
			// break the rule
			g[l] = 0
			break
		}
		g[l] = last[l]
	}
	// g[l] 是递增的最后一个位置
	l--

	if l == x {
		return (x + 1) * x / 2
	}

	var res int
	pos := n + 1
	for r := x; r > 0; r-- {
		// r fixed to be deleted
		for l > 0 && g[l] > pos {
			l--
		}
		res += l + 1
		if first[r] == 0 {
			continue
		}
		if last[r] > pos {
			// r必须包括在内
			break
		}

		pos = first[r]
	}
	return res
}
