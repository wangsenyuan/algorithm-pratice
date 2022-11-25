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
		s := readString(reader)
		res := solve(s)
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

func solve(s string) int {
	// cnt[0] = 01
	// cnt[1] = 10
	cnt := make([]int, 2)

	n := len(s)

	for i := 1; i < n; i++ {
		if s[i-1] == '0' && s[i] == '1' {
			cnt[0]++
		}
		if s[i-1] == '1' && s[i] == '0' {
			cnt[1]++
		}
	}

	var res int
	for i := 0; i < n; i++ {
		// flip s[i]
		a, b := cnt[0], cnt[1]
		if i > 0 {
			if s[i-1] == '0' && s[i] == '1' {
				// 00
				a--
			}
			if s[i-1] == '1' && s[i] == '0' {
				b--
			}
			if s[i-1] == '0' && s[i] == '0' {
				a++
			}
			if s[i-1] == '1' && s[i] == '1' {
				b++
			}
		}
		if i+1 < n {
			if s[i] == '0' && s[i+1] == '1' {
				a--
			}
			if s[i] == '1' && s[i+1] == '0' {
				b--
			}
			if s[i] == '0' && s[i+1] == '0' {
				b++
			}
			if s[i] == '1' && s[i+1] == '1' {
				a++
			}
		}
		if a == b {
			res++
		}
	}
	return res
}
