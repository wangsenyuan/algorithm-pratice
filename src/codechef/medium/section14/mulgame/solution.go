package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
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

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, q, m := readThreeNums(reader)
		A := readNNums(reader, n)
		queries := make([][]int, q)
		for i := 0; i < q; i++ {
			queries[i] = readNNums(reader, 2)
		}
		res := solve(n, q, m, A, queries)
		//fmt.Println(res)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

func solve(N int, Q int, M int, A []int, queries [][]int) int {
	cnt := make([]int, M+1)

	ops := make(map[[2]int]int)

	for _, q := range queries {
		l, r := q[0], q[1]
		l--
		r--
		op := [...]int{A[l] + A[r], A[l]}
		ops[op]++
	}

	for k, w := range ops {
		u, v := k[0], k[1]
		for j := 0; j <= M; j += u {
			if j+v <= M {
				cnt[j+v] += w
			}
			if j+u <= M {
				cnt[j+u] -= w
			}
		}
	}

	ans := cnt[0]

	for i := 1; i <= M; i++ {
		cnt[i] += cnt[i-1]
		ans = max(ans, cnt[i])
	}
	return ans
}

func solve1(N int, Q int, M int, A []int, queries [][]int) int {
	C := int(math.Sqrt(float64(M)))

	was := make([]bool, C+1)
	cnt := make([]int, M+1)
	F := make([][]int, C+1)
	for i := 0; i <= C; i++ {
		F[i] = make([]int, C+1)
	}

	naive := func(l, r int) {
		md := l + r
		for i := 0; i <= M; i += md {
			if i+l <= md {
				cnt[i+l]++
			}
			if i+md <= M {
				cnt[i+md]--
			}
		}
	}

	todo := make([]int, 0, Q)

	for i := 0; i < Q; i++ {
		l, r := queries[i][0], queries[i][1]
		l--
		r--
		md := A[l] + A[r]
		if md > C {
			naive(A[l], A[r])
		} else {
			if !was[md] {
				todo = append(todo, md)
			}
			was[md] = true
			F[md][A[l]]++
			F[md][0]--
		}
	}

	for i := 1; i <= M; i++ {
		for _, x := range todo {
			cnt[i] += F[x][i%x]
		}
	}

	ans := cnt[0]

	for i := 1; i <= M; i++ {
		cnt[i] += cnt[i-1]
		ans = max(ans, cnt[i])
	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
