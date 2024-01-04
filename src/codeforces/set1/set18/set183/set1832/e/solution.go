package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	nums := readNNums(reader, 6)

	res := solve(nums[0], nums[1], nums[2], nums[3], nums[4], nums[5])

	fmt.Println(res)
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

const MOD = 998244353

func add(a, b, m int) int {
	a += b
	a %= m
	return a
}
func mul(a, b, m int) int {
	return a * b % m
}
func pow(a, b, m int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = mul(r, a, m)
		}
		a = mul(a, a, m)
		b >>= 1
	}
	return r
}

func prefSum(arr []int) []int {
	ans := make([]int, len(arr)+1)

	for i := 0; i < len(arr); i++ {
		ans[i+1] = add(ans[i], arr[i], MOD)
	}

	return ans
}

func solve(n int, a1 int, x int, y int, m int, k int) int {

	a := make([]int, n)
	a[0] = a1

	for i := 1; i < n; i++ {
		a[i] = add(mul(a[i-1], x, m), y, m)
	}

	for i := 0; i <= k; i++ {
		a = prefSum(a)
	}

	var ans int
	for i := 1; i <= n; i++ {
		ans ^= (a[i+1] * i)
	}

	return ans
}
