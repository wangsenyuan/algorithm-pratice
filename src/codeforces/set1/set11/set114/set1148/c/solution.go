package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	p := readNNums(reader, n)

	res := solve(p)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", cur[0], cur[1]))
	}
	fmt.Print(buf.String())
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

func solve(p []int) [][]int {
	n := len(p)
	h := n / 2
	pos := make([]int, n)
	for i := 0; i < n; i++ {
		p[i]--
		pos[p[i]] = i
	}

	var res [][]int

	swap := func(i, j int) {
		i, j = min(i, j), max(i, j)
		res = append(res, []int{i + 1, j + 1})
		p[i], p[j] = p[j], p[i]
		pos[p[i]] = i
		pos[p[j]] = j
	}

	for i := 0; i < n; i++ {
		// i <= pos[i] 成立
		if p[i] != i {
			if pos[i]-i >= h {
				swap(i, pos[i])
			} else if n-1-pos[i] >= h {
				swap(pos[i], n-1)
				swap(i, n-1)
			} else if i >= h {
				// p[i] - 0 >= h
				swap(0, pos[i])
				swap(0, i)
				swap(0, pos[0])
			} else {
				// i < h but p[i] 不能到 n-1位置
				// i.....0
				// 0 到了pos[i]的位置, i 到了0的位置
				swap(0, pos[i])
				// i到了n-1的位置
				swap(0, n-1)
				// i到了i
				swap(i, pos[i])
				swap(0, pos[0])
			}
		}
	}

	return res
}
