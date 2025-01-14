package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)

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

func process(reader *bufio.Reader) int {
	n, k, p := readThreeNums(reader)
	a := readNNums(reader, n)
	return solve(a, k, p)
}

const inf = 1 << 60

func solve(a []int, k int, p int) int {
	n := len(a)
	if k*p >= n {
		fp := make([][]int, k+1)
		for i := range k + 1 {
			fp[i] = make([]int, p)
			for j := range p {
				fp[i][j] = inf
			}
		}
		fp[0][0] = 0

		dp := make([]int, k+1)

		var sum int
		for _, x := range a {
			sum += x
			for j := k; j > 0; j-- {
				dp[j] = inf
				for d := range p {
					// d = 4的时候
					// sum = 16, p = 12
					if sum < d {
						break
					}
					dp[j] = min(dp[j], fp[j-1][(sum-d)%p]+d)
				}
			}
			for j := 1; j <= k; j++ {
				if dp[j] >= 0 {
					fp[j][sum%p] = min(fp[j][sum%p], dp[j])
				}
			}
		}

		return dp[k]
	}
	var x int

	lis := make([]int, n)
	var pos int

	for i, num := range a {
		x = (x + num) % p
		j := search(pos, func(j int) bool {
			return lis[j] > x
		})
		lis[j] = x
		if j == pos {
			pos++
		}
		if i == n-1 && j >= k-1 {
			return x
		}
	}

	return x + p
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
