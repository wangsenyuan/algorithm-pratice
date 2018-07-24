package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
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

func readFloat64(bs []byte, from int, val *float64) int {
	i := from
	sign := 1
	if bs[i] == '-' {
		sign = -1
		i++
	}

	var dec float64
	for i < len(bs) && bs[i] != '.' && bs[i] != ' ' {
		dec = dec*10 + float64(bs[i]-'0')
		i++
	}
	*val = dec

	if i == len(bs) || bs[i] == ' ' {
		//no fraction
		return i
	}
	i++
	var frac float64
	base := 1.0
	for i < len(bs) && bs[i] != ' ' {
		frac = frac*10 + float64(bs[i]-'0')
		base *= 10
		i++
	}
	*val += frac / base
	return i * sign
}

func readNFloat64s(scanner *bufio.Scanner, n int) []float64 {
	res := make([]float64, n)

	pos := 0
	scanner.Scan()
	bs := scanner.Bytes()
	//fmt.Printf("[debug] %s\n", string(bs))
	for i := 0; i < n; i++ {
		for pos < len(bs) && bs[pos] == ' ' {
			pos++
		}

		pos = readFloat64(bs, pos, &res[i])
	}
	return res
}

func main() {
	f, err := os.Open("./D-large-practice.in")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	tc := readNum(scanner)

	for i := 1; i <= tc; i++ {
		R, C := readTwoNums(scanner)
		G := make([][]byte, R)
		for j := 0; j < R; j++ {
			scanner.Scan()
			G[j] = scanner.Bytes()
		}
		res := solve(R, C, G)
		if res {
			fmt.Printf("Case #%d: A\n", i)
		} else {
			fmt.Printf("Case #%d: B\n", i)
		}
	}
}

func solve(R, C int, G [][]byte) bool {

	var doit func(r, c, b int) int
	N := 1 << uint(C+1)

	memo := make([][][]int, R)
	for i := 0; i < R; i++ {
		memo[i] = make([][]int, C)
		for j := 0; j < C; j++ {
			memo[i][j] = make([]int, N)
			for k := 0; k < N; k++ {
				memo[i][j][k] = -1
			}
		}
	}

	doit = func(r, c, b int) int {
		if c == C {
			c = 0
			r++
			if r == R {
				return 0
			}
		}
		if memo[r][c][b] >= 0 {
			return memo[r][c][b]
		}
		b2 := (b << 1) & (N - 1)
		var ret int
		if G[r][c] != '.' {
			ret = doit(r, c+1, b2)
		} else {
			// no match with previous cells
			ret = doit(r, c+1, b2+1)
			if c > 0 && b&1 == 1 {
				// match with left cell
				ret = max(ret, 1+doit(r, c+1, b2-2))
			}

			if c > 0 && (b&(1<<uint(C))) > 0 {
				// match with top-left cell, no affect with b2
				ret = max(ret, 1+doit(r, c+1, b2))
			}

			if b&(1<<uint(C-1)) > 0 {
				// match with top cell,
				ret = max(ret, 1+doit(r, c+1, b2-(1<<uint(C))))
			}

			if c < C-1 && b&(1<<uint(C-2)) > 0 {
				// match with top-right cell
				ret = max(ret, 1+doit(r, c+1, b2-(1<<uint(C-1))))
			}
		}
		memo[r][c][b] = ret
		return ret
	}

	var kr, kc int
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if G[i][j] == 'K' {
				kr, kc = i, j
			}
		}
	}

	// when no king in the matching
	a := doit(0, 0, 0)
	G[kr][kc] = '.'
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			for k := 0; k < N; k++ {
				memo[i][j][k] = -1
			}
		}
	}
	// king in the matching
	b := doit(0, 0, 0)
	return b > a
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
