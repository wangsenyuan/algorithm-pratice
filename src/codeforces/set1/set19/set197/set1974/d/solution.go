package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for range tc {
		readNum(reader)
		s := readString(reader)
		res := solve(s)
		buf.WriteString(res)
		buf.WriteByte('\n')
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

const action = "NSEW"

var vals = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func id(x byte) int {
	return strings.IndexByte(action, x)
}

func solve(s string) string {

	var x, y int

	for _, c := range []byte(s) {
		i := id(c)
		x += vals[i][0]
		y += vals[i][1]
	}

	if abs(x)%2 == 1 || abs(y)%2 == 1 {
		return "NO"
	}
	n := len(s)
	res := make([]byte, n)
	for i := range n {
		res[i] = 'R'
	}
	if abs(x)+abs(y) == 0 {
		if n == 2 {
			return "NO"
		}
		res[0] = 'H'
		first := id(s[0])
		for i := 1; i < n; i++ {
			second := id(s[i])
			if first^1 == second {
				res[i] = 'H'
				break
			}
		}
		return string(res)
	}

	if x != 0 {
		x /= 2
		for j, c := range []byte(s) {
			i := id(c)
			if x > 0 && i == 2 {
				res[j] = 'H'
				x--
			} else if x < 0 && i == 3 {
				res[j] = 'H'
				x++
			}

			if x == 0 {
				break
			}
		}
	}

	if y != 0 {
		y /= 2
		// y != 0
		for j, c := range []byte(s) {
			i := id(c)
			if y > 0 && i == 0 {
				res[j] = 'H'
				y--
			} else if y < 0 && i == 1 {
				res[j] = 'H'
				y++
			}
			if y == 0 {
				break
			}
		}
	}

	return string(res)
}

func abs(num int) int {
	return max(num, -num)
}
