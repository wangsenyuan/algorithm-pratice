package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
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
	b := readNNums(reader, 2*n)
	res := solve(b)

	s := fmt.Sprintf("%v", res)

	fmt.Println(s[1 : len(s)-1])
}

func getPrimes(x int) (lpf []int, primes []int) {
	// x <= 2750131
	lpf = make([]int, x+1)
	for i := 2; i <= x; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, j := range primes {
			if i*j > x {
				break
			}
			lpf[i*j] = j
			if i%j == 0 {
				break
			}
		}
	}
	return
}

func solve(a []int) []int {
	x := slices.Max(a)

	lpf, primes := getPrimes(x)

	from := make([]int, x+1)

	for _, v := range primes {
		if v <= len(primes) {
			from[primes[v-1]] = v
		}
	}

	sort.Ints(a)
	reverse(a)

	// 考虑最大的那个数，如果这个数是个质数，那么它必须在右边，否则必须在左边
	freq := make([]int, x+1)
	for _, v := range a {
		freq[v]++
	}

	var res []int

	for _, v := range a {
		if freq[v] == 0 {
			continue
		}
		freq[v]--
		// v是目前最大的数
		if lpf[v] == v {
			// 它必须在右边
			w := from[v]
			res = append(res, w)
			freq[w]--
		} else {
			res = append(res, v)
			w := v / lpf[v]
			freq[w]--
		}
	}

	return res
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
