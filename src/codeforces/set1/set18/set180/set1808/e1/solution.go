package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k, m := readThreeNums(reader)

	res := solve(n, k, m)

	fmt.Println(res)
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

func solve(n int, k int, m int) int {
	// n digits, (0...k-1) meets, some x = (sum of other digits) % k
	add := func(a, b int) int {
		a += b
		if a >= m {
			a -= m
		}
		return a
	}

	sub := func(a, b int) int {
		return add(a, m-b)
	}

	mul := func(a, b int) int {
		return a * b % m
	}

	pow := func(a, b int) int {
		r := 1
		for b > 0 {
			if b&1 == 1 {
				r = mul(r, a)
			}
			a = mul(a, a)
			b >>= 1
		}
		return r
	}

	dp := make([]int, k)
	fp := make([]int, k)

	calc := func(sum int) int {
		// forbid those x, where 2 * x % m = sum
		for i := 0; i < k; i++ {
			dp[i] = 0
		}
		dp[0] = 1
		for i := 1; i <= n; i++ {
			// process digit i
			for j := 0; j < k; j++ {
				fp[j] = 0
			}
			for j := 0; j < k; j++ {
				for x := 0; x < k; x++ {
					if (2*x)%k == sum {
						// forbid
						continue
					}
					fp[(j+x)%k] = add(fp[(j+x)%k], dp[j])
				}
			}
			copy(dp, fp)
		}

		return dp[sum]
	}

	// 2 * x % m = sum % m
	// 如果确定了sum，那么x就相当于确定了

	// 总数
	res := pow(k, n)

	for sum := 0; sum < k; sum++ {
		res = sub(res, calc(sum))
	}

	return res
}
