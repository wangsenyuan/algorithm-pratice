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
		s, _ := reader.ReadString('\n')
		res := solve(len(s)-1, s)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
	}

	fmt.Print(buf.String())
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
func solve(n int, s string) bool {
	var one bool

	for i := 1; i < n; i++ {
		if s[i] == s[i-1] && s[i] == '0' {
			if one {
				return false
			}
		}
		one = one || s[i] == '1' && s[i-1] == '1'
	}
	return true
}
func solve1(n int, s string) bool {
	// if position i is kept
	// if s[i] = 0, then all '1' before it should be removed,
	// if s[i] = 1, then all '0' after it should be removed
	// let dp[i] means if there are continues 1 before i,
	dp := make([]bool, n)
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] || s[i] == '1' && s[i-1] == '1'
	}
	if !dp[n-1] {
		// remove all ones
		return true
	}
	// fp means there are continues 0 so far
	var fp bool
	for i := n - 2; i >= 0; i-- {
		if s[i] == '0' {
			if !dp[i] && !fp {
				// remove all ones before, and all zeros after
				return true
			}
		}
		fp = fp || s[i] == '0' && s[i+1] == '0'
	}
	if !fp {
		// return all zeros
		return true
	}
	return false
}
