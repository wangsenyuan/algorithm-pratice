package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	A := readNNums(reader, n)
	z, res := solve(A)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", z))
	buf.WriteString(res)
	buf.WriteByte('\n')
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

func solve(A []int) (int, string) {
	// concat(a, b) * concat(b, a) % 3
	// 是否等于 a * b % 3?
	// 似乎不是的
	// (a * 100 + b) * (b * 10 + a)
	// 如果将a和b替换成 a % 3, b % 3 (但是保留长度）
	// (a * 100 + b) * (b * 10 + a) 和变化前mod 3的结果是一致的
	// 这时候长度也没有关系了，因为 b * a 最多到4，不会长生进位
	// 在跟定z的情况下， 可以判定i和j是否有冲突

	cnt := make([]int, 2)
	n := len(A)
	for i := 0; i < n; i++ {
		x := A[i] % 3
		x = x * x % 3
		// 0, or 1
		cnt[x]++
	}
	var z int
	res := make([]byte, n)
	if cnt[0] >= n/2 {
		z = 2
		m := n / 2
		for i := 0; i < n; i++ {
			x := A[i] % 3
			x = x * x % 3
			if x == 0 && m > 0 {
				res[i] = '0'
				m--
			} else {
				res[i] = '1'
			}
		}
	} else {
		z = 0
		m := n / 2
		for i := 0; i < n; i++ {
			x := A[i] % 3
			x = x * x % 3
			if x == 1 && m > 0 {
				res[i] = '1'
				m--
			} else {
				res[i] = '0'
			}
		}
	}

	return z, string(res)
}
