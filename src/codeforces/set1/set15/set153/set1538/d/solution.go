package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		a, b, k := readThreeNums(reader)
		res := solve(a, b, k)

		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
	}
	fmt.Print(buf.String())
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

const N = 500000

var primes []int

func init() {
	pf := make([]int, N)
	for i := 2; i < N; i++ {
		if pf[i] == 0 {
			pf[i] = i
			primes = append(primes, i)
		}
		for _, j := range primes {
			if i*j >= N {
				break
			}
			pf[i*j] = j
			if i%j == 0 {
				break
			}
		}
	}
}

func solve(a int, b int, k int) bool {
	if a < b {
		a, b = b, a
	}

	n := count(a) + count(b)
	var m int
	if a == b {
		m = 0
	} else if a%b == 0 {
		m = 1
	} else {
		m = 2
	}

	if k < m || k > n {
		return false
	}

	if k == 1 {
		return m == 1
	}
	return true
}

func count(num int) int {
	// how much to change num to 1
	var res int

	for _, i := range primes {
		if i*i > num {
			break
		}
		for num%i == 0 {
			res++
			num /= i
		}
	}

	if num > 1 {
		res++
	}

	return res
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
