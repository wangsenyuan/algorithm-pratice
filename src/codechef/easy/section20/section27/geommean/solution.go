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
		n, X := readTwoNums(reader)
		A := readNNums(reader, n)
		res := solve(X, A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

type Pair struct {
	first, second int
}

func getPrimeFactors(num int) []Pair {

	var res []Pair

	for i := 2; i <= num/i; i++ {
		if num%i == 0 {
			var cnt int
			for num%i == 0 {
				cnt++
				num /= i
			}
			res = append(res, Pair{i, cnt})
		}
	}

	if num > 1 {
		res = append(res, Pair{num, 1})
	}

	return res
}

func solve(X int, A []int) int64 {
	prime_factors := getPrimeFactors(X)

	var arr [][]int

	var ans int64

	for i := 0; i < len(A); i++ {
		var factors []int
		for j := 0; j < len(prime_factors); j++ {
			var p = prime_factors[j].first
			var cnt int
			for A[i]%p == 0 {
				cnt++
				A[i] /= p
			}
			factors = append(factors, cnt)
		}
		if A[i] > 1 {
			ans += process(arr, prime_factors)
			arr = arr[0:0]
		} else {
			arr = append(arr, factors)
		}
	}

	ans += process(arr, prime_factors)

	return ans
}

func process(arr [][]int, prime_factors []Pair) int64 {
	k := len(prime_factors)

	cnt := make(map[Key]int64)

	curr := NewKey(0)

	cnt[curr]++

	for i := 0; i < len(arr); i++ {
		base := NewKey(1)
		for j := 0; j < k; j++ {
			curr = curr.Add(base.Mul(arr[i][j]))
			curr = curr.Sub(base.Mul(prime_factors[j].second))
			base = base.Mul(k + 1)
		}
		cnt[curr]++
	}
	var ans int64

	for _, v := range cnt {
		ans += int64(v) * int64(v-1) / 2
	}
	return ans
}

const MOD1 = 100000000007
const MOD2 = 100000000009

type Key struct {
	first  int64
	second int64
}

func NewKey(v int) Key {
	V := int64(v)
	return Key{V % MOD1, V % MOD2}
}

func (this Key) Add(that Key) Key {
	return Key{(this.first + that.first) % MOD1, (this.second + that.second) % MOD2}
}

func (this Key) Sub(that Key) Key {
	return this.Add(Key{MOD1 - that.first, MOD2 - that.second})
}

func (this Key) Mul(v int) Key {
	V := int64(v)
	return Key{(this.first * V) % MOD1, (this.second * V) % MOD2}
}
