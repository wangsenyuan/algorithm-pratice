package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	statements := make([]string, n)
	for i := 0; i < n; i++ {
		statements[i] = readString(reader)
	}

	fmt.Println(solve(statements))
}

func readFloat64(bs []byte, pos int, r *float64) int {
	var first float64
	var second float64
	exp := 1.0
	var dot bool
	for pos < len(bs) && (bs[pos] == '.' || isDigit(bs[pos])) {
		if bs[pos] == '.' {
			dot = true
		} else if !dot {
			first = first*10 + float64(bs[pos]-'0')
		} else {
			second = second*10 + float64(bs[pos]-'0')
			exp *= 10
		}
		pos++
	}
	*r = first + second/exp
	//fmt.Fprintf(os.Stderr, "%.6f ", *r)
	return pos
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
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

func solve(statement []string) string {
	n := len(statement)

	var p int
	for i := 0; i < n; i++ {
		s := statement[i]
		s = strings.TrimSpace(s)
		if len(s) == 0 {
			continue
		}
		statement[p] = s
		p++
	}
	n = p

	R := make([]int, n)
	stack := make([]int, n)
	p = 0
	var throwAt int
	for i := 0; i < n; i++ {
		x := strings.ToLower(statement[i])
		if x == "try" {
			stack[p] = i
			p++
			continue
		}
		if strings.HasPrefix(x, "catch") {
			R[stack[p-1]] = i
			p--
		} else if strings.HasPrefix(x, "throw") {
			throwAt = i
		}
	}

	exception := getArguments(statement[throwAt])

	for i := throwAt - 1; i >= 0; i-- {
		x := strings.ToLower(statement[i])
		if x == "try" {
			j := R[i]
			if j < throwAt {
				continue
			}
			// j > throwAt
			argStr := getArguments(statement[j])
			args := strings.Split(argStr, ",")
			c := strings.TrimSpace(args[0])
			if c == exception {
				// good find it
				return strings.Trim(strings.TrimSpace(args[1]), "\"")
			}
		}
	}

	return "Unhandled Exception"
}

func getArguments(s string) string {
	i := strings.IndexByte(s, '(')
	j := strings.IndexByte(s, ')')
	x := s[i+1 : j]
	return strings.TrimSpace(x)
}
