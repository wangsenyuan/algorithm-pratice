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
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(a []int) int {
	// 把0变成1后，它前面的1和原值0造成的inversions将取消，同时产生新的1和0的inversions
	// 把1变成0后，它前面的1和现值0将产生新的inversions，且原来和后面的0造成inversions建取消
	n := len(a)
	// 前面有多少个1
	l1 := make([]int, n)
	// 后面有多少个0
	r0 := make([]int, n)

	for i := 0; i < n; i++ {
		if i > 0 {
			l1[i] = l1[i-1]
		}
		if a[i] == 1 {
			l1[i]++
		}
	}

	for i := n - 1; i >= 0; i-- {
		if i+1 < n {
			r0[i] = r0[i+1]
		}
		if a[i] == 0 {
			r0[i]++
		}
	}

	var ans int
	var best int

	for i := 0; i < n; i++ {
		if a[i] == 0 {
			// l1[i] = l1[i-1]
			ans += l1[i]
			// 只需要加一次，不能加两次
		}
		var tmp int
		// flip it
		if a[i] == 0 {
			tmp -= l1[i]
			tmp += r0[i] - 1
		} else {
			tmp -= r0[i]
			tmp += l1[i] - 1
		}
		if tmp > best {
			best = tmp
		}
	}

	return ans + best
}
