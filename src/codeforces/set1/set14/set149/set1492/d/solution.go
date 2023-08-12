package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	a, b, k := readThreeNums(reader)
	ok, x, y := solve(a, b, k)

	if !ok {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
		fmt.Println(x)
		fmt.Println(y)
	}
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

func solve(a int, b int, k int) (bool, string, string) {
	// x has a 0s, b 1s
	// y has a 0s, b 1s
	// x - y has k 1s
	// (x >= y)
	// 如果k=0, easy => let x = y, and b leading 1s, plus a 0s
	// 10000 - 00001 => 1111 也就是说 一个1 + x个0， 可以长生 x个1
	// 假设用 1个1和k个0 => k个1 (有足够0的情况下）
	// 也就是说k的数量 <= a才行， 否则肯定不行的

	n := a + b
	x := make([]byte, n)
	y := make([]byte, n)
	for i := 0; i < n; i++ {
		if i < b {
			x[i] = '1'
		} else {
			x[i] = '0'
		}
	}
	copy(y, x)

	if k == 0 {
		return true, string(x), string(y)
	}

	if k > a+b-2 {
		// x will have more 0 than y
		return false, "", ""
	}

	// 1111000
	// 1011001
	// 0011111

	for i := 1; i < b; i++ {
		if i+k < n && y[i+k] == '0' {
			y[i] = '0'
			y[i+k] = '1'
			return true, string(x), string(y)
		}
	}
	return false, "", ""
}
