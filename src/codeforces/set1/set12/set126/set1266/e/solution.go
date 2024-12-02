package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}

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

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	a := readNNums(reader, n)
	m := readNum(reader)
	queries := make([][]int, m)
	for i := range m {
		queries[i] = readNNums(reader, 3)
	}
	return solve(a, queries)
}

func solve(a []int, queries [][]int) []int {
	n := len(a)

	m := make([]int, n)

	var sum int
	for i := range n {
		sum += a[i]
	}

	record := make([]map[int]int, n)

	for i := 0; i < n; i++ {
		record[i] = make(map[int]int)
	}

	update := func(v int, mv int) {
		sum -= max(a[v]-m[v], 0)
		m[v] = mv
		sum += max(a[v]-m[v], 0)
	}

	add := func(s int, t int, u int) {
		s--
		u--
		if v, ok := record[s][t]; ok {
			update(v, m[v]-1)
		}

		delete(record[s], t)

		if u < 0 {
			return
		}
		record[s][t] = u
		update(u, m[u]+1)
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		add(cur[0], cur[1], cur[2])
		ans[i] = sum
	}

	return ans
}
