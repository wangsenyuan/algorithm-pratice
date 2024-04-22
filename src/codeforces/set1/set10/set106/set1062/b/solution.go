package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	res := solve(n)

	fmt.Printf("%d %d\n", res[0], res[1])
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

func solve(n int) []int {
	lpf := make([]int, n+1)
	var primes []int
	for i := 2; i <= n; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for j := 0; j < len(primes); j++ {
			if i*primes[j] > n {
				break
			}
			lpf[i*primes[j]] = primes[j]
			if i%primes[j] == 0 {
				break
			}
		}
	}

	freq := make(map[int]int)

	for tmp := n; tmp > 1; tmp /= lpf[tmp] {
		freq[lpf[tmp]]++
	}

	res := 1

	// 需要将v变成正好能被 2整除完
	// 先要把大家变成同样的v
	var mv int
	for k, v := range freq {
		res *= k
		mv = max(mv, v)
	}

	var op int
	var mul int
	for _, v := range freq {
		// v != mv 只需要一次，就可以
		if v != mv {
			mul = 1
			break
		}
	}

	for mv > 1 {
		if mv&1 == 1 {
			// 这里可以在mul的时候，直接乘上一个足够的数
			mv++
			mul = 1
		}
		mv >>= 1
		op++
	}

	return []int{res, op + mul}
}
