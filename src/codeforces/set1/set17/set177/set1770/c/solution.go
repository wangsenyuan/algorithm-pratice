package main

import "sort"

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		p := readNNums(reader, n)
		res := solve(p)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
	}
	fmt.Print(buf.String())
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(a []int) bool {
	// gcd(a + x, b + x) = 1
	// 如果gcd(a, b) = 1, 是不是加任何x都是成立的?
	// gcd(a, b) = gcd(a - b, b)
	// gcd(a + x, b + x) = gcd(a - b, b + x)
	// 假设对于b来说，所有比它大的数，和它的差，组成一个数列
	// 假设是 d, 然后得到d的因数列表（要排除1）
	n := len(a)
	sort.Ints(a)
	for i := 0; i+1 < n; i++ {
		if a[i] == a[i+1] {
			return false
		}
	}

	var primes []int
	set := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		if !set[i] {
			primes = append(primes, i)
			set[i] = true
		}
		for j := 0; j < len(primes); j++ {
			if i*primes[j] > n {
				break
			}
			set[i*primes[j]] = true
			if i%primes[j] == 0 {
				break
			}
		}
	}

	cnt := make([]int, n+1)
	for _, p := range primes {
		for i := 0; i < n; i++ {
			x := a[i] % p
			cnt[x]++
		}

		mv := cnt[0]
		for i := 0; i < p; i++ {
			mv = min(mv, cnt[i])
			cnt[i] = 0
		}

		if mv >= 2 {
			return false
		}
	}

	return true
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
