package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, x := readTwoNums(reader)
	songs := make([][]int, n)
	for i := 0; i < n; i++ {
		songs[i] = readNNums(reader, 2)
	}

	res := solve(x, songs)

	fmt.Println(res)
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

const mod = 1e9 + 7

func add(a ...int) int {
	res := a[0]
	for i := 1; i < len(a); i++ {
		res += a[i]
		if res >= mod {
			res -= mod
		}
	}
	return res
}

func mul(a ...int) int {
	res := a[0]
	for i := 1; i < len(a); i++ {
		res *= a[i]
		res %= mod
	}
	return res
}

func solve(T int, songs [][]int) int {
	n := len(songs)

	// n <= 15 可以用 bitmask

	N := 1 << n
	dp := make([][4]int, N)

	dp[0][3] = 1

	var ans int

	for state := 0; state < N; state++ {
		for last := 0; last < 4; last++ {
			for i := 0; i < n; i++ {
				if (state>>i)&1 == 0 && songs[i][1] != last+1 {
					// can append
					next := state | (1 << i)
					dp[next][songs[i][1]-1] = add(dp[next][songs[i][1]-1], dp[state][last])
				}
			}
			var sum int
			for i := 0; i < n; i++ {
				if (state>>i)&1 == 1 {
					sum += songs[i][0]
				}
			}
			if sum == T {
				ans = add(ans, dp[state][last])
			}
		}
	}

	return ans
}
