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
	n, k := readTwoNums(reader)
	a := readNNums(reader, k)
	return solve(n, a)
}

func solve(n int, a []int) int {
	// n := len(a)˝
	first := make([]int, n+1)
	last := make([]int, n+1)
	for i, x := range a {
		if first[x] == 0 {
			first[x] = i + 1
		}
		last[x] = i + 1
	}

	check := func(i int, j int) bool {
		// (i, j) pair
		if first[i] == 0 || first[j] == 0 {
			// i 没有出现过，那么一致使用i，最后再换成j
			// 或者一开始是i，但是马上换成j
			return true
		}
		return last[j] < first[i]
	}

	var res int
	for i := 1; i < n; i++ {
		// (i, i + 1)
		if check(i, i+1) {
			res++
		}
		if check(i+1, i) {
			res++
		}
		if first[i] == 0 {
			res++
		}
	}
	if first[n] == 0 {
		res++
	}
	return res
}
