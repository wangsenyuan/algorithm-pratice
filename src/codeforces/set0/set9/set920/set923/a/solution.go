package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	fmt.Println(solve(n))
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\r' || s[i] == '\n' {
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

func solve(x2 int) int {
	lps := make([]int, x2+1)
	var primes []int
	for i := 2; i <= x2; i++ {
		if lps[i] == 0 {
			lps[i] = i
			primes = append(primes, i)
		}
		for j := 0; j < len(primes); j++ {
			if i*primes[j] > x2 {
				break
			}
			lps[i*primes[j]] = primes[j]
			if i%primes[j] == 0 {
				break
			}
		}
	}

	fs := make(map[int]int)
	for tmp := x2; tmp > 1; tmp /= lps[tmp] {
		fs[lps[tmp]]++
	}

	var arr []int
	for k := range fs {
		arr = append(arr, k)
	}
	// arr 是 x2的 primes factors
	sort.Ints(arr)

	res := x2

	// x2 = x1 > x0
	for _, p := range arr {
		// p < x0 && 2 * p = x2
		// 如果 x0 = x2 - p, 那么 x0 就能整除 p, 那么就不会增长
		// 且 如果x0 < x2 - p, 那么 x1 = x2 - p
		x0 := max(x2-p+1, p+1)
		res = min(res, x0)
	}

	// x2 > x1 > x0
	for _, p := range primes {
		// x1 是 p的倍数
		for x1 := 2 * p; x1 < x2; x1 += p {
			x0 := max(x1-p+1, p+1)
			if x0 >= x1 {
				continue
			}
			// x0 < x1
			// 此时需要存在一个x2的prime factor px
			// px < x1, and x1 >= x2 - px + 1
			for _, px := range arr {
				if px >= x1 {
					break
				}
				if x1 >= x2-px+1 {
					// found one
					res = min(res, x0)
					break
				}
			}
		}
	}

	return res
}
