package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, q, m := readThreeNums(reader)
	a := readNNums(reader, n)
	queries := make([][]int, q)
	for i := 0; i < q; i++ {
		queries[i] = readNNums(reader, 3)
	}
	b := readNNums(reader, m)

	res := solve(a, queries, b)
	var buf bytes.Buffer
	for i := 0; i < m; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}

	buf.WriteByte('\n')

	fmt.Print(buf.String())
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

func solve(a []int, queries [][]int, b []int) []int {
	n := len(a)

	pos := make([]int, n)
	for i := 0; i < n; i++ {
		pos[i] = -1
	}

	process := func(x int) int {
		if pos[x] >= 0 {
			return pos[x]
		}

		pos[x] = x

		for i := len(queries) - 1; i >= 0; i-- {
			cur := queries[i]
			l, r := cur[1]-1, cur[2]-1
			if l <= pos[x] && pos[x] <= r {
				if cur[0] == 1 {
					pos[x]--
					if pos[x] < l {
						pos[x] = r
					}
				} else {
					d := pos[x] - l
					pos[x] = r - d
				}
			}
		}

		return pos[x]
	}

	ans := make([]int, len(b))
	for i, x := range b {
		ans[i] = a[process(x-1)]
	}

	return ans
}
