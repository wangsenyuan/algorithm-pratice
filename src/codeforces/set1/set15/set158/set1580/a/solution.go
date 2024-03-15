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
		n, m := readTwoNums(reader)

		mat := make([]string, n)

		for i := 0; i < n; i++ {
			mat[i] = readString(reader)[:m]
		}

		res := solve(mat)

		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(mat []string) int {
	n := len(mat)
	m := len(mat[0])
	if n < m {
		return solve1(mat, 4, 3)
	}
	mat = transpose(mat)

	return solve1(mat, 3, 4)
}

func transpose(mat []string) []string {
	n, m := len(mat), len(mat[0])
	buf := make([][]byte, m)
	for i := 0; i < m; i++ {
		buf[i] = make([]byte, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			buf[j][i] = mat[i][j]
		}
	}
	res := make([]string, m)
	for i := 0; i < m; i++ {
		res[i] = string(buf[i])
	}
	return res
}

func solve1(mat []string, a int, b int) int {
	n := len(mat)
	m := len(mat[0])

	sum := make([][]int, m)
	for i := 0; i < m; i++ {
		sum[i] = make([]int, n)
	}

	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			if mat[i][j] == '1' {
				sum[j][i]++
			}
			if i > 0 {
				sum[j][i] += sum[j][i-1]
			}
		}
	}

	best := 20

	pref := make([]int, m+1)
	dp := make([]int, 21)
	rec := make([]int, m)
	process := func(top int, bot int) {
		// top,bot都变成0，但是中间都要变成0
		pref[0] = 0

		for j := 0; j < m; j++ {
			var tmp int
			if mat[top][j] == '0' {
				// need to change to 1
				tmp++
			}
			ones := sum[j][bot-1] - sum[j][top]
			// mid change from 1 to 0
			tmp += ones
			if mat[bot][j] == '0' {
				tmp++
			}
			pref[j+1] = pref[j] + tmp
			// change 0 to 1
			rec[j] = bot - top - 1 - ones
		}
		for j := 0; j <= 20; j++ {
			dp[j] = -1
		}
		// 还是不够
		for j := b; j < m; j++ {
			if rec[j-b] < 20 {
				// 这里还有个要求是 j - l + 1 >= 4
				dp[rec[j-b]] = j - b
			}
			// 如果以j为有边界
			// 这个就是需要计算的部分
			if rec[j] >= best {
				// already worse
				continue
			}
			for k := 0; k < 20-rec[j]; k++ {
				// 修改了k次的列， 每次修改的都记录最近的即可
				l := dp[k]
				if l >= 0 {
					// 中间的变化， 不包括第j，l列
					best = min(best, rec[j]+k+pref[j]-pref[l+1])
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := i + a; j < n; j++ {
			process(i, j)
		}
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
