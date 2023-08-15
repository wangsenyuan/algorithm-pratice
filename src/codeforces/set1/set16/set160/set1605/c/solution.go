package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		s := readString(reader)[:n]
		res := solve(s)
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

func solve(s string) int {
	// 如果存在aa，=> 2
	// 假设存在一段 l, r  cnt[a] > cnt[b] & cnt[a] > cnt[c]
	// cnt[a] + cnt[b] + cnt[c] = r - l + 1
	// 因为不存在aa, 所以要么ab,要么ac
	// 如果出现了连续的bc，可以直接跳过它们，这是因为，bc，至少需要两个a，如果这两个a在同一侧，它们要么给出了更好的答案
	n := len(s)
	for i := 0; i+1 < n; i++ {
		if s[i] == 'a' && s[i+1] == 'a' {
			return 2
		}
	}
	// 然后检查3个长度
	for i := 0; i+2 < n; i++ {
		if s[i] == 'a' && s[i+2] == 'a' {
			return 3
		}
	}
	// 然后检查4个长度
	for i := 0; i+3 < n; i++ {
		if s[i] == 'a' && s[i+3] == 'a' {
			if s[i+1] != s[i+2] {
				return 4
			}
		}
	}
	// abbabbaccacca
	for i := 0; i+7 <= n; i++ {
		if s[i] == 'a' && s[i+3] == 'a' && s[i+6] == 'a' && s[i+1] != s[i+4] {
			return 7
		}
	}
	return -1
}
