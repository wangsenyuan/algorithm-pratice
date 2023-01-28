package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line := readString(reader)
	res := solve(line)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
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

var valid_endings map[string]int

func init() {
	valid_endings = make(map[string]int)
	valid_endings["lios"] = 0
	valid_endings["etr"] = 1
	valid_endings["initis"] = 2
	valid_endings["liala"] = 3
	valid_endings["etra"] = 4
	valid_endings["inites"] = 5
}

func solve(line string) bool {
	words := strings.Split(line, " ")
	gender := check(words[0])
	if gender < 0 {
		return false
	}
	if len(words) == 1 {
		return true
	}
	gender /= 3
	// 0 for masculine, 1 for feminine
	var i int
	for i < len(words) {
		x := check(words[i])
		if x/3 != gender {
			return false
		}
		x %= 3
		if x == 0 {
			i++
		} else {
			break
		}
	}
	if i == len(words) {
		// all adj
		return false
	}
	if check(words[i])%3 != 1 {
		// not noun
		return false
	}
	i++
	for i < len(words) {
		x := check(words[i])
		if x/3 != gender || x%3 != 2 {
			return false
		}
		i++
	}
	return true
}

func check(word string) int {
	for end, v := range valid_endings {
		if strings.HasSuffix(word, end) {
			return v
		}
	}
	return -1
}
