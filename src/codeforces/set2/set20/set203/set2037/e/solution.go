package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)

	ask := func(l int, r int) int {
		fmt.Printf("? %d %d\n", l, r)
		return readNum(reader)
	}

	for range tc {
		n := readNum(reader)
		res := solve(n, ask)
		fmt.Printf("! %s\n", res)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func solve(n int, ask func(l int, r int) int) string {
	ans := ask(1, n)
	if ans == 0 {
		// 000 / 1111 = 0
		return "IMPOSSIBLE"
	}
	if n == 2 {
		return "01"
	}
	buf := make([]byte, n+1)
	for i := 0; i <= n; i++ {
		buf[i] = '0'
	}

	query := func(l int, r int) int {
		if l >= r {
			return 0
		}
		return ask(l, r)
	}

	var cnt0 int
	r := n - 1
	for r > 0 {
		cur := query(1, r)
		if cur < ans {
			cnt0 = ans - cur
			buf[r+1] = '1'
			ans = cur
			r--
			break
		}
		r--
	}
	// 现在知道i前面cnt0的数量了
	for r >= 0 {
		cur := query(1, r)
		if cur < ans {
			buf[r+1] = '1'
			ans = cur
		} else {
			// buf[r+1] = '0'
			if cnt0 > 0 {
				cnt0--
			} else {
				buf[r+1] = '1'
			}
		}
		r--
	}

	return string(buf[1:])
}
