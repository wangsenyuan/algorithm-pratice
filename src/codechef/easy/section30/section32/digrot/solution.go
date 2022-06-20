package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	// hint(105)
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		res := solve(n)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

func solve(n int) int {
	// n <= 10e8
	// r(n) <= 10e9

	lrot := func(n int) int {
		var buf []int
		for n > 0 {
			buf = append(buf, n%10)
			n /= 10
		}
		reverse(buf)
		reverse(buf[1:])
		reverse(buf)
		return toNum(buf)
	}
	rrot := func(n int) int {
		var buf []int
		for n > 0 {
			buf = append(buf, n%10)
			n /= 10
		}
		reverse(buf)
		reverse(buf[:len(buf)-1])
		reverse(buf)
		return toNum(buf)
	}

	var best int
	best = max(rrot(lrot(n)), best)
	best = max(lrot(rrot(n)), best)

	for i, cur := 0, n; i < 6; i++ {
		cur = lrot(cur)
		best = max(best, cur)
	}

	for i, cur := 0, n; i < 6; i++ {
		cur = rrot(cur)
		best = max(best, cur)
	}

	return best
}

func toNum(arr []int) int {
	var res int
	for i := 0; i < len(arr); i++ {
		res = res*10 + arr[i]
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
