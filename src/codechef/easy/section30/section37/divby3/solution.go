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

const INF = 1 << 30

func solve(s string) int {
	n := len(s)
	var cnt int
	var sum int
	for i := 0; i < n; i++ {
		sum *= 2
		if s[i] == '1' {
			cnt++
			sum++
		}
		sum %= 3
	}

	if sum == 0 {
		return 0
	}

	if cnt == n {
		return -1
	}

	cnt0 := n - cnt

	if min(cnt0, cnt) == 1 {
		for i := 0; i < n-1; i++ {
			if s[i] != s[i+1] {
				var pos int
				if s[i] == '1' {
					pos = n - i
				} else {
					pos = n - (i + 1)
				}
				pos %= 2
				if sum == 2 && pos == 1 {
					return 1
				}
				if sum == 1 && pos == 0 {
					return 1
				}
				return -1
			}
		}
	}

	var positions []int

	mark := make([]bool, 2)

	for i := 0; i < n-1; i++ {
		if s[i] != s[i+1] {
			if s[i] == '1' {
				positions = append(positions, (n-i)%2)
			} else {
				positions = append(positions, (n-(i+1))%2)
			}
		}
	}
	for _, it := range positions {
		mark[it] = true
	}

	if len(positions) == 1 {
		if sum == 2 && positions[0] == 1 {
			return 1
		}
		if sum == 1 && positions[0] == 0 {
			return 1
		}
		return 3
	}

	if sum == 1 {
		if mark[0] {
			return 1
		}
		return 2
	}

	if mark[1] {
		return 1
	}
	return 2
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
