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
		L, R := readTwoNums(reader)
		res := solve(L, R)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const N = 100011

var good []int64
var sum []int

func init() {
	good = make([]int64, N)
	var mul int64 = 1
	var idx int
	for left := int64(0); idx < N; left++ {
		if left == mul {
			mul *= 10
		}
		for mid := int64(0); mid < 10 && idx < N; mid++ {
			num := (left*10+mid)*mul + reverse(left)
			good[idx] = num
			idx++
		}
	}
	//sort.Sort(Int64s(good))
	good = good[1:]
	sum = make([]int, N)

	for i := 0; i < N-10; i++ {
		sum[i] = int(good[i] % MOD)
		if i > 0 {
			sum[i] += sum[i-1]
			if sum[i] >= MOD {
				sum[i] -= MOD
			}
		}
	}
}

func reverse(num int64) int64 {
	var res int64
	for num > 0 {
		res = res*10 + num%10
		num /= 10
	}
	return res
}

type Int64s []int64

func (this Int64s) Len() int {
	return len(this)
}

func (this Int64s) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this Int64s) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func solve(L, R int) int {
	L--
	R--
	tmp := sum[R]
	tmp = tmp - sum[L] + MOD
	tmp %= MOD
	return pow(int(good[L]%MOD), tmp)
}

func pow(a, b int) int {
	A := int64(a)
	var R int64 = 1
	for b > 0 {
		if b&1 == 1 {
			R = (R * A) % MOD
		}
		A = (A * A) % MOD
		b >>= 1
	}
	return int(R)
}
