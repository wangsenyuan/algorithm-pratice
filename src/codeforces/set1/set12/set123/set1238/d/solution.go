package main

import (
	"bufio"
	"fmt"
	"os"
)

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func main() {
	reader := bufio.NewReader(os.Stdin)

	readString(reader)
	s := readString(reader)
	res := solve(s)
	fmt.Println(res)
}

func solve(s string) int {
	n := len(s)
	// 所有的子串
	res := n * (n + 1) / 2
	// BBA 这个也不是good的
	// BBBBA 或者 AAAAB

	last := make([]int, 2)
	last[0] = -1
	last[1] = -1

	for i := 0; i < n; i++ {
		x := int(s[i] - 'A')

		if last[x] < last[x^1] {
			res -= i - last[x] - 1
		}
		last[x] = i
	}

	next := make([]int, 2)
	next[0] = n
	next[1] = n

	for i := n - 1; i >= 0; i-- {
		x := int(s[i] - 'A')
		if next[x^1] < next[x] {
			res -= next[x] - i - 1
		}
		next[x] = i
		if i+1 < n && s[i] != s[i+1] {
			// 这个被多减了一次
			res++
		}
	}

	res -= n

	return res
}
