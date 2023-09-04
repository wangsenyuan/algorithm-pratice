package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	t := readNum(reader)

	for t > 0 {
		t--
		n, m := readTwoNums(reader)
		A := readNNums(reader, n)
		Q := make([][]int, m)
		for i := 0; i < m; i++ {
			Q[i] = readNNums(reader, 2)
		}
		res := solve(A, Q)
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d\n", x))
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

func solve(A []int, Q [][]int) []int {
	// 先不考虑限制
	// 1号能够胜利，直到它遇到下一个比a[1]大的数字
	// 只需要考虑第一圈，直到遇到n，剩余的就后面胜利的只会是n了
	n := len(A)
	pos := make([]int, n+1)
	for i, a := range A {
		pos[a-1] = i
	}

	m := pos[n-1]
	winner := make([]int, m+1)
	winner[0] = 0
	L := make([]int, n)
	R := make([]int, n)
	for i := 0; i < n; i++ {
		L[i] = -1
		R[i] = -1
	}
	for i := 1; i < m; i++ {
		if A[winner[i-1]] < A[i] {
			winner[i] = i
		} else {
			winner[i] = winner[i-1]
		}
		if L[winner[i]] < 0 {
			L[winner[i]] = i
		}
		R[winner[i]] = i
	}

	ans := make([]int, len(Q))

	for j, cur := range Q {
		i, k := cur[0], cur[1]
		i--
		if i > m {
			//它在n的后面, will never win
			ans[j] = 0
			continue
		}
		if i == m {
			// i = n - 1
			if k < m {
				ans[j] = 0
			} else {
				if m == 0 {
					ans[j] = k
				} else {
					// 假设m=2，那么需要2次round后，才能到达n，在第二次时，n将获胜 (+1), 且在剩下的 k - m 次后一直获胜
					ans[j] = k - m + 1
				}
			}
			continue
		}
		// i < m
		if L[i] < 0 || k < L[i] {
			// no win at all
			ans[j] = 0
			continue
		}
		ans[j] = min(k, R[i]) - L[i] + 1
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
