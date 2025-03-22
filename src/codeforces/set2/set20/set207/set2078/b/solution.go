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
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	buf.WriteTo(os.Stdout)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
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

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	gates := make([]string, n)
	for i := 0; i < n; i++ {
		gates[i] = readString(reader)
	}
	return solve(n, gates)
}

func solve(n int, gates []string) int {

	read := func(s string) (op1 byte, first int, op2 byte, second int) {
		op1 = s[0]
		pos := readInt([]byte(s), 2, &first)
		op2 = s[pos+1]
		pos += 3
		readInt([]byte(s), pos, &second)
		return
	}

	next := make([]int, n+1)

	for i := n - 1; i >= 0; i-- {
		op1, first, op2, second := read(gates[i])
		next[i] = next[i+1]
		if op1 == '+' && op2 == '+' {
			continue
		}
		if op1 == 'x' && op2 == 'x' {
			if first > second {
				next[i] = 0
			} else if first < second {
				next[i] = 1
			}
		} else if op1 == 'x' {
			next[i] = 0
		} else if op2 == 'x' {
			next[i] = 1
		}
	}

	cur := []int{1, 1}

	for i := range n {
		op1, first, op2, second := read(gates[i])
		var more int
		if op1 == '+' {
			more += first
		} else {
			more += (first - 1) * cur[0]
		}
		if op2 == '+' {
			more += second
		} else {
			more += (second - 1) * cur[1]
		}
		if next[i+1] == 0 {
			cur[0] += more
		} else {
			cur[1] += more
		}
	}
	return cur[0] + cur[1]
}
