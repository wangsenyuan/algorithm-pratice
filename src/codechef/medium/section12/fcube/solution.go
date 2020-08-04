package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := make([]int64, n)
		s, _ := reader.ReadBytes('\n')
		var pos int = -1
		for i := 0; i < n; i++ {
			var x uint64
			pos = readUint64(s, pos+1, &x)
			A[i] = int64(x)
		}

		fmt.Println(solve(n, A))
	}
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

const MOD = 1000000007
const MX = 1000001

var primes []int64

func init() {
	mark := make([]bool, MX)

	for x := 2; x < MX; x++ {
		if !mark[x] {
			primes = append(primes, int64(x))
			for y := x; y < MX; y += x {
				mark[y] = true
			}
		}
	}
}

func solve(n int, A []int64) int {
	cnt := make(map[int64]int)

	for i := 0; i < n; i++ {
		for j := 0; j < len(primes); j++ {
			for A[i]%primes[j] == 0 {
				cnt[primes[j]]++
				A[i] /= primes[j]
			}
		}
	}

	B := make([]int64, n)

	for i := 0; i < n; i++ {
		B[i] = A[i]
		for j := 0; j < n; j++ {
			g := gcd(A[i], A[j])
			if g > 1 && g < A[i] {
				cnt[g]++
				cnt[A[i]/g]++
				B[i] = 1
				break
			}
		}
	}

	// if B[i] > 1, it may be a square of prime number
	for i := 0; i < n; i++ {
		x := int64(math.Sqrt(float64(B[i])))

		for d := -2; d <= 2; d++ {
			y := x + int64(d)
			if y > 1 && y*y == B[i] {
				cnt[y] += 2
				B[i] = 1
				break
			}
		}
	}

	// else it is self a prime number
	for i := 0; i < n; i++ {
		if B[i] > 1 {
			cnt[B[i]]++
		}
	}

	var ans int64 = 1

	for k, v := range cnt {
		if v%3 != 0 {
			v += 3 - v%3
		}

		k %= MOD

		for i := 0; i < v; i++ {
			ans = (ans * k) % MOD
		}
	}

	return int(ans)
}

func gcd(a, b int64) int64 {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
