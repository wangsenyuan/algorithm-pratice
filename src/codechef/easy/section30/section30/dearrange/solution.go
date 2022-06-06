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

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
		buf.WriteString(fmt.Sprintf("%d\n", len(res)))
		for _, op := range res {
			buf.WriteString(fmt.Sprintf("%d %d\n", op.l, op.r))
			for i := 0; i < len(op.P); i++ {
				buf.WriteString(fmt.Sprintf("%d ", op.P[i]))
			}
			buf.WriteByte('\n')
		}
	}
	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

var perms [][][]int

func init() {
	perms = make([][][]int, 4)
	for i := 4; i < 8; i++ {
		perms[i-4] = permutations(i)
	}
}

func solve(A []int) []Operation {
	n := len(A)
	L := 1
	for L <= n && A[L-1] == L {
		L++
	}
	if L > n {
		return nil
	}
	// L is the first index that A[L-1] != L
	R := n
	for R > 0 && A[R-1] == R {
		R--
	}
	// R is the last A[R-1] != R
	p := L
	for p <= R && A[p-1] != p {
		p++
	}
	B := make([]int, n)
	copy(B, A)

	var ops []Operation

	if p > R {
		sort.Ints(B)
		ops = append(ops, Operation{L, R, B})
		return ops
	}

	if n == 3 {
		ops = append(ops, Operation{1, 2, []int{2, 3, 1}})
		ops = append(ops, Operation{1, 3, []int{1, 2, 3}})
	} else {

		place := func(l, r int) {
			le := r - l
			for _, perm := range perms[le-4] {
				for i, j := range perm {
					B[l+i] = A[l+j]
				}
				ok := true
				for i := 0; i < le; i++ {
					if B[l+i] == A[l+i] || B[l+i] == l+i+1 {
						ok = false
						break
					}
				}
				if ok {
					return
				}
			}
		}

		for i := 0; i < n; {
			j := i + 4
			if i+7 >= n {
				j = n
			}

			place(i, j)

			i = j
		}

		ops = append(ops, Operation{1, n, B})

		C := make([]int, n)
		copy(C, A)
		sort.Ints(C)

		ops = append(ops, Operation{1, n, C})
	}

	return ops
}

type Operation struct {
	l, r int
	P    []int
}

func permutations(n int) [][]int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	var res [][]int
	buf := make([]int, n)
	var loop func(i int, flag int)
	loop = func(i int, flag int) {
		if i == n {
			cur := make([]int, n)
			copy(cur, buf)
			res = append(res, cur)
			return
		}
		for j := 0; j < n; j++ {
			if (flag>>uint(j))&1 == 0 && arr[j] != i {
				buf[i] = arr[j]
				loop(i+1, flag|(1<<uint(j)))
			}
		}
	}

	loop(0, 0)

	return res
}
