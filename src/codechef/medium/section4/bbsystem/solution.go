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
		res := solve(n)
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

const N_MAX = 2229283
const MOD = 500009
const SQRT_N_MAX = 1494
const DIV_MAX = 321

var R [N_MAX]int

func init() {
	ps := getPrime(SQRT_N_MAX)

	minDiv := make([]int, N_MAX)

	for _, p := range ps {
		for j := p; j < N_MAX; j += p {
			if minDiv[j] == 0 {
				minDiv[j] = p
			}
		}
	}

	div := make([]int, N_MAX)
	per := make([]int, N_MAX)

	deg := make([]int16, N_MAX)

	div[1] = 1

	for i := 2; i < N_MAX; i++ {
		mi := minDiv[i]
		if mi == 0 {
			mi = i
		}
		j := i / mi
		if minDiv[j] != mi {
			deg[i] = 2
			per[i] = div[j]
			div[i] = div[j] * 2
		} else {
			deg[i] = deg[j] + 1
			per[i] = per[j]
			div[i] = per[j] * int(deg[i])
		}
	}
	sum := make([]int, DIV_MAX)
	R[0] = 1
	R[1] = 1
	for i := 2; i < N_MAX; i++ {
		sum[div[i]]++
		R[i] = int(int64(R[i-1]) * int64(sum[div[i]]) % MOD)
	}
}

func getPrime(n int) []int {
	n2 := n / 2
	res := make([]int, n2)
	for i := 0; i < n2; i++ {
		res[i] = 1
	}

	for i := 3; i*i < n; i += 2 {
		if res[i>>1] == 1 {
			for j := (i * i) >> 1; j < n2; j += i {
				res[j] = 0
			}
		}
	}

	var p int
	res[p] = 2
	p++
	for i := 1; i < n2; i++ {
		if res[i] == 1 {
			res[p] = 2*i + 1
			p++
		}
	}
	return res[:p]
}

func solve(n int) int {
	if n >= N_MAX {
		return MOD - 1
	}
	return R[n] - 1
}
