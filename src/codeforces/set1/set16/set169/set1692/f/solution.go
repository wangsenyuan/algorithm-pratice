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

	for tc > 0 {
		tc--
		n := readNum(reader)
		arr := readNNums(reader, n)
		res := solve(arr)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(arr []int) bool {
	dp := make([]int, 10)
	fp := make([]int, 10)
	for _, num := range arr {
		r := num % 10
		for x := 0; x < 10; x++ {
			if fp[x] > 0 && (x+r)%10 == 3 {
				return true
			}
		}
		for x := 0; x < 10; x++ {
			fp[(x+r)%10] += dp[x]
		}

		dp[r]++
	}

	return false
}

func solve1(arr []int) bool {
	freq := make([]int, 10)

	for _, num := range arr {
		freq[num%10]++
	}

	for a := 0; a < 10; a++ {
		if freq[a] == 0 {
			continue
		}
		freq[a]--
		for b := a; b < 10; b++ {
			if freq[b] == 0 {
				continue
			}
			freq[b]--

			for c := b; c < 10; c++ {
				if freq[c] == 0 {
					continue
				}
				r := (a + b + c) % 10
				if r == 3 {
					return true
				}
			}

			freq[b]++
		}

		freq[a]++
	}

	return false
}
