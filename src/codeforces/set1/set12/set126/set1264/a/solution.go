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
		p := readNNums(reader, n)
		res := solve(p)
		buf.WriteString(fmt.Sprintf("%d %d %d\n", res[0], res[1], res[2]))
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
func solve(p []int) []int {
	n := len(p)
	if n <= 3 {
		// g > 0, s > 0, g > 0
		// g < s && g < b 肯定达不到
		return []int{0, 0, 0}
	}

	// n = 7, 那么前面只有能3个，所以要从2开始算
	r := n/2 - 1

	for r >= 0 && p[r] == p[r+1] {
		r--
	}
	// r是铜牌和无牌的分界线
	if r < 3 {
		return []int{0, 0, 0}
	}

	prev := make([]int, r+1)

	for i := 1; i <= r; i++ {
		prev[i] = prev[i-1]
		if p[i-1] > p[i] {
			prev[i] = i - 1
		}
	}

	for i := 0; i < n/2; i++ {
		if p[i] > p[i+1] {
			g := i + 1
			// s > g, b > g
			// and g + s + b <= r + 1
			// 那么在 2 * g + 1 到 r - (g + 1) 中间找到一个 j, p[j] > p[j+1]
			b := r - (g + 1) + 1
			if b <= i {
				break
			}
			// 至少保证有 g + 1 个 铜牌
			j := prev[b]
			// j是银牌和铜牌的分界线
			b = r - j
			s := r + 1 - g - b
			if s > g {
				return []int{g, s, b}
			}
		}
	}
	return []int{0, 0, 0}
}
