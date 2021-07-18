package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc, p := readTwoNums(reader)

	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n, q := readTwoNums(reader)
		Q := make([][]int, q)
		for i := 0; i < q; i++ {
			Q[i] = readNNums(reader, 2)
		}
		res := solve(n, p, Q)
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
}

func solve(n, p int, Q [][]int) []int {
	if p == 0 {
		// or
		return solveMin(n, Q)
	}
	// and
	return solveMax(n, Q)
}

const P = 3000007
const MOD1 = 1000000009
const MOD2 = 1000000007

func solveMax(n int, Q [][]int) []int {

	calc := func(hash []int64, p int64, m int64, cur []int) {
		l := cur[0]
		r := cur[1]
		if l > r {
			l, r = r, l
		}
		l--
		hash[l] = (hash[l] + p) % m
		hash[r] = (hash[r] - p) % m
		for hash[r] < 0 {
			hash[r] += m
		}
		hash[r] %= m
	}

	var p1 int64 = 1
	var p2 int64 = 2
	hash1 := make([]int64, n+1)
	hash2 := make([]int64, n+2)

	for i := 1; i <= len(Q); i++ {
		calc(hash1, int64(i)*p1%MOD1, MOD1, Q[i-1])
		p1 = (p1 * P) % MOD1
		calc(hash2, int64(i)*p2%MOD2, MOD2, Q[i-1])
		p2 = (p2 * P) % MOD2
	}

	mem := make(map[Pair][]int)

	for i := 0; i < n; i++ {
		if i > 0 {
			hash1[i] = (hash1[i] + hash1[i-1]) % MOD1
			hash2[i] = (hash2[i] + hash2[i-1]) % MOD2
		}
		key := Pair{int(hash1[i]), int(hash2[i])}
		if _, found := mem[key]; !found {
			mem[key] = make([]int, 0, 2)
		}
		mem[key] = append(mem[key], i+1)
	}

	ans := make([]int, n)

	for i := 0; i < n; i++ {
		key := Pair{int(hash1[i]), int(hash2[i])}
		if key.first == 0 && key.second == 0 {
			ans[i] = n
		} else {
			l := len(mem[key])
			ans[i] = mem[key][l-1]
			mem[key] = mem[key][:l-1]
		}
	}
	return ans
}

type Pair struct {
	first  int
	second int
}

func solveMin(n int, Q [][]int) []int {
	pref := make([]int, n+1)
	for _, cur := range Q {
		l, r := cur[0], cur[1]
		if l > r {
			l, r = r, l
		}
		l--
		pref[l]++
		pref[r]--
	}

	res := make([]int, n)

	for i := 0; i < n; i++ {
		if i > 0 {
			pref[i] += pref[i-1]
		}
		if pref[i] > 0 {
			res[i] = i + 1
		} else {
			res[i] = 1
		}
	}

	return res
}
