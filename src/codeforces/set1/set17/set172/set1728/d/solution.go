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

	tc := readNum(reader)

	for tc > 0 {
		tc--
		s := readString(reader)
		res := solve(s)
		buf.WriteString(fmt.Sprintf("%s\n", res))
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

func solve(s string) string {
	n := len(s)
	// n is even and n <= 2000
	// 假设当前用户时alice, 处理l...r
	// 如果 s[l] < s[r], 且dp[l+1...r] 是draw, 那么 alice 胜
	//    或者 s[l] > s[r] 但是 dp[l, r-1]是draw， 那么也是alice 胜
	//  如果 s[l] == s[r] 且 dp[l+1..r]和 dp[l...r-1]都是draw，那么就是draw
	// 否则 alice 败
	// 如果当前是bob, 也可以有同样当然分析
	// 是不是以一对进行分析更合理？
	// 如果只剩下一个的时候，貌似没法判断呐.

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}
	fp := make([][][]int, n)
	for i := 0; i < n; i++ {
		fp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			fp[i][j] = make([]int, 2)
			for k := 0; k < 2; k++ {
				fp[i][j][k] = -1
			}
		}
	}
	// 0 for draw, 1 for win, 2 for lose
	var alice func(l int, r int) int
	var bob func(l int, r int, a int) int

	alice = func(l int, r int) int {
		if l > r {
			return 0
		}
		// l < r always true
		if dp[l][r] < 0 {
			tmp1 := bob(l+1, r, 0)
			if tmp1 == 2 {
				dp[l][r] = 1
				return 1
			}
			tmp2 := bob(l, r-1, 1)
			if tmp2 == 2 {
				dp[l][r] = 1
				return 1
			}
			if tmp1 == 0 || tmp2 == 0 {
				dp[l][r] = 0
			} else {
				dp[l][r] = 2
			}
		}

		return dp[l][r]
	}
	bob = func(l int, r int, a int) int {
		if fp[l][r][a] < 0 {
			var x byte
			if a == 0 {
				// left of l
				x = s[l-1]
			} else {
				// right of r
				x = s[r+1]
			}
			// try win
			tmp1 := alice(l+1, r)
			if tmp1 == 2 || tmp1 == 0 && x > s[l] {
				// alice lose
				fp[l][r][a] = 1
				return 1
			}
			tmp2 := alice(l, r-1)
			if tmp2 == 2 || tmp2 == 0 && x > s[r] {
				fp[l][r][a] = 1
				return 1
			}
			// can't win, try draw
			if tmp1 == 0 && x == s[l] || tmp2 == 0 && x == s[r] {
				fp[l][r][a] = 0
				return 0
			}
			// lose
			fp[l][r][a] = 2
		}

		return fp[l][r][a]
	}

	ans := alice(0, n-1)
	if ans == 1 {
		return "Alice"
	}

	if ans == 2 {
		return "Bob"
	}

	return "Draw"
}
