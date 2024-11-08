package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	res := solve(n)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, e := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", e[0], e[1]))
	}

	fmt.Print(buf.String())
}

func getPrimes(n int) ([]int, []int) {
	var primes []int
	lpf := make([]int, n+1)
	for i := 2; i <= n; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, j := range primes {
			if i*j > n {
				break
			}
			lpf[i*j] = j
			if i%j == 0 {
				break
			}
		}
	}
	return primes, lpf
}

func solve(n int) [][]int {
	_, lpf := getPrimes(n * 2)

	var res [][]int
	deg := make([]int, n)
	// 先组成一个环
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		deg[i]++
		deg[j]++
		res = append(res, []int{i + 1, j + 1})
	}

	// 总数不是质数
	for i := 0; lpf[len(res)] != len(res); i = (i + 1) % n {
		if deg[i] == 3 {
			continue
		}
		j := (i + 2) % n
		deg[j]++
		res = append(res, []int{i + 1, j + 1})
	}

	return res
}
