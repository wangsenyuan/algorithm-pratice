package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	a := readNNums(reader, n-1)
	b := readNNums(reader, n-1)
	res := solve(a, b)
	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("YES\n"))
	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	buf.WriteByte('\n')

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

func solve(a []int, b []int) []int {
	n := len(a) + 1
	dp := make([][]int, n)
	dp[n-1] = make([]int, 4)
	for i := 0; i <= 3; i++ {
		dp[n-1][i] = i
	}

	for i := n - 2; i >= 0; i-- {
		dp[i] = make([]int, 4)
		for j := 0; j < 4; j++ {
			dp[i][j] = -1
		}
		for x := 0; x < 4; x++ {
			if dp[i+1][x] >= 0 {
				for y := 0; y < 4; y++ {
					if a[i] == x|y && b[i] == x&y {
						dp[i][y] = x
					}
				}
			}
		}
	}

	res := make([]int, n)
	res[0] = -1
	for i := 0; i < 4; i++ {
		if dp[0][i] >= 0 {
			res[0] = i
			break
		}
	}

	for i := 0; i+1 < n; i++ {
		if res[i] < 0 {
			return nil
		}
		res[i+1] = dp[i][res[i]]
	}

	if res[n-1] < 0 {
		return nil
	}

	return res
}
