package main

import (
	"bufio"
	"bytes"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res := process(reader)

	var buf bytes.Buffer
	for _, x := range res {
		if x {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func process(reader *bufio.Reader) []bool {
	n, q := readTwoNums(reader)
	word := readString(reader)[:n]
	cmds := make([]string, q)
	for i := 0; i < q; i++ {
		cmds[i] = readString(reader)
	}
	return solve(word, cmds)
}

func solve(word string, cmds []string) []bool {
	n := len(word)
	next := make([][]int, n+2)
	for i := range n + 2 {
		next[i] = make([]int, 26)
		for j := range 26 {
			next[i][j] = n
		}
	}
	for i := n - 1; i >= 0; i-- {
		copy(next[i], next[i+1])
		x := int(word[i] - 'a')
		next[i][x] = i
	}

	dp := make([][][]int, 255)
	for i := range 255 {
		dp[i] = make([][]int, 255)
		for j := range 255 {
			dp[i][j] = make([]int, 255)
			for k := range 255 {
				dp[i][j][k] = n
			}
		}
	}

	dp[0][0][0] = -1

	s := make([][]byte, 3)

	ans := make([]bool, len(cmds))

	ll := make([]int, 3)

	add := func(i int, c byte) {
		s[i] = append(s[i], c)
		// 先设置最大值
		clear(ll)
		ll[i] = len(s[i]) - 1
		// 为啥这里要重新计算一遍呢？
		for l0 := ll[0]; l0 <= len(s[0]); l0++ {
			for l1 := ll[1]; l1 <= len(s[1]); l1++ {
				for l2 := ll[2]; l2 <= len(s[2]); l2++ {
					if l0+l1+l2 == 0 {
						continue
					}
					dp[l0][l1][l2] = n
					if l0 > 0 {
						dp[l0][l1][l2] = min(dp[l0][l1][l2], next[dp[l0-1][l1][l2]+1][int(s[0][l0-1]-'a')])
					}
					if l1 > 0 {
						dp[l0][l1][l2] = min(dp[l0][l1][l2], next[dp[l0][l1-1][l2]+1][int(s[1][l1-1]-'a')])
					}
					if l2 > 0 {
						dp[l0][l1][l2] = min(dp[l0][l1][l2], next[dp[l0][l1][l2-1]+1][int(s[2][l2-1]-'a')])
					}
				}
			}
		}
	}

	for i, cur := range cmds {
		j := int(cur[2] - '1')
		if cur[0] == '+' {
			c := cur[4]
			add(j, c)
		} else {
			s[j] = s[j][:len(s[j])-1]
		}
		ans[i] = dp[len(s[0])][len(s[1])][len(s[2])] < n
	}

	return ans
}
