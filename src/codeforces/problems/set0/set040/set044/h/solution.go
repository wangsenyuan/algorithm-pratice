package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s := readString(reader)

	res := solve(s)

	fmt.Println(res)
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
		if s[i] == '\n' || s[i] == '\r' {
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

func solve(num string) int64 {
	n := len(num)

	dp := make([][]int64, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int64, 10)
	}

	for i := 0; i < 10; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < n; i++ {
		// dp[n][i]
		for x := 0; x < 10; x++ {
			sum := x + int(num[i]-'0')
			half := sum / 2
			dp[i][half] += dp[i-1][x]
			if half*2 < sum {
				dp[i][half+1] += dp[i-1][x]
			}
		}
	}

	var res int64

	for i := 0; i < 10; i++ {
		res += dp[n-1][i]
	}

	ok := true
	prev := int(num[0] - '0')
	for i := 1; i < n && ok; i++ {
		sum := prev + int(num[i]-'0')
		half := sum / 2
		if half == int(num[i]-'0') {
			prev = half
			continue
		}
		if sum%2 == 1 && half+1 == int(num[i]-'0') {
			prev = half + 1
			continue
		}
		ok = false
	}

	if ok {
		res--
	}

	return res
}
