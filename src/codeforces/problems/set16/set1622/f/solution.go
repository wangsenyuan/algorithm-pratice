package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	res := solve(n)

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%d\n", len(res)))

	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	buf.WriteByte('\n')
	fmt.Print(buf.String())
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

var pr []int
var f []int64
var rf map[int64]int

func prepare(MAX_N int) {
	pr = make([]int, MAX_N)

	rd := rand.NewSource(17)
	hs := make([]int64, MAX_N)

	for i := 2; i < MAX_N; i++ {
		if pr[i] == 0 {
			for j := i; j < MAX_N; j += i {
				if pr[j] == 0 {
					pr[j] = i
				}
			}
		}
		hs[i] = rd.Int63()
	}

	f = make([]int64, MAX_N)

	rf = make(map[int64]int)

	for i := 2; i < MAX_N; i++ {
		f[i] = f[i-1]
		x := i
		for x > 1 {
			f[i] ^= hs[pr[x]]
			x /= pr[x]
		}
		rf[f[i]] = i
	}
}

func solve(n int) []int {
	prepare(n + 1)
	check := func() []int {
		var fp int64

		for i := 2; i <= n; i++ {
			fp ^= f[i]
		}

		if fp == 0 {
			return nil
		}
		if v, ok := rf[fp]; ok {
			return []int{v}
		}
		for i := 2; i <= n; i++ {
			x := f[i] ^ fp
			if v, ok := rf[x]; ok && v != i {
				return []int{v, i}
			}
		}
		return []int{2, n / 2, n}
	}
	rem := check()
	ans := make([]int, 0, n-len(rem))

	for i := 1; i <= n; i++ {
		if !find(rem, i) {
			ans = append(ans, i)
		}
	}
	return ans
}

func find(arr []int, target int) bool {
	for _, num := range arr {
		if num == target {
			return true
		}
	}
	return false
}
