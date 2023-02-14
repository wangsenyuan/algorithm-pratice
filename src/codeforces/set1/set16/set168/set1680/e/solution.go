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
		n := readNum(reader)
		S := make([]string, 2)
		S[0] = readString(reader)
		S[1] = readString(reader)
		res := solve(n, S)
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

const INF = 1 << 30

func solve(n int, S []string) int {
	// 肯定是停在某个 *
	first := make([]int, 2)
	last := make([]int, 2)

	for i := 0; i < 2; i++ {
		first[i] = -1
		last[i] = -1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < 2; j++ {
			if S[j][i] == '*' {
				if first[j] < 0 {
					first[j] = i
				}
				last[j] = i
			}
		}
	}
	if first[0] < 0 {
		// no * at first row
		return last[1] - first[1]
	}
	if first[1] < 0 {
		return last[0] - first[0]
	}

	a := min(first[0], first[1])
	b := max(last[0], last[1])
	dp := make([]int, 2)
	fp := make([]int, 2)
	if S[1][a] == '*' {
		dp[0]++
	}
	if S[0][a] == '*' {
		dp[1]++
	}
	for i := a + 1; i <= b; i++ {
		for j := 0; j < 2; j++ {
			fp[j] = INF
		}
		for j := 0; j < 2; j++ {
			var x int
			if S[1-j][i] == '*' {
				x++
			}
			fp[j] = min(fp[j], dp[j]+1+x, dp[1-j]+2)
		}
		copy(dp, fp)
	}

	return min(dp[0], dp[1])
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

func max(arr ...int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
}
