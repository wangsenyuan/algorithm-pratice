package main

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
		nums := readNNums(reader, n)
		res := solve(nums)

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

const X = 100_010

var primes []int
var lpf []int

func init() {
	lpf = make([]int, X)
	for i := 2; i < X; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for j := 0; j < len(primes); j++ {
			if primes[j]*i >= X {
				break
			}
			lpf[primes[j]*i] = primes[j]
			if i%primes[j] == 0 {
				break
			}
		}
	}
}

func solve(a []int) bool {

	cnt := make(map[int]int)

	for _, x := range a {
		if x < X && lpf[x] == x {
			// a prime number
			cnt[x]++
			if cnt[x] > 1 {
				return true
			}
		} else {
			for i := 0; i < len(primes) && primes[i]*primes[i] <= x; i++ {
				// 看起来这样很慢，但其实是比较快的
				if x%primes[i] == 0 {
					cnt[primes[i]]++
					if cnt[primes[i]] > 1 {
						return true
					}
					for x%primes[i] == 0 {
						x /= primes[i]
					}
				}
			}
			if x > 0 {
				cnt[x]++
				if cnt[x] > 1 {
					return true
				}
			}
		}
	}

	return false
}
