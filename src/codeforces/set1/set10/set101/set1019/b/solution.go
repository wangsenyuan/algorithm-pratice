package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	ask := func(i int) int {
		fmt.Printf("? %d\n", i)
		return readNum(reader)
	}

	res := solve(n, ask)

	fmt.Printf("! %d\n", res)
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

func solve(n int, ask func(int) int) int {
	vis := make([]bool, n+1)
	a := make([]int, n+1)

	h := n / 2

	// get b[i] = a[i] - a[i + h]
	query := func(i int) int {
		if !vis[i] {
			a[i] = ask(i)
			vis[i] = true
		}
		if !vis[i+h] {
			a[i+h] = ask(i + h)
			vis[i] = true
		}
		return a[i] - a[i+h]
	}

	cur := query(1)
	if cur == 0 {
		return 1
	}
	if abs(cur)&1 == 1 {
		return -1
	}
	// cur = -2 or cur = 2
	l, r := 1, h
	for l < r {
		mid := (l + r) / 2
		tmp := query(mid)
		if tmp == 0 {
			return mid
		}
		if sign(tmp) == sign(cur) {
			l = mid + 1
			cur = tmp
		} else {
			r = mid
		}
	}

	return l
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func sign(num int) int {
	if num < 0 {
		return -1
	}
	if num == 0 {
		return 0
	}
	return 1
}
