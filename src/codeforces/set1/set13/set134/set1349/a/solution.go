package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

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
	a := readNNums(reader, n)
	res := solve(a)

	fmt.Println(res)
}

const inf = 1 << 30

func solve(a []int) int {
	ma := slices.Max(a) + 1
	lpf := make([]int, ma)

	var primes []int
	for i := 2; i < ma; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, j := range primes {
			if i*j >= ma {
				break
			}
			lpf[i*j] = j
			if i%j == 0 {
				break
			}
		}
	}

	pw1 := make([][]int, ma)

	freq := make([]int, ma)

	for _, num := range a {
		for num > 1 {
			x := lpf[num]
			var cnt int
			for num%x == 0 {
				cnt++
				num /= x
			}
			freq[x]++
			pw1[x] = append(pw1[x], cnt)
		}
	}

	n := len(a)

	for _, x := range primes {
		// 说明x肯定在某个地方不存在
		if freq[x] < n-1 {
			pw1[x] = pw1[x][:0]
			continue
		}
		if freq[x] < n {
			pw1[x] = append(pw1[x], 0)
		}
	}

	res := 1
	for _, x := range primes {
		if len(pw1[x]) > 1 {
			sort.Ints(pw1[x])
			res *= pow(x, pw1[x][1])
		}
	}

	return res
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r *= a
		}
		a *= a
		b >>= 1
	}
	return r
}
