package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ask := func(y int) int {
		fmt.Printf("%d\n", y)
		return readNum(reader)
	}
	m, n := readTwoNums(reader)
	solve(m, n, ask)
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

func solve(m int, n int, ask func(int) int) {
	// n个p是不知道的
	// 还的好好想想

	p := make([]int, n)
	for i := range n {
		p[i] = ask(1)
		// 如果p[i] = 1, 说明现在回复的是真话
		// p[i] = -1, 说明回复的是假话
		if p[i] == 0 {
			return
		}
	}

	l, r := 1, m
	var id int
	for l <= r {
		mid := (l + r) / 2
		v := ask(mid)
		if v == 0 {
			return
		}
		if p[id%n] == 1 {
			// 它回复的是正确的
			if v < 0 {
				// x < mid
				r = mid
			} else {
				l = mid + 1
			}
		} else {
			// 它回复的是假话
			if v < 0 {
				l = mid + 1
			} else {
				r = mid
			}
		}
		id++
	}
}
