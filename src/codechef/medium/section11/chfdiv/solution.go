package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		var N, K uint64
		s, _ := reader.ReadBytes('\n')
		pos := readUint64(s, 0, &N)
		readUint64(s, pos+1, &K)
		fmt.Println(solve(int64(N), int64(K)))
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
const MAX_N = 1000007

var primes []int

func init() {
	set := make([]bool, MAX_N)
	primes = make([]int, 0, 10000)

	for x := 2; x < MAX_N; x++ {
		if !set[x] {
			primes = append(primes, x)
			for y := x + x; y < MAX_N; y += x {
				set[y] = true
			}
		}
	}
}

func solve(N, K int64) int {
	var ans int64 = 1
	for _, p := range primes {
		P := int64(p)
		if P > N {
			break
		}
		if K%P == 0 {
			continue
		}
		var cnt int64
		for P <= N {
			cnt += N / P
			P *= int64(p)
		}
		ans *= pow(int64(p), cnt)
		ans %= MOD
	}

	return int(ans)
}

func pow(a, b int64) int64 {
	var r int64 = 1
	for b > 0 {
		if b&1 == 1 {
			r *= a
			r %= MOD
		}
		a *= a
		a %= MOD
		b >>= 1
	}
	return r
}
