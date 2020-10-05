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

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		s, _ := reader.ReadBytes('\n')
		var V uint64
		pos := readUint64(s, 0, &V)
		var N int
		pos = readInt(s, pos+1, &N)
		s, _ = reader.ReadBytes('\n')
		P := make([]uint64, 2)
		pos = readUint64(s, 0, &P[0])
		pos = readUint64(s, pos+1, &P[1])
		var A0, B0, C0, M0 uint64
		pos = readUint64(s, pos+1, &A0)
		pos = readUint64(s, pos+1, &B0)
		pos = readUint64(s, pos+1, &C0)
		pos = readUint64(s, pos+1, &M0)
		s, _ = reader.ReadBytes('\n')
		Q := make([]uint64, 2)
		pos = readUint64(s, 0, &Q[0])
		pos = readUint64(s, pos+1, &Q[1])
		var A1, B1, C1, M1 uint64
		pos = readUint64(s, pos+1, &A1)
		pos = readUint64(s, pos+1, &B1)
		pos = readUint64(s, pos+1, &C1)
		pos = readUint64(s, pos+1, &M1)
		res := solve(N, V, P, A0, B0, C0, M0, Q, A1, B1, C1, M1)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

const MOD = 1000000007

func solve(N int, V uint64, P []uint64, A0, B0, C0 uint64, M0 uint64, Q []uint64, A1, B1, C1 uint64, M1 uint64) int {
	if V%MOD == 0 {
		return 0
	}

	A0 = (A0 * A0) % M0
	A1 = (A1 * A1) % M1
	B0 %= M0
	B1 %= M1
	C0 %= M0
	C1 %= M1
	PP := make([]uint64, N)
	copy(PP, P)
	QQ := make([]uint64, N)
	copy(QQ, Q)
	var power uint64 = 1
	for i := 0; i < N; i++ {
		if i > 1 {
			PP[i] = (((A0*PP[i-1])%M0+(B0*PP[i-2])%M0)%M0 + C0) % M0
			QQ[i] = (((A1*QQ[i-1])%M1+(B1*QQ[i-2])%M1)%M1 + C1) % M1
		}
		dim := PP[i]*M1 + QQ[i] + 1
		dim--
		if dim == 0 {
			return 1
		}

		dim %= MOD - 1
		power = (power * dim) % (MOD - 1)
	}

	return int(fastPow(V%MOD, power))
}

func fastPow(a, b uint64) uint64 {
	var res uint64 = 1
	for b != 0 {
		if b&1 == 1 {
			res = (res * a) % MOD
		}
		a = (a * a) % MOD
		b >>= 1
	}
	return res
}
