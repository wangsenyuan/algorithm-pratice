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
		var i int
		for i < len(s) && s[i] != ' ' {
			i++
		}
		start := s[:i]
		i++
		for i < len(s) && s[i] != ' ' {
			i++
		}
		end := s[len(start)+1 : i]
		var l, r int
		pos := readInt([]byte(s), i+1, &l)
		readInt([]byte(s), pos+1, &r)
		res := solve(start, end, l, r)
		buf.WriteString(res)
		buf.WriteByte('\n')
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

var days = []string{"saturday", "sunday", "monday", "tuesday", "wednesday", "thursday", "friday"}

func solve(start string, end string, l int, r int) string {
	// let end - start = 7 * x + diff
	var diff int
	for i := 0; i < 7; i++ {
		if days[i] == start {
			for j := 0; j < 7; j++ {
				diff++
				if days[(i+j)%7] == end {
					break
				}
			}
			break
		}
	}
	var count int
	var days int
	for diff <= r {
		if diff >= l {
			days = diff
			count++
		}
		diff += 7
	}

	if count == 1 {
		return fmt.Sprintf("%d", days)
	}
	if count == 0 {
		return "impossible"
	}
	return "many"
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
