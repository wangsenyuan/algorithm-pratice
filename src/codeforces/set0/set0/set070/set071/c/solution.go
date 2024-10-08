package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	w := readNNums(reader, n)
	res := solve(w)

	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
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

func solve(knights []int) bool {
	n := len(knights)

	var primes []int
	marked := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		if !marked[i] {
			primes = append(primes, i)
		}
		for _, p := range primes {
			if i*p > n {
				break
			}
			marked[i*p] = true
			if i%p == 0 {
				break
			}
		}
	}

	sum := make([]int, n)

	check := func(m int) bool {
		// m边形, n % m = 0
		cnt := n / m
		for i := 0; i < n; i++ {
			sum[i%cnt] += knights[i]
		}
		for i := 0; i < cnt; i++ {
			if sum[i] == m {
				return true
			}
			sum[i] = 0
		}
		return false
	}

	// 4边形，不是质数
	if n%4 == 0 && check(4) {
		return true
	}

	for _, p := range primes {
		if p > 2 && n%p == 0 && check(p) {
			return true
		}
	}

	return false
}
