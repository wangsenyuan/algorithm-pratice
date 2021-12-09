package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, k := readTwoNums(reader)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readNum(reader)
	}
	res := solve(k, nums)
	var buf bytes.Buffer
	for _, ans := range res {
		buf.WriteString(fmt.Sprintf("%d\n", ans))
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
const MAX_N = 200010

var F [MAX_N]int64
var I [MAX_N]int64
var sum [MAX_N]int64
var mobius [MAX_N]int64
var divs [][]int

func pow(a, b int64) int64 {
	var r int64 = 1
	for b > 0 {
		if b&1 == 1 {
			r = (r * a) % MOD
		}
		a = (a * a) % MOD
		b >>= 1
	}
	return r
}

func mul(a, b int64) int64 {
	a *= b
	a %= MOD
	return a
}

func add(a, b int64) int64 {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func sub(a, b int64) int64 {
	return add(a, MOD-b)
}

func init() {
	F[0] = 1
	for i := 1; i < MAX_N; i++ {
		F[i] = int64(i) * F[i-1] % MOD
	}

	I[MAX_N-1] = pow(F[MAX_N-1], MOD-2)

	for i := MAX_N - 2; i >= 0; i-- {
		I[i] = int64(i+1) * I[i+1] % MOD
	}

	for i := 0; i < MAX_N; i++ {
		mobius[i] = 1
	}
	mobius[1] = 0

	set := make([]bool, MAX_N)

	for i := 2; i < MAX_N; i++ {
		if !set[i] {

			i2 := int64(i) * int64(i)
			for j := i; j < MAX_N; j += i {
				set[j] = true
				// *= -1
				mobius[j] = mul(mobius[j], MOD-1)
				if int64(j)%(i2) == 0 {
					mobius[j] = 0
				}
			}
		}
	}

	divs = make([][]int, MAX_N)
	for i := 1; i < MAX_N; i++ {
		divs[i] = make([]int, 0, 2)
	}
	for i := 1; i < MAX_N; i++ {
		for j := i; j < MAX_N; j += i {
			divs[j] = append(divs[j], i)
		}
	}

	for i := 1; i < MAX_N; i++ {
		for _, x := range divs[i] {
			sum[i] = add(sum[i], mul(mobius[x], int64(i/x)))
		}
	}
}

func nCr(n, r int) int64 {
	if r < 0 || n < r {
		return 0
	}
	return mul(F[n], mul(I[r], I[n-r]))
}

func solve(k int, nums []int) []int {
	res := make([]int64, 200001)
	prev := make([]int64, 200001)
	var ans int64
	for i := k; i <= 200000; i++ {
		for _, x := range divs[i] {
			now := add(int64(x), sum[x])
			ans = sub(ans, mul(prev[x], now))
			newCoff := nCr(i/x, k)
			ans = add(ans, mul(newCoff, now))
			prev[x] = newCoff
		}
		den := mul(I[i], mul(F[k], F[i-k]))
		res[i] = mul(den, ans)
	}

	ret := make([]int, len(nums))

	for i, n := range nums {
		ret[i] = int(res[n])
	}

	return ret
}
