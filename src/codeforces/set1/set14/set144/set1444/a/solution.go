package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	var buf bytes.Buffer

	for n > 0 {
		n--
		nums := readNInt64s(reader, 2)
		res := calc(nums[0], nums[1])
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Println(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(P []int64, Q []int64) []int64 {
	ans := make([]int64, len(P))
	for i := 0; i < len(P); i++ {
		ans[i] = calc(P[i], Q[i])
	}
	return ans
}

func calc(p int64, q int64) int64 {
	// find some x
	// p % x = 0
	// and x % q != 0
	// if p % q != 0 => p
	if p%q != 0 {
		return p
	}
	// p % q == 0
	// 对它们进行质因数分解
	// 对于任意的质因数x, fp[x] >= fq[x]
	// 那么只要有一个质因数的个数，要比 fq[x]小
	// 比如如果 p % 2 == 0, 但是q % 2 != 0
	// 那么只要除掉一个2就可以了
	k := int(q)

	cnt := make(map[int]int)
	var best int64 = 1

	for x := 2; x <= k/x; x++ {

		for k%x == 0 {
			cnt[x]++
			k /= x
		}
	}

	if k > 1 {
		cnt[k]++
	}

	for x := range cnt {
		tmp := p
		z := int64(x)
		for tmp%z == 0 {
			tmp /= z
			if tmp%q != 0 {
				best = max(best, tmp)
			}
		}
	}

	return best
}

func pow(a, b int) int64 {
	x := int64(a)
	var res int64 = 1
	for b > 0 {
		if b&1 == 1 {
			res *= x
		}
		x *= x
		b >>= 1
	}
	return res
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
