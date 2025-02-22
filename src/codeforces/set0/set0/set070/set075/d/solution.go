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
	arrs := make([][]int, n)
	for i := range n {
		s, _ := reader.ReadBytes('\n')
		var m int
		pos := readInt(s, 0, &m) + 1
		arrs[i] = make([]int, m)
		for j := range m {
			pos = readInt(s, pos, &arrs[i][j]) + 1
		}
	}
	xs := readNNums(reader, m)
	return solve(arrs, xs)
}

func solve(arrs [][]int, xs []int) int {
	n := len(arrs)
	states := make([]state, n)
	for i, cur := range arrs {
		states[i] = calc(cur)
	}

	best := -inf
	pref := -inf

	for _, i := range xs {
		state := states[i-1]
		best = max(best, state.best)
		best = max(best, pref+state.pref)
		pref = max(pref+state.sum, state.suf)
	}

	return best
}

type state struct {
	pref int
	suf  int
	sum  int
	best int
}

const inf = 1 << 60

func calc(arr []int) state {
	best := -inf
	var sum int
	for _, x := range arr {
		sum += x
		best = max(best, sum)
		if sum < 0 {
			sum = 0
		}
	}
	var pref = -inf
	sum = 0
	for _, x := range arr {
		sum += x
		pref = max(pref, sum)
	}

	var suf = -inf
	sum = 0
	for i := len(arr) - 1; i >= 0; i-- {
		sum += arr[i]
		suf = max(suf, sum)
	}

	return state{pref, suf, sum, best}
}
