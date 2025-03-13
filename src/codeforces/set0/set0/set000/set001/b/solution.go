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
	for range tc {
		s := readString(reader)
		r := solve(s)
		buf.WriteString(r)
		buf.WriteByte('\n')
	}
	buf.WriteTo(os.Stdout)
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(s string) string {
	if s[0] == 'R' {
		i := 1
		for i < len(s) && isDigit(s[i]) {
			i++
		}
		if i < len(s) && i > 1 && s[i] == 'C' {
			j := i + 1
			for j < len(s) && isDigit(s[j]) {
				j++
			}
			if j == len(s) {
				// it is the RXCY pattern
				//x := parseAsNum(s[1:i])
				y := parseAsNum(s[i+1:])
				return convertToLetters(y) + s[1:i]
			}
		}
	}
	//it is BC23 pattern
	var num int
	var i int
	for i < len(s) && isUpper(s[i]) {
		num = num*26 + int(s[i]-'A')
		num++
		i++
	}
	//num++
	return "R" + s[i:] + "C" + fmt.Sprintf("%d", num)
}

func parseAsNum(s string) int {
	var res int
	for i := 0; i < len(s); i++ {
		res = res*10 + int(s[i]-'0')
	}
	return res
}

func convertToLetters(num int) string {
	var buf bytes.Buffer
	for num > 0 {
		num--
		x := byte(num%26 + 'A')
		buf.WriteByte(x)
		num = num / 26
	}
	bs := buf.Bytes()
	reverse(bs)
	return string(bs)
}

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func isUpper(x byte) bool {
	return x >= 'A' && x <= 'Z'
}

func isDigit(x byte) bool {
	return x >= '0' && x <= '9'
}
