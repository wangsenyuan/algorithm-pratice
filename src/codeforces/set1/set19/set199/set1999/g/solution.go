package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	ask := func(a, b int) int {
		fmt.Printf("? %d %d\n", a, b)
		return readNum(reader)
	}

	for tc > 0 {
		tc--

		res := solve(ask)

		fmt.Printf("! %d\n", res)
	}
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

func solve(ask func(a int, b int) int) int {
	l, r := 2, 999

	for (r - l + 1) >= 3 {
		m1 := l + (r-l+1)/3
		m2 := r - (r-l+1)/3
		ans := ask(m1, m2)
		if m1*m2 == ans {
			// m2 < x
			l = m2 + 1
		} else if m1*(m2+1) == ans {
			// m1 < x <= m2
			l = m1 + 1
			r = m2
		} else {
			// (m1 + 1) * (m2 + 1) = ans
			// x <= m1
			r = m1
		}
	}

	if l == r {
		return l
	}
	// l + 1 == r
	ans := ask(l, l)
	if l*l == ans {
		return r
	}
	return l
}
