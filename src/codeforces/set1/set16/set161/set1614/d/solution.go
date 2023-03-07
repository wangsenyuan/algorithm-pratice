package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	//var buf bytes.Buffer
	n := readNum(reader)
	A := readNNums(reader, n)
	res := solve(A)
	fmt.Println(res)
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

const MOD = 1e9 + 7

func modMul(a, b int) int {
	return int(int64(a) * int64(b) % MOD)
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = modMul(r, a)
		}
		a = modMul(a, a)
		b >>= 1
	}
	return r
}

func solve(A []int) int64 {
	// A[i] <= 5 * 1e6
	n := len(A)
	var x int
	for i := 0; i < n; i++ {
		if x < A[i] {
			x = A[i]
		}
	}

	cnt := make([]int, x+1)
	for i := 0; i < n; i++ {
		cnt[A[i]]++
	}

	vis := make([]bool, x+1)
	var pr []int

	for i := 2; i <= x; i++ {
		if !vis[i] {
			pr = append(pr, i)
			if x/i < i {
				continue
			}
			for j := i * i; j <= x; j += i {
				vis[j] = true
			}
		}
	}

	for _, p := range pr {
		for j := x / p; j > 0; j-- {
			cnt[j] += cnt[j*p]
		}
	}
	var best int64
	fp := make([]int64, x+1)
	for i := x; i >= 1; i-- {
		fp[i] = int64(cnt[i]) * int64(i)
		for j := 0; j < len(pr) && pr[j] <= x/i; j++ {
			fp[i] = max(fp[i], fp[pr[j]*i]+int64(i)*int64(cnt[i]-cnt[pr[j]*i]))
		}
		best = max(best, fp[i])
	}
	return best
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
