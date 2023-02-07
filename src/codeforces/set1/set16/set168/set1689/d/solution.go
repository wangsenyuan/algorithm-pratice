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
		G := make([]string, n)
		for i := 0; i < n; i++ {
			G[i] = readString(reader)[:m]
		}
		res := solve(n, m, G)
		buf.WriteString(fmt.Sprintf("%d %d\n", res[0], res[1]))
	}

	fmt.Print(buf.String())
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

func solve(n int, m int, G []string) []int {
	// 是不是只需要找到最远的4个black cell即可？

	targets := make([]Pair, 4)
	targets[0] = Pair{n, m}
	targets[1] = Pair{-n, -m}
	targets[2] = Pair{n, -m}
	targets[3] = Pair{-n, m}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if G[i][j] == 'B' {
				if i+j < targets[0].first+targets[0].second {
					targets[0] = Pair{i, j}
				}
				if i+j > targets[1].first+targets[1].second {
					targets[1] = Pair{i, j}
				}
				if i-j < targets[2].first-targets[2].second {
					targets[2] = Pair{i, j}
				}

				if i-j > targets[3].first-targets[3].second {
					targets[3] = Pair{i, j}
				}
			}
		}
	}

	res := n + m

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			tmp := 0
			for k := 0; k < 4; k++ {
				dx := abs(i - targets[k].first)
				dy := abs(j - targets[k].second)
				tmp = max(tmp, dx+dy)
			}
			res = min(res, tmp)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			tmp := 0
			for k := 0; k < 4; k++ {
				dx := abs(i - targets[k].first)
				dy := abs(j - targets[k].second)
				tmp = max(tmp, dx+dy)
			}
			if tmp == res {
				return []int{i + 1, j + 1}
			}
		}
	}
	return nil
}

type Pair struct {
	first  int
	second int
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
