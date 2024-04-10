package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	readNum(reader)
	s := readString(reader)
	m := readNum(reader)
	words := make([]string, m)
	for i := 0; i < m; i++ {
		words[i] = readString(reader)
	}

	res := solve(s, words)

	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func solve(s string, words []string) int {
	n := len(s)

	var pos int

	var flag int
	for i := 0; i < n; i++ {
		if s[i] != '*' {
			flag |= 1 << int(s[i]-'a')
			pos |= 1 << i
		}
	}

	var valid []string

	for _, w := range words {
		ok := true
		for i := 0; i < n; i++ {
			if (pos>>i)&1 == 1 {
				if w[i] != s[i] {
					ok = false
					break
				}
			} else {
				x := int(w[i] - 'a')
				if (flag>>x)&1 == 1 {
					ok = false
					break
				}
			}
		}
		if ok {
			valid = append(valid, w)
		}
	}

	check := func(x int) bool {
		for _, w := range valid {
			ok := false
			for i := 0; i < n; i++ {
				if (pos>>i)&1 == 1 {
					// already revealed
					continue
				}
				if int(w[i]-'a') == x {
					ok = true
					break
				}
			}
			if !ok {
				return false
			}
		}
		return true
	}

	var res int
	for x := 0; x < 26; x++ {
		if (flag>>x)&1 == 0 && check(x) {
			res++
		}
	}

	return res
}
