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
		A := readNNums(reader, n)
		buf.WriteString(fmt.Sprintf("%d\n", solve(n, A)))
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


const MOD = 998244353
const N = 1000010
var F [N]int64
var IF [N]int64

func init() {
	F[0] = 1
	for i := 1; i < N; i++ {
		F[i] = (int64(i) * F[i-1]) % MOD
	}
	IF[N-1] = pow(F[N-1], MOD - 2)
	for i := N - 2; i >= 0; i-- {
		IF[i] = (int64(i + 1) * IF[i+1]) % MOD
	}
}

func pow(a, b int64) int64 {
	var r int64 = 1
	for b > 0 {
		if b & 1 == 1 {
			r = (r * a) % MOD
		}
		a = (a * a) %MOD
		b >>= 1
	}

	return r
}

func solve(n int, A []int) int {
	cnt := make(map[int]int)
	for _, a := range A {
		cnt[a]++
	}

	var ans int64 = 1

	for i := 1; i <= n; i++ {
		if cnt[i] % i != 0 {
			return 0
		}
		x := cnt[i] / i
		ans = (ans * IF[x]) % MOD
		tmp := pow(int64(i), MOD-2)

		for j := i; j <= cnt[i]; j += i {
			ans = (ans * tmp) % MOD
			ans = (ans * F[j]) % MOD
			ans = (ans * IF[j - i]) % MOD
		}
	}

	return int(ans)
}