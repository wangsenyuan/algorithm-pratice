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
		n, m := readTwoNums(reader)
		mat := make([][]int, n)
		for i := 0; i < n; i++ {
			mat[i] = readNNums(reader, m)
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

const inf = 1 << 20

func solve(mat [][]int) int {
	n := len(mat)
	m := len(mat[0])
	sum := make([][]int, n)
	for i := 0; i < n; i++ {
		sum[i] = make([]int, m)
	}

	check := func(l int) bool {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if mat[i][j] >= l {
					sum[i][j] = 1
				} else {
					sum[i][j] = 0
				}
				if i > 0 {
					sum[i][j] += sum[i-1][j]
				}
				if j > 0 {
					sum[i][j] += sum[i][j-1]
				}
				if i > 0 && j > 0 {
					sum[i][j] -= sum[i-1][j-1]
				}
				if i >= l-1 && j >= l-1 {
					tmp := sum[i][j]
					if i-l >= 0 {
						tmp -= sum[i-l][j]
					}
					if j-l >= 0 {
						tmp -= sum[i][j-l]
					}
					if i-l >= 0 && j-l >= 0 {
						tmp += sum[i-l][j-l]
					}
					if tmp == l*l {
						return true
					}
				}
			}
		}
		return false
	}

	l, r := 1, min(inf, m, n)+1

	for l < r {
		mid := (l + r) / 2
		if !check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r - 1
}

func solve1(mat [][]int) int {
	n := len(mat)
	m := len(mat[0])
	// 需要直到在[l][l]中的最小值
	// 如何快速的直到在给定l时，[i...i+l][j...j+l]中的最小值
	// 假设先找出了第一个l矩阵的最小值
	// 如果直到每行到c为止，最多l个元素的最小值
	// 那是否是连续l行的最小值？貌似是的
	// 每行的最小值，可以用双端队列进行计算
	dp := make([][]Pair, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]Pair, n)
	}
	pos := make([][]int, m)
	for i := 0; i < m; i++ {
		pos[i] = make([]int, 2)
	}

	row := make([]int, m)

	check := func(l int) bool {
		if l == 1 {
			return true
		}
		for i := 0; i < m; i++ {
			// tail
			pos[i][0] = 0
			// head
			pos[i][1] = 0
		}

		for r := 0; r < n; r++ {
			var tail, head int

			for c := 0; c < m; c++ {
				for head > tail && mat[r][c] < mat[r][row[head-1]] {
					head--
				}
				row[head] = c
				head++
				if c >= l && c-l == row[tail] {
					tail++
				}
				// r行到c为止，长度为l的arr的最小值
				mv := mat[r][row[tail]]

				for pos[c][1] > pos[c][0] && mv < dp[c][pos[c][1]-1].first {
					pos[c][1]--
				}

				dp[c][pos[c][1]] = Pair{mv, r}
				pos[c][1]++

				if r >= l && r-l == dp[c][pos[c][0]].second {
					pos[c][0]++
				}

				mv = dp[c][pos[c][0]].first
				if r >= l-1 && c >= l-1 && mv >= l {
					return true
				}
			}
		}
		return false
	}

	l, r := 1, min(inf, m, n)+1

	for l < r {
		mid := (l + r) / 2
		if !check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r - 1
}

type Pair struct {
	first  int
	second int
}

func min(arr ...int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if res > arr[i] {
			res = arr[i]
		}
	}
	return res
}
