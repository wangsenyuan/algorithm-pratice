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
		_, k := readTwoNums(reader)
		s := readString(reader)
		res := solve(s, k)
		buf.WriteString(fmt.Sprintf("%s\n", res))
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
		if s[i] == '\n' || s[i] == '\r' {
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

const INF = 1 << 30

func solve(s string, k int) string {
	buf := []byte(s)
	n := len(buf)
	var cnt int
	for i := 0; i < n; i++ {
		cnt += int(buf[i] - '0')
	}
	if cnt == 0 {
		// all 0, no change
		return s
	}

	if cnt == n {
		return all_zero(n)
	}

	for i := 0; i < n; i++ {
		if s[i] == '1' {
			buf[i] = '0'
			if i > 0 && s[i-1] == '0' {
				buf[i-1] = '1'
			}
			if i+1 < n && s[i+1] == '0' {
				buf[i+1] = '1'
			}
		}
	}

	k--

	if k == 0 {
		return string(buf)
	}

	left := make([]int, n)

	for i := 0; i < n; i++ {
		if buf[i] == '1' {
			left[i] = i
		} else if i > 0 {
			left[i] = left[i-1]
		} else {
			left[i] = -1
		}
	}

	right := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		if buf[i] == '1' {
			right[i] = i
		} else if i+1 < n {
			right[i] = right[i+1]
		} else {
			right[i] = n
		}
	}

	for i := 0; i < n; i++ {
		first := INF
		if buf[i] == '1' {
			first = 0
		} else {
			if left[i] >= 0 {
				first = i - left[i]
			}
			if right[i] < n {
				first = min(first, right[i]-i)
			}
		}
		if first > k {
			buf[i] = '0'
		} else {
			if (k-first)&1 == 0 {
				buf[i] = '1'
			} else {
				buf[i] = '0'
			}
		}
	}

	return string(buf)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func all_zero(n int) string {
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = '0'
	}
	return string(buf)
}
