package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	res := solve(n)
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

const inf = 1 << 30

func solve(n int) int {
	var step func(n int, m int) int

	step = func(n int, m int) int {
		if m == 0 {
			return inf
		}
		if m == 1 {
			return n - 1
		}
		return step(m, n%m) + n/m
	}

	ans := n - 1
	for i := 1; i < n; i++ {
		ans = min(ans, step(n, i))
	}
	return ans
}
