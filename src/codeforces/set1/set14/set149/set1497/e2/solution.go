package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		a := readNNums(reader, n)
		res := solve(a, k)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

const X = 1_000_000_9

var lps [X]int

func init() {
	var primes []int
	for x := 2; x < X; x++ {
		if lps[x] == 0 {
			lps[x] = x
			primes = append(primes, x)
		}
		for j := 0; j < len(primes); j++ {
			if x*primes[j] >= X {
				break
			}
			lps[x*primes[j]] = primes[j]
			if x%primes[j] == 0 {
				break
			}
		}
	}
}

func solve(a []int, k int) int {
	n := len(a)
	if k >= n-1 {
		return 1
	}

	a = convert(a)
	ma := slices.Max(a)

	freq := make([]int32, ma+1)

	fp := make([][]int32, n)
	for i := 0; i < n; i++ {
		fp[i] = make([]int32, k+1)
	}

	for j := 0; j <= k; j++ {
		l := n
		now := 0
		for i := n - 1; i >= 0; i-- {
			for l-1 >= 0 && now+check(freq[a[l-1]] > 0) <= j {
				l--
				now += check(freq[a[l]] > 0)
				freq[a[l]]++
			}
			fp[i][j] = int32(l)
			if freq[a[i]] > 1 {
				now--
			}
			freq[a[i]]--
		}
	}

	dp := make([][]int32, n+1)
	dp[0] = make([]int32, k+1)
	for i := 1; i <= n; i++ {
		dp[i] = make([]int32, k+1)
		for j := 0; j <= k; j++ {
			if j > 0 {
				dp[i][j] = dp[i][j-1]
			} else {
				dp[i][j] = 1 << 30
			}
			for x := 0; x <= j; x++ {
				dp[i][j] = min(dp[i][j], dp[fp[i-1][x]][j-x]+1)
			}
		}
	}

	return int(dp[n][k])
}

func check(b bool) int {
	if b {
		return 1
	}
	return 0
}

func convert(a []int) []int {
	for i, num := range a {
		res := 1
		var last int
		var cnt int
		for num > 1 {
			if lps[num] == last {
				cnt++
			} else {
				if cnt%2 == 1 {
					res *= last
				}
				last = lps[num]
				cnt = 1
			}
			num /= lps[num]
		}
		if cnt%2 == 1 {
			res *= last
		}
		a[i] = res
	}

	return a
}
