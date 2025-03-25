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
	n := readNum(reader)
	a := readNNums(reader, n)
	return solve(a)
}

const H = 30

func solve(a []int) int {
	// 如果a[i][d]进行切分？

	var dfs func(d int, arr []int) int

	dfs = func(d int, arr []int) int {
		if d < 0 {
			return 0
		}
		tmp := make([][]int, 2)
		for _, num := range arr {
			x := (num >> d) & 1
			tmp[x] = append(tmp[x], num)
		}
		if len(tmp[0]) == 0 {
			return dfs(d-1, tmp[1])
		}
		if len(tmp[1]) == 0 {
			return dfs(d-1, tmp[0])
		}
		res := min(dfs(d-1, tmp[0]), dfs(d-1, tmp[1]))
		return 1<<d | res
	}

	return dfs(H-1, a)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
