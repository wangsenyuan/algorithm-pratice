package main

import (
	"bufio"
	"fmt"
	"io"
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

	for scanner.Scan() {
		err := scanner.Err()
		if err == io.EOF {
			break
		}
		var n int
		readInt(scanner.Bytes(), 0, &n)

		mixtures := readNNums(scanner, n)

		fmt.Println(solve(n, mixtures))
	}
}

const INF = 1 << 30

func solve(n int, mixtures []int) int64 {
	// f[i][j] = least smoke after mixing [i, j]
	f := make([][]int64, n)
	// g[i][j] = color when get least smoke after mixing [i, j]
	g := make([][]int, n)

	for i := 0; i < n; i++ {
		f[i] = make([]int64, n)
		g[i] = make([]int, n)
		g[i][i] = mixtures[i]
		if i < n-1 {
			f[i][i+1] = int64(mixtures[i] * mixtures[i+1])
		}
		for j := i + 2; j < n; j++ {
			f[i][j] = INF
		}
		for j := i + 1; j < n; j++ {
			g[i][j] = (g[i][j-1] + mixtures[j]) % 100
		}
	}

	for k := 3; k <= n; k++ {
		for i := 0; i+k <= n; i++ {
			j := i + k - 1
			var sum int64 = INF
			for x := i; x < j; x++ {
				tmp := int64(g[i][x]*g[x+1][j]) + f[i][x] + f[x+1][j]
				if tmp < sum {
					sum = tmp
				}
			}
			f[i][j] = sum
		}
	}
	return f[0][n-1]
}
