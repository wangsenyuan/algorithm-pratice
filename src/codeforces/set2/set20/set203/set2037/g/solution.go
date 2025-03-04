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
	n := readNum(reader)
	a := readNNums(reader, n)
	return solve(a)
}

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func solve(a []int) int {
	x := slices.Max(a)
	lpf := make([]int, x+1)
	var primes []int

	for i := 2; i < len(lpf); i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, j := range primes {
			if i*j >= len(lpf) {
				break
			}
			lpf[i*j] = j
			if i%j == 0 {
				break
			}
		}
	}

	pfs := make([][]int, x+1)
	for i := 2; i <= x; i++ {
		for y := i; y > 1; {
			z := lpf[y]
			pfs[i] = append(pfs[i], z)
			for y%z == 0 {
				y /= z
			}
		}
	}

	fp := make([]int, x+1)

	update := func(num int, v int) {
		fs := pfs[num]
		m := len(fs)
		for state := 1; state < 1<<m; state++ {
			y := 1
			for i := range m {
				if (state>>i)&1 == 1 {
					y *= fs[i]
				}
			}
			fp[y] = add(fp[y], v)
		}
	}

	n := len(a)

	update(a[n-1], 1)

	// 少于一百万的质数虽然有很多，但是对于单个数来说，最多有7个
	// 2 * 3 * 5 * 7 * 11 * 13 * 17
	// 1 << 7 = 128
	// 那么就可以用这个了

	for i := n - 2; i >= 0; i-- {
		var ways int
		fs := pfs[a[i]]
		m := len(fs)
		for state := 1; state < 1<<m; state++ {
			var cnt int
			p := 1
			for i := range m {
				if (state>>i)&1 == 1 {
					cnt++
					p *= fs[i]
				}
			}
			if cnt&1 == 1 {
				ways = add(ways, fp[p])
			} else {
				ways = sub(ways, fp[p])
			}
		}
		update(a[i], ways)
		if i == 0 {
			return ways
		}
	}

	return 0
}
