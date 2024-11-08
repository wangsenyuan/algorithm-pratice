package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
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

	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	queries := make([]int, m)
	for i := range m {
		queries[i] = readNum(reader)
	}

	ans := solve(a, queries)

	var buf bytes.Buffer
	for _, x := range ans {
		buf.WriteString(fmt.Sprintf("%d %d\n", x[0], x[1]))
	}

	fmt.Print(buf.String())
}

func solve(arr []int, queries []int) [][]int {
	x := slices.Max(arr)

	var res [][]int

	n := len(arr)
	front, end := 0, 0

	for {
		a, b := arr[front], arr[(front+1)%n]
		res = append(res, []int{a, b})
		front = (front + 1) % n
		if a > b {
			arr[front] = a
			arr[end] = b
		} else {
			arr[front] = b
			arr[end] = a
		}
		end = (end + 1) % n
		if a == x || b == x {
			break
		}
	}

	tmp := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		tmp[i] = arr[(front+1+i)%n]
	}

	ans := make([][]int, len(queries))

	for i, m := range queries {
		if m <= len(res) {
			ans[i] = res[m-1]
		} else {
			m -= len(res)
			m--
			ans[i] = []int{x, tmp[m%(n-1)]}
		}
	}
	return ans
}
