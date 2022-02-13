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
		s := readString(reader)
		k := readNum(reader)
		res := solve(n, s, k)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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

func solve(n int, s string, m int) int {
	//if start from position i, and ends at j (greedily), get a substring as B
	//then the anser is (j - i) - m
	// try find dp[i] = j?
	v := make([][]int, 26)
	for i := 0; i < 26; i++ {
		v[i] = make([]int, 0, 2)
	}

	for i := 0; i < n; i++ {
		x := int(s[i] - 'a')
		v[x] = append(v[x], i)
	}

	itr := make([]int, 26)
	next := make([]int, n)
	for i := 0; i < n; i++ {
		x := int(s[i] - 'a')
		j := (x + 1) % 26
		for {
			if len(v[j]) == itr[j] {
				next[i] = -1
				break
			}
			if v[j][itr[j]] <= i {
				itr[j]++
			} else {
				next[i] = v[j][itr[j]]
				break
			}
		}
	}
	dp := make([][]int, n)

	for i := n - 1; i >= 0; i-- {
		dp[i] = make([]int, 20)
		dp[i][0] = next[i]
		for j := 1; j < 20; j++ {
			if dp[i][j-1] >= 0 {
				dp[i][j] = dp[dp[i][j-1]][j-1]
			} else {
				dp[i][j] = -1
			}
		}
	}

	ans := 1 << 30

	for i := 0; i < n; i++ {
		if s[i] == 'a' {
			st := i
			k := m - 1
			for j := 0; j < 20; j++ {
				if k&1 == 1 {
					st = dp[st][j]
				}
				if st == -1 {
					break
				}
				k >>= 1
			}

			if st != -1 {
				ans = min(ans, st-i+1)
			}
		}
	}

	if ans == 1<<30 {
		return -1
	}
	return ans - m

}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
