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
	a := readNNums(reader, n)

	res := solve(a)

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

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

const X = 1010

func solve(a []int) int {
	n := len(a)
	dp := make([]int, n+1)

	divs := make(map[int][]int)

	getFactors := func(num int) []int {
		if v, ok := divs[num]; ok {
			return v
		}
		var cur []int
		for i := 1; i <= num/i; i++ {
			if num%i == 0 {
				cur = append(cur, i)
				if num != i*i {
					cur = append(cur, num/i)
				}
			}
		}
		sort.Ints(cur)
		divs[num] = cur

		return cur
	}

	dp[0] = 1
	for i := 1; i <= n; i++ {
		num := a[i-1]
		fs := getFactors(num)
		fp := make(map[int]int)
		for _, x := range fs {
			if x > i {
				break
			}
			// x <= i
			fp[x] = add(fp[x], dp[x-1])
		}

		for x, v := range fp {
			dp[x] = add(dp[x], v)
		}
	}

	var res int
	for j := 1; j <= n; j++ {
		res = add(res, dp[j])
	}
	return res
}
