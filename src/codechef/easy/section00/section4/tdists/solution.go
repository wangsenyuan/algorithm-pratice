package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		x, y := readTwoNums(reader)
		res := solve(x, y)
		if len(res) == 0 {
			buf.WriteString("NO\n")
			continue
		}
		buf.WriteString("YES\n")
		buf.WriteString(fmt.Sprintf("%d\n", len(res)))
		for j := 2; j <= len(res); j++ {
			buf.WriteString(fmt.Sprintf("%d %d\n", res[j-1], j))
		}
	}
	fmt.Print(buf.String())
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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func solve(even, odd int) []int {
	if odd%2 != 0 {
		return nil
	}
	// let x nodes at even levels
	// let y nodes at odd levels
	// then 2 * x * y = odd
	// x * x + y * y - n = even
	// x * x + y * y + 2 * x * y = even + 2 * odd = n * n
	// x + y = n
	odd /= 2
	var n int
	x := 1
	for x <= odd/x {
		if odd%x != 0 {
			x++
			continue
		}
		y := odd / x
		tmp := int64(x)*int64(x) + int64(y)*int64(y)
		if tmp == int64(even) {
			n = x + y
			break
		}
		x++
	}
	if n == 0 {
		return nil
	}
	y := n - x
	P := make([]int, n)
	for j := 1; j <= y; j++ {
		P[j] = 1
	}
	for j := y + 1; j < n; j++ {
		P[j] = 2
	}
	return P
}
