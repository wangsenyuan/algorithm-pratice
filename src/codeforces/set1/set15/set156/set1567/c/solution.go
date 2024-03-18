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
		res := solve(n)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

const X = 20

var ways [X]int

func init() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			ways[i+j]++
		}
	}
}

func solve(n int) int {
	arr := getDigits(n)

	var x int
	var y int
	for i, num := range arr {
		if i&1 == 0 {
			x = x*10 + num
		} else {
			y = y*10 + num
		}
	}

	return (x+1)*(y+1) - 2
}

func solve1(n int) int {
	// a + b => n

	arr := getDigits(n)

	m := len(arr)

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, 4)
		for j := 0; j < 4; j++ {
			dp[i][j] = -1
		}
	}

	var dfs func(i int, carry int) int
	dfs = func(i int, carry int) int {
		cur := arr[i]
		x := (carry >> 1) & 1

		if i == m-1 {
			// 最后一位
			if carry&1 == 1 {
				// invalid
				return 0
			}
			if x == 1 {
				// 有进位
				cur += 10
			}
			// a + b = cur
			return ways[cur]
		}
		if dp[i][carry] < 0 {
			if x == 1 {
				cur += 10
			}
			var res int
			for expect := 0; expect < 2; expect++ {
				if cur >= expect {
					res += ways[cur-expect] * dfs(i+1, ((carry&1)<<1)|expect)
				}
			}
			dp[i][carry] = res
		}

		return dp[i][carry]
	}

	// 包含了 (n, 0) (0, n)
	res := dfs(0, 0) - 2

	return res
}

func getDigits(n int) []int {
	var arr []int
	for tmp := n; tmp > 0; tmp /= 10 {
		arr = append(arr, tmp%10)
	}

	reverse(arr)
	return arr
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
