package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	ask := func(l int, r int) int {
		fmt.Printf("? %d %d\n", l, r)
		return readNum(reader)
	}

	n := readNum(reader)

	res := solve(n, ask)

	var buf bytes.Buffer
	buf.WriteString("!")
	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf(" %d", res[i]))
	}
	fmt.Println(buf.String())
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

func solve(n int, ask func(l int, r int) int) []int {
	// 至少要问两个cage
	// 假设知道 s1 = A[1] + A[2]
	//         s2 = A[2] + A[3]
	///        s(n-1) = A[n-1] + A[n]
	//  s1 - s2 = A[1] - A[3]
	//   A[1] + A[2] + A[3] = x
	// s1 - s2 + x = 2 * A[1] + A[2]
	// s1 = A[1] + A[2]
	// -s2 + x = A[1]
	a := ask(1, 2)
	b := ask(2, 3)
	c := ask(1, 3)
	res := make([]int, n+1)
	res[1] = -b + c
	res[2] = a - res[1]
	res[3] = c - res[2] - res[1]
	for i := 4; i <= n; i++ {
		res[i] = ask(i-1, i)
		res[i] = res[i] - res[i-1]
	}

	return res[1:]
}
