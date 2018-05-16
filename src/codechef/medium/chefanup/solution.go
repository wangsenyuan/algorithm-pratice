package main

import (
	"sort"
	"bufio"
	"os"
	"fmt"
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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	tmp := int64(0)
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
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
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	t := readNum(scanner)

	for t > 0 {
		t--
		scanner.Scan()
		var n, K int
		pos := readInt(scanner.Bytes(), 0, &n)
		pos = readInt(scanner.Bytes(), pos+1, &K)
		var L int64
		readInt64(scanner.Bytes(), pos+1, &L)
		res := solve(n, K, L)
		fmt.Printf("%d", res[0])
		for i := 1; i < n; i++ {
			fmt.Printf(" %d", res[i])
		}
		fmt.Println()
	}
}

func solve(N int, K int, L int64) []int {
	res := make([]int, N)
	X := int64(K)
	i := N - 1
	L--
	for L > 0 {
		res[i] = int(L % X)
		i--
		L /= X
	}

	for i := 0; i < N; i++ {
		res[i]++
	}
	return res
}
func solve1(N int, K int, L int64) []int {
	INF := L
	kn := make([]int64, N+1)

	kn[0] = 1
	for i := 1; i <= N; i++ {
		if kn[i-1] == INF {
			kn[i] = INF
		} else {
			kn[i] = kn[i-1] * int64(K)
			if kn[i] < 0 || kn[i] > INF {
				kn[i] = INF
			}
		}
	}
	L--
	res := make([]int, N)
	var last int
	for L > 0 {
		m := sort.Search(N+1, func(i int) bool {
			return kn[i] > L
		})

		//m the length of continus K from right
		//kn[m] > L, kn[m-1] <= L
		m--
		left := N - m - 1
		for i := last; i < left; i++ {
			res[i] = 1
		}

		x := 1
		for x <= K {
			y := int64(x)
			if y*kn[m] > L {
				break
			}
			x++
		}
		// x * kn[m] >= L
		res[left] = x
		L -= int64(x-1) * kn[m]
		last = left + 1
	}

	for i := last; i < N; i++ {
		res[i] = 1
	}

	return res
}
