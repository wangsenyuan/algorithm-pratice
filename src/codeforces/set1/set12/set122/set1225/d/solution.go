package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func process(reader *bufio.Reader) int {
	n, k := readTwoNums(reader)
	a := readNNums(reader, n)
	return solve(k, a)
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

func solve(k int, a []int) int {
	x := slices.Max(a)
	lpf := make([]int, x+1)
	var primes []int

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

	// a * b = x ** k
	// 那么就是 a, b的质因数的power的和，要能等于k
	// 比如 x = 6, k = 3
	// (2**3) * (3 ** 3)

	type pair struct {
		first  int
		second int
	}

	get := func(num int) (res int, pw []pair) {
		res = 1
		for num > 1 {
			x := lpf[num]
			var p int
			for num%x == 0 {
				p++
				num /= x
			}
			p %= k
			if p > 0 {
				pw = append(pw, pair{x, p})
			}
			res *= pow(x, p)
		}
		slices.SortFunc(pw, func(a, b pair) int {
			return a.first - b.first
		})
		return
	}

	freq := make([]int, x+1)
	mx := make([]int, x+1)
	var res int
	for _, num := range a {
		v, arr := get(num)

		if len(arr) == 0 {
			res += freq[1]
		} else {
			expect := 1
			for _, cur := range arr {
				v, p := cur.first, cur.second
				p = k - p
				if p > mx[v] || p < 0 {
					// no such value
					expect = 0
					break
				}
				expect *= pow(v, p)
			}

			if expect <= x {
				res += freq[expect]
			}

			for _, cur := range arr {
				v, p := cur.first, cur.second
				mx[v] = max(mx[v], p)
			}

		}

		freq[v]++
	}

	return res
}
