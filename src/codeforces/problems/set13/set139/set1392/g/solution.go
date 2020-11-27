package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, k := readThreeNums(reader)

	a, _ := reader.ReadString('\n')
	b, _ := reader.ReadString('\n')

	swaps := make([][]int, n)

	for i := 0; i < n; i++ {
		swaps[i] = readNNums(reader, 2)
	}

	res := solve(n, m, k, a, b, swaps)

	fmt.Printf("%d\n", res[0])
	fmt.Printf("%d %d\n", res[1], res[2])
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

func solve(n, m, k int, s, t string, swaps [][]int) []int {
	K := 1 << k
	dp := make([]int, K)
	fp := make([]int, K)
	for i := 0; i < K; i++ {
		dp[i] = n
		fp[i] = -n
	}
	S := []byte(s)
	T := []byte(t)
	dp[getMask(S, k)] = 0
	fp[getMask(T, k)] = 0

	p := make([]int, k)
	for i := 0; i < k; i++ {
		p[i] = i
	}

	for i := 1; i <= n; i++ {
		l, r := swaps[i-1][0], swaps[i-1][1]
		l--
		r--
		p[l], p[r] = p[r], p[l]
		for j := 0; j < k; j++ {
			S[p[j]] = s[j]
			T[p[j]] = t[j]
		}
		mask := getMask(S, k)
		dp[mask] = min(dp[mask], i)
		mask = getMask(T, k)
		fp[mask] = i
	}
	x := countDigit(s, '1')
	y := countDigit(t, '1')
	var best int
	var left, right int = 1, 1
	for mask := K - 1; mask >= 0; mask-- {
		if fp[mask]-dp[mask] >= m {

			z := countDigitOne(mask)
			tmp := k - (x + y - 2*z)
			if tmp >= best {
				best = tmp
				left = dp[mask] + 1
				right = fp[mask]
			}
		}

		for j := 0; j < k; j++ {
			if mask&(1<<j) > 0 {
				dp[mask^(1<<j)] = min(dp[mask^(1<<j)], dp[mask])
				fp[mask^(1<<j)] = max(fp[mask^(1<<j)], fp[mask])
			}
		}
	}

	return []int{best, left, right}
}

func countDigitOne(num int) int {
	var res int
	for num > 0 {
		res++
		num -= num & -num
	}
	return res
}

func countDigit(s string, x byte) int {
	var res int
	for i := 0; i < len(s); i++ {
		if s[i] == x {
			res++
		}
	}
	return res
}

func getMask(s []byte, n int) int {
	var res int
	for i := 0; i < n; i++ {
		res |= int(s[i]-'0') << i
	}
	return res
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
