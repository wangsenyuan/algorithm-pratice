package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	a := readNNums(reader, n)
	b := readString(reader)

	res := solve(a, b)

	fmt.Printf("%d %d\n", res[0], res[1])
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\r' || s[i] == '\n' {
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

const inf = 1e9

func solve(a []int, b string) []int {
	l := []int{-inf, inf}
	r := []int{-inf, inf}

	n := len(a)
	for i := 4; i < n; i++ {
		s := b[i-4 : i+1]
		if s == "00000" {
			l[1] = min(l[1], max(a[i-4], a[i-3], a[i-2], a[i-1], a[i]))
		} else if s == "00001" {
			l[0] = max(l[0], 1+max(a[i-4], a[i-3], a[i-2], a[i-1], a[i]))
		} else if s == "11111" {
			r[0] = max(r[0], min(a[i-4], a[i-3], a[i-2], a[i-1], a[i]))
		} else if s == "11110" {
			r[1] = min(r[1], min(a[i-4], a[i-3], a[i-2], a[i-1], a[i])-1)
		}
	}

	return []int{l[0], r[1]}
}
