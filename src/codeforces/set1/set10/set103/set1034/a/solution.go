package main

import (
	"bufio"
	"fmt"
	"os"
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

func solve(a []int) int {
	X := max_of(a) + 1
	lpf := make([]int, X)
	var primes []int
	for i := 2; i < X; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for j := 0; j < len(primes); j++ {
			if i*primes[j] >= X {
				break
			}
			lpf[i*primes[j]] = primes[j]
			if i%primes[j] == 0 {
				break
			}
		}
	}
	var g int
	for _, num := range a {
		g = gcd(g, num)
	}

	getFactors := func(num int) map[int]int {
		res := make(map[int]int)
		for num > 1 {
			res[lpf[num]]++
			num /= lpf[num]
		}
		return res
	}

	fs := getFactors(g)
	ws := make(map[int]int)

	for _, num := range a {
		xs := getFactors(num)
		for k, v := range xs {
			if v > fs[k] {
				ws[k]++
			}
		}
	}
	n := len(a)
	res := n

	for _, p := range primes {
		// g 中包含这个
		if ws[p] > 0 {
			res = min(res, n-ws[p])
		}
	}

	if res == n {
		return -1
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func max_of(arr []int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
}
