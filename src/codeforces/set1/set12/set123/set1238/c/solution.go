package main

import (
	"bufio"
	"bytes"
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
	var buf bytes.Buffer

	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		h, n := readTwoNums(reader)
		a := readNNums(reader, n)
		res := solve(h, a)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

const inf = 1 << 30

func solve(h int, p []int) int {
	// 下落的过程，一层层下来，在中间的某些位置x的下一层(x - 1)
	// 处存在一个挡板，所以，可以额外花费一个宝石在x-2处出现挡板
	// 越过这个位置
	// 所以必须考虑某个p[i]的前一个位置，需要花费多少
	n := len(p)

	var ans int

	var l int

	for i := range n {
		if i > 0 && p[i-1]-1 > p[i] {
			if l > 0 {
				ans += (i - l) & 1
			} else {
				ans += 1 - (i-l)&1
			}
			l = i
		}
	}

	if p[n-1] > 1 {
		if l > 0 {
			ans += (n - l) & 1
		} else {
			ans += 1 - (n-l)&1
		}
	}

	return ans
}
