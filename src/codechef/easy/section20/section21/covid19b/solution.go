package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		V := readNNums(reader, n)
		a, b := solve(n, V)

		fmt.Printf("%d %d\n", a, b)
	}
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

func solve(n int, V []int) (int, int) {

	// lets fist find meet[i][j] = when they will meet
	meet := make([][]Pair, n)
	for i := 0; i < n; i++ {
		meet[i] = make([]Pair, n)
		for j := 0; j < n; j++ {
			meet[i][j] = Pair{-1, -1}
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if V[j] >= V[i] {
				//safe
				continue
			}
			// i + V[i] * t = j + V[j] * t
			// (V[i] - V[j]) * t = j - i
			a := j - i
			b := V[i] - V[j]
			c := gcd(a, b)
			meet[i][j] = Pair{a / c, b / c}
			meet[j][i] = Pair{a / c, b / c}
		}
	}

	find := func(i int) int {
		//if i inffect
		set := make(map[int]Pair)
		set[i] = Pair{0, 1}

		for {
			var found = -1
			var time Pair
			for k, p := range set {
				for i := 0; i < n; i++ {
					if i == k || meet[i][k].first < 0 {
						continue
					}
					if _, found := set[i]; found {
						continue
					}
					q := meet[i][k]
					if p.Less(q) {
						found = i
						time = q
						break
					}
				}
				if found >= 0 {
					break
				}
			}
			if found < 0 {
				break
			}
			set[found] = time
		}

		return len(set)
	}

	var best = n
	var worst int

	for i := 0; i < n; i++ {
		cnt := find(i)
		if cnt < best {
			best = cnt
		}
		if cnt > worst {
			worst = cnt
		}
	}
	return best, worst
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

type Pair struct {
	first  int
	second int
}

func (this Pair) Less(that Pair) bool {
	return this.first*that.second <= this.second*that.first
}
