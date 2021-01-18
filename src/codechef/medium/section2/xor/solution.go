package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, k := readTwoNums(reader)
	A := make([]int, n)
	for i := 0; i < n; i++ {
		A[i] = readNum(reader)
	}
	res := solve(n, k, A)
	var buf bytes.Buffer

	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	fmt.Println(buf.String())
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

const MAX_K = 250000

type Pair struct {
	first, second int
}

const H = 31

func solve(n int, k int, A []int) []int {
	sort.Ints(A)

	middle := func(L, R, h int) int {
		M := L
		for M < R && A[M]&(1<<uint(h)) == 0 {
			M++
		}
		return M
	}

	C2 := func(l int) int {
		L := int64(l)
		V := L * (L - 1) / 2
		if V > MAX_K {
			return MAX_K + 1
		}
		return int(V)
	}

	mult := func(a, b int) int {
		v := int64(a) * int64(b)
		if v > MAX_K {
			return MAX_K + 1
		}
		return int(v)
	}

	mid := make([]int, n)
	midB := make([]int, n)
	midC := make([]int, n)

	B := make([]Pair, n)
	C := make([]Pair, n)
	B[0] = Pair{0, n}
	bLen := 1

	NB := make([]Pair, n)
	NC := make([]Pair, n)

	res := make([]int, 0, k)

	var cross bool

	for h := H - 1; h >= 0; h-- {
		if !cross {
			var cnt0 int
			for i := 0; i < bLen; i++ {
				L, R := B[i].first, B[i].second
				M := middle(L, R, h)
				mid[i] = M
				if cnt0 < k {
					cnt0 += C2(R-M) + C2(M-L)
				}
			}

			var nbLen int

			for i := 0; i < bLen; i++ {
				L, R := B[i].first, B[i].second
				M := mid[i]
				if cnt0 >= k {
					if L < M {
						NB[nbLen] = Pair{L, M}
						nbLen++
					}
					if M < R {
						NB[nbLen] = Pair{M, R}
						nbLen++
					}
				} else {
					for u := L; u < M; u++ {
						for v := L; v < u; v++ {
							res = append(res, A[u]^A[v])
						}
					}
					for u := M; u < R; u++ {
						for v := M; v < u; v++ {
							res = append(res, A[u]^A[v])
						}
					}
					if L < M && M < R {
						NB[nbLen] = Pair{L, M}
						C[nbLen] = Pair{M, R}
						nbLen++
					}
				}
			}

			if cnt0 < k {
				cross = true
				k -= cnt0
			}
			bLen = nbLen
			copy(B[:bLen], NB[:bLen])
		} else {
			var cnt0 int
			for i := 0; i < bLen; i++ {
				L1, R1 := B[i].first, B[i].second
				M1 := middle(L1, R1, h)
				midB[i] = M1
				L2, R2 := C[i].first, C[i].second
				M2 := middle(L2, R2, h)
				midC[i] = M2
				if cnt0 < k {
					cnt0 += mult(M1-L1, M2-L2) + mult(R1-M1, R2-M2)
				}
			}

			var nbLen int
			for i := 0; i < bLen; i++ {
				L1, R1 := B[i].first, B[i].second
				M1 := midB[i]
				L2, R2 := C[i].first, C[i].second
				M2 := midC[i]

				if cnt0 >= k {
					if L1 < M1 && L2 < M2 {
						NB[nbLen] = Pair{L1, M1}
						NC[nbLen] = Pair{L2, M2}
						nbLen++
					}
					if M1 < R1 && M2 < R2 {
						NB[nbLen] = Pair{M1, R1}
						NC[nbLen] = Pair{M2, R2}
						nbLen++
					}
				} else {
					for u := L1; u < M1; u++ {
						for v := L2; v < M2; v++ {
							res = append(res, A[u]^A[v])
						}
					}
					for u := M1; u < R1; u++ {
						for v := M2; v < R2; v++ {
							res = append(res, A[u]^A[v])
						}
					}
					if L1 < M1 && L2 < M2 {
						NB[nbLen] = Pair{L1, M1}
						NC[nbLen] = Pair{L2, M2}
						nbLen++
					}
					if M1 < R1 && M2 < R2 {
						NB[nbLen] = Pair{M1, R1}
						NC[nbLen] = Pair{M2, R2}
						nbLen++
					}
				}
			}

			if cnt0 < k {
				k -= cnt0
			}
			bLen = nbLen
			copy(B[:bLen], NB[:bLen])
			copy(C[:bLen], NC[:bLen])
		}
	}

	var val int
	if cross {
		val = A[B[0].first] ^ A[C[0].first]
	}
	for k > 0 {
		res = append(res, val)
		k--
	}
	sort.Ints(res)

	return res
}
