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
	t := readString(reader)
	res := solve(s, t)

	fmt.Println(res)
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

const mod = 1000000007

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func solve(s string, x string) int {
	n := len(s)

	res := 1
	for i := 0; i < n; i++ {
		if s[i] == '?' {
			res = mul(res, 10)
		}
		if x[i] == '?' {
			res = mul(res, 10)
		}
	}

	res = sub(res, countLessWays(s, x))
	res = sub(res, countLessWays(x, s))
	res = add(res, countEqualWays(s, x))

	return res
}

func countLessWays(s string, x string) int {
	n := len(s)
	res := 1
	for i := 0; i < n; i++ {
		// we want all s[i] <= x[i]
		if s[i] != '?' && x[i] != '?' {
			if s[i] > x[i] {
				return 0
			}
			continue
		}
		if s[i] == '?' && x[i] == '?' {
			res = mul(res, 55)
		} else if s[i] == '?' {
			u := int(x[i] - '0')
			res = mul(res, u+1)
		} else if x[i] == '?' {
			u := int(s[i] - '0')
			res = mul(res, 9-u+1)
		}
	}

	return res
}

func countEqualWays(s string, x string) int {
	n := len(s)
	res := 1
	for i := 0; i < n; i++ {
		if s[i] != '?' && x[i] != '?' {
			if s[i] != x[i] {
				return 0
			}
			continue
		}
		if s[i] == '?' && x[i] == '?' {
			res = mul(res, 10)
		}
	}

	return res
}
