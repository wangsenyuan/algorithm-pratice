package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	p := make([]string, n)
	for i := 0; i < n; i++ {
		p[i] = readString(reader)
	}
	w := readString(reader)
	lucky := readString(reader)

	res := solve(p, w, lucky)

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

func solve(p []string, w string, lucky string) string {
	// 把lucky也放进去， 遇到lucky， 就当作一次replace

	slices.SortFunc(p, func(a, b string) int {
		return len(a) - len(b)
	})

	// dp[i] 表示能够替换的最大值
	n := len(w)

	check := func(s string, x string) bool {
		for i := 0; i < len(s); i++ {
			if upperCase(s[i]) != upperCase(x[i]) {
				return false
			}
		}
		return true
	}

	diff := make([]int, n+1)

	for i := 0; i < n; i++ {
		for _, x := range p {
			if i+len(x) > n {
				break
			}
			if check(w[i:i+len(x)], x) {
				diff[i]++
				diff[i+len(x)]--
			}
		}
	}

	for i := 1; i < n; i++ {
		diff[i] += diff[i-1]
	}

	res := []byte(w)

	a := lowerCase(lucky[0])
	A := upperCase(lucky[0])

	for i := 0; i < n; i++ {
		if diff[i] > 0 {
			if lowerCase(res[i]) == a {
				// must change
				if res[i] >= 'a' && res[i] <= 'z' {
					if a == 'a' {
						res[i] = 'b'
					} else {
						res[i] = 'a'
					}
				} else {
					if A == 'A' {
						res[i] = 'B'
					} else {
						res[i] = 'A'
					}
				}
			} else if res[i] >= 'a' && res[i] <= 'z' {
				res[i] = a
			} else {
				res[i] = A
			}
		}
	}

	return string(res)
}

func lowerCase(x byte) byte {
	if x >= 'a' && x <= 'z' {
		return x
	}
	return byte(x - 'A' + 'a')
}

func upperCase(x byte) byte {
	if x >= 'A' && x <= 'Z' {
		return x
	}
	return byte(x - 'a' + 'A')
}
