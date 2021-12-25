package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

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

func readTripNums(scanner *bufio.Scanner) (a int, b int, c int) {
	scanner.Scan()
	bs := scanner.Bytes()
	x := readInt(bs, 0, &a)
	x = readInt(bs, x+1, &b)
	readInt(bs, x+1, &c)
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	t := readNum(scanner)
	for i := 0; i < t; i++ {
		a, e, v := readTripNums(scanner)
		res := solve(a, e, v)
		fmt.Println(res)
	}
}

const MOD int64 = 1000000007

func solve(A, B, V int) int {
	f := make([][][]int64, A+B+1)
	for n := 0; n <= A+B; n++ {
		f[n] = make([][]int64, A+1)
		for a := 0; a <= A; a++ {
			f[n][a] = make([]int64, V+1)
			for v := 0; v <= V; v++ {
				f[n][a][v] = -1
			}
		}
	}
	//f[0][0][0] = 1

	var dp func(n, a, v int) int64
	dp = func(n, a, v int) int64 {
		if f[n][a][v] >= 0 {
			return f[n][a][v]
		}
		if n < a {
			f[n][a][v] = 0
		} else if a == 0 {
			//all E
			f[n][a][v] = 1
		} else if v < a {
			f[n][a][v] = 0
		} else if n == 1 {
			f[n][a][v] = 1
		} else {
			f[n][a][v] = 0
			nl, nh := n/2, n-n/2
			for aa := max(0, a-nh); aa <= nl && aa <= a; aa++ {
				p := (dp(nl, aa, v-a) * dp(nh, a-aa, v-a)) % MOD
				var q int64
				if v > a {
					q = (dp(nl, aa, v-a-1) * dp(nh, a-aa, v-a-1)) % MOD
				}
				f[n][a][v] = (f[n][a][v] + p - q + MOD) % MOD
			}
			f[n][a][v] = (f[n][a][v] + dp(n, a, v-1)) % MOD
		}

		return f[n][a][v]
	}

	N := A + B
	res := dp(N, A, V)
	if V > 0 {
		res = (res - dp(N, A, V-1) + MOD) % MOD
	}

	return int(res)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
