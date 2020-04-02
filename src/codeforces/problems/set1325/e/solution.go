package main

import (
	"bufio"
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	A := readNNums(reader, n)

	fmt.Println(solve(n, A))
}

const MAX_A = 1000003

var primes []int
var lp []int
var d [][]int

func init() {
	lp = make([]int, MAX_A)
	primes = make([]int, MAX_A)
	d = make([][]int, MAX_A)
	var n int

	for i := 2; i < MAX_A; i++ {
		if lp[i] == 0 {
			primes[n] = i
			n++
			for j := i; j < MAX_A; j += i {
				lp[j] = i
			}
		}
		if lp[i] == i {
			d[i] = []int{i}
		} else {
			var k int
			for k < len(d[i/lp[i]]) && d[i/lp[i]][k] != lp[i] {
				k++
			}
			if k < len(d[i/lp[i]]) {
				// duplicate
				d[i] = make([]int, len(d[i/lp[i]])-1)
				copy(d[i], d[i/lp[i]][:k])
				copy(d[i][k:], d[i/lp[i]][k+1:])
			} else {
				d[i] = make([]int, len(d[i/lp[i]])+1)
				copy(d[i], d[i/lp[i]])
				d[i][k] = lp[i]
			}
		}
	}

	primes = primes[:n]
}

func solve(n int, A []int) int {
	E := make([][]int, n)
	V := make([][]int, MAX_A)

	for i := 0; i < MAX_A; i++ {
		V[i] = make([]int, 0, 10)
	}

	for i := 0; i < n; i++ {
		a := A[i]

		if len(d[a]) == 0 {
			// it is already a perfect square number
			return 1
		}
		if len(d[a]) == 1 {
			// it is a prime
			E[i] = []int{d[a][0], 1}
		} else {
			// it has to have only two valid prime factors, because the 7 divisors limitation
			E[i] = []int{d[a][0], d[a][1]}
		}

		V[E[i][0]] = append(V[E[i][0]], i)
		V[E[i][1]] = append(V[E[i][1]], i)
	}

	ans := math.MaxInt32

	dist := make([]int, MAX_A)

	que := make([]Pair, 2*len(primes))

	for _, i := range primes {
		if int64(i)*int64(i) > MAX_A {
			break
		}

		for _, j := range primes {
			dist[j] = 0
		}

		var front int
		var end int

		for _, j := range V[i] {
			que[end] = Pair{j, E[j][0] == i}
			end++
			dist[E[j][0]^E[j][1]^i] = 1
		}

		for front < end {
			cur := que[front]
			front++
			node := E[cur.first][cur.GetSecondAsInt()]

			for _, u := range V[node] {
				if u != cur.first {
					next := Pair{u, E[u][0] == node}
					nnode := E[next.first][next.GetSecondAsInt()]

					if dist[nnode] == 0 && nnode != i {
						que[end] = next
						end++
						dist[nnode] = dist[node] + 1
					} else {
						ans = min(ans, dist[node]+dist[nnode]+1)
					}
				}
			}
		}

	}

	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Pair struct {
	first  int
	second bool
}

func (p Pair) GetSecondAsInt() int {
	if p.second {
		return 1
	}
	return 0
}
