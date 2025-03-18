package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	fmt.Println(solve(n))
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

func d(x int) int {
	if x < 10 {
		return x
	}
	var s int
	for x > 0 {
		s += x % 10
		x /= 10
	}
	return d(s)
}

func solve(n int) int {
	if n <= 3 {
		return 0
	}
	lpf := make([]int, n+1)
	var primes []int
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

	freq := make([]int, 10)
	var ans int
	// 迭代C
	// (1, 1, 1)
	freq[1]++

	ans--

	for c := 2; c <= n; c++ {
		s := d(c)

		freq[s]++

		cnt := 1
		for num := c; num > 1; {
			p := lpf[num]
			var tmp int
			for num%p == 0 {
				tmp++
				num /= p
			}
			cnt *= (tmp + 1)
		}
		ans -= cnt
	}

	for a := 1; a <= 9; a++ {
		for b := 1; b <= 9; b++ {
			c := d(a * b)
			ans += freq[a] * freq[b] * freq[c]
		}
	}

	return ans
}

func bruteForce(n int) int {
	var res int
	for a := 1; a <= n; a++ {
		for b := 1; b <= n; b++ {
			for c := 1; c <= n; c++ {
				if d(d(a)*d(b)) == d(c) && a*b != c {
					res++
				}
			}
		}
	}
	return res
}
