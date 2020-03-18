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
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func main() {

	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()

	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n := readNum(scanner)
		P := readNNums(scanner, 3)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(scanner, 2)
		}
		res := solve(n, P, E)
		f.WriteString(fmt.Sprintf("%d\n", res))
	}
}

func solve(n int, P []int, E [][]int) int64 {
	conns := make([][]int, n)

	for i := 0; i < n; i++ {
		conns[i] = make([]int, 0, 3)
	}

	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		conns[u] = append(conns[u], v)
		conns[v] = append(conns[v], u)
	}

	in := make([]int, n)
	out := make([]int, n)
	parent := make([]int, n)

	var time int
	var euler func(p, u int)

	euler = func(p, u int) {
		parent[u] = p
		time++
		in[u] = time

		for _, v := range conns[u] {
			if p == v {
				continue
			}
			euler(u, v)
		}
		out[u] = time
	}

	euler(-1, 0)

	bit := NewBIT(n)

	A := make([]int, n)
	B := make([]int, n)

	calc := func(p, q []int) int64 {
		var sum int64
		for i := 0; i < len(q); i++ {
			sum += int64(q[i])
		}
		var res int64
		for i := 0; i < len(p); i++ {
			x := int64(p[i])
			y := sum - int64(q[i])
			res += x * y
		}
		return res
	}

	var ans int64

	for u := 0; u < n; u++ {
		var ii int

		for _, v := range conns[u] {
			if v == parent[u] {
				continue
			}
			A[ii] = bit.GetRange(in[v], out[v])
			B[ii] = out[v] - in[v] + 1 - A[ii]
			ii++
		}

		A[ii] = bit.GetRange(1, in[u]-1) + bit.GetRange(out[u]+1, n)
		B[ii] = n - (out[u] - in[u] + 1) - A[ii]
		ii++
		if P[1] == 1 {
			ans += calc(B[:ii], B[:ii]) / 2
		} else if P[1] == 2 {
			ans += calc(A[:ii], B[:ii])
		} else {
			ans += calc(A[:ii], A[:ii]) / 2
		}

		bit.Set(in[u])
	}

	return ans
}

type BIT struct {
	arr []int
	n   int
}

func NewBIT(n int) BIT {
	arr := make([]int, n+1)
	return BIT{arr, n}
}

func (bit *BIT) Set(pos int) {
	// pos++
	for pos <= bit.n {
		bit.arr[pos] += 1
		pos += pos & -pos
	}
}

func (bit BIT) Get(pos int) int {
	var res int
	// pos++
	for pos > 0 {
		res += bit.arr[pos]
		pos = pos & (pos - 1)
	}
	return res
}

func (bit BIT) GetRange(left, right int) int {
	return bit.Get(right) - bit.Get(left-1)
}
