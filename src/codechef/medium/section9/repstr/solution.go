package main

import (
	"bufio"
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

func fillNNums(bs []byte, n int, res []int) {
	x := 0
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		L, H := readTwoNums(reader)

		if H == 0 {
			break
		}
		S, _ := reader.ReadString('\n')
		S = S[:len(S)-1]
		a, b := solve(L, H, S)
		fmt.Printf("%d %d\n", a, b)
	}
}

const BASE = 27
const BASE2 = 28
const MOD = 1000133
const MOD2 = 1000213

func solve(L, H int, S string) (int, int) {

	N := len(S)
	if N < L {
		return 0, 0
	}

	find := func(K int) int {
		if K > N {
			return 0
		}
		if K == N {
			return 1
		}

		mem := make(map[int64]map[int64]int)

		var hash1, hash2 int64
		var b1, b2 int64 = 1, 1
		for i := 0; i < K; i++ {
			hash1 = hash1*BASE + int64(S[i]-'a'+1)
			hash1 %= MOD
			hash2 = hash2*BASE2 + int64(S[i]-'a'+1)
			hash2 %= MOD2

			b1 *= BASE
			b1 %= MOD
			b2 *= BASE2
			b2 %= MOD2
		}

		updateCount(mem, hash1, hash2)

		for i := K; i < N; i++ {
			hash1 = hash1*BASE + int64(S[i]-'a'+1)
			hash1 %= MOD

			hash1 -= (b1 * int64(S[i-K]-'a'+1)) % MOD
			if hash1 < 0 {
				hash1 += MOD
			}

			hash2 = hash2*BASE2 + int64(S[i]-'a'+1)
			hash2 %= MOD2
			hash2 -= (b2 * int64(S[i-K]-'a'+1)) % MOD2
			if hash2 < 0 {
				hash2 += MOD
			}

			updateCount(mem, hash1, hash2)
		}

		var res int

		for _, v := range mem {
			for _, c := range v {
				if res < c {
					res = c
				}
			}
		}

		return res
	}

	cnt0 := find(L)

	var left = L
	var right = H + 1

	for left < right {
		mid := (left + right) / 2
		cnt := find(mid)
		if cnt < cnt0 {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return cnt0, right - 1
}

func updateCount(mem map[int64]map[int64]int, a, b int64) {
	if _, found := mem[a]; !found {
		mem[a] = make(map[int64]int)
	}
	mem[a][b]++
}

func pow(a, b int) int {
	A := int64(a)
	R := int64(1)
	for b > 0 {
		if b&1 == 1 {
			R = (R * A) % MOD
		}
		A = (A * A) % MOD
		b >>= 1
	}
	return int(R)
}
