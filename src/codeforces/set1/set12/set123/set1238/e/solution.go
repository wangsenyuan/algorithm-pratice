package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	//tc := readNum(scanner)
	tc := 1

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		s := readString(reader)[:n]
		res := solve(s, m)

		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Println(buf.String())
}

const inf = 1 << 31

func solve(s string, m int) int {
	n := len(s)
	// 考虑ab的出现的次数是x， 如果在perm中, dist(pos[a], pos[b])的距离是y
	// 那么它的贡献就是 x * dist()
	contr := make([][]int, 26)
	for i := 0; i < 26; i++ {
		contr[i] = make([]int, 26)
	}

	for i := 0; i+1 < n; i++ {
		a := int(s[i] - 'a')
		b := int(s[i+1] - 'a')
		contr[a][b]++
		contr[b][a]++
	}

	// 考虑在排好 mask的情况下，再加入i时的情况
	// dp[mask][i] 表示排好mask的字母后，且以i为最后一个字母时的最小cost
	M := 1 << m

	mb := make([]int, M)
	cnt := make([]int, M)
	for i := 1; i < M; i++ {
		cnt[i] = 1 + cnt[i&(i-1)]
		for j := 0; j < m; j++ {
			if (i>>j)&1 == 1 {
				mb[i] = j
				break
			}
		}
	}
	fp := make([][]int, M)
	for state := 0; state < M; state++ {
		fp[state] = make([]int, m)
	}

	for state := 0; state < M; state++ {
		for i := 0; i < m; i++ {
			j := mb[state]
			fp[state][i] = fp[state^(1<<j)][i] + contr[i][j]
		}
	}

	dp := make([]int, M)
	for i := 0; i < M; i++ {
		dp[i] = inf
	}
	dp[0] = 0

	for state := 0; state < M; state++ {
		for i := 0; i < m; i++ {
			if (state>>i)&1 == 1 {
				continue
			}
			pos := cnt[state]
			next := state | (1 << i)
			dp[next] = min(dp[next], dp[state]+pos*(fp[state][i]-fp[(M-1)^next][i]))
		}
	}

	return dp[M-1]
}

func digit_count(num int) int {
	var res int
	for num > 0 {
		res++
		num -= num & -num
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
