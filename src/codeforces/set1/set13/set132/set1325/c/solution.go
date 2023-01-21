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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

func fillNNums(bs []byte, n int, res []int) {
	x := 0
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 2)
	}
	res := solve(n, E)

	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}

func solve(n int, E [][]int) []int {
	degree := make([]int, n)
	pivot := -1
	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		degree[u]++
		degree[v]++
		if degree[u] > 2 {
			pivot = u
		}
		if degree[v] > 2 {
			pivot = v
		}
	}
	ans := make([]int, len(E))
	for i := 0; i < len(ans); i++ {
		ans[i] = -1
	}

	var from int

	if pivot > 0 {
		for i := 0; i < len(E) && from < 3; i++ {
			u, v := E[i][0]-1, E[i][1]-1
			if u == pivot || v == pivot {
				ans[i] = from
				from++
			}
		}
	}

	for i := 0; i < len(E); i++ {
		if ans[i] < 0 {
			ans[i] = from
			from++
		}
	}

	return ans
}
