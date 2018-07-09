package main

import (
	"bufio"
	"fmt"
	"os"
	"bytes"
	"sort"
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
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		n, K := readTwoNums(scanner)
		scanner.Scan()
		word := scanner.Bytes()

		mat := make([][]float64, 26)
		for i := 0; i < 26; i++ {
			mat[i] = readNFloat64s(scanner, 26)
		}
		words := make([][]byte, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			words[i] = scanner.Bytes()
		}

		fmt.Printf("%.10f\n", solve(n, K, word, mat, words))

		tc--
	}
}

func solve(n int, K int, word []byte, mat [][]float64, words [][]byte) float64 {
	sort.Sort(ByteSlices(words))

	var j int
	for i := 1; i <= n; i++ {
		if i < n && bytes.Compare(words[i], words[i-1]) == 0 {
			continue
		}
		words[j] = words[i-1]
		j++
	}
	words = words[:j]
	mat = matrixPow(mat, K)

	calculate := func(dst []byte) float64 {
		if len(dst) != len(word) {
			return 0
		}

		res := 1.0

		for i := 0; i < len(dst); i++ {
			x, y := word[i], dst[i]
			a, b := int(x-'a'), int(y-'a')
			if a < 0 || a >= 26 || b < 0 || b >= 26 {
				res = 0.0
				break
			}
			res *= mat[a][b]
		}

		return res
	}

	var res float64

	for _, dst := range words {
		res += calculate(dst)
	}

	return res
}

func multiple(a, b [][]float64) [][]float64 {
	res := make([][]float64, 26)
	for i := 0; i < 26; i++ {
		res[i] = make([]float64, 26)
	}

	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			for k := 0; k < 26; k++ {
				res[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return res
}

func matrixPow(a [][]float64, n int) [][]float64 {
	if n == 0 {
		return unitMatrix()
	}

	b := matrixPow(a, n>>1)
	c := multiple(b, b)

	if n&1 == 1 {
		return multiple(a, c)
	}
	return c
}

func unitMatrix() [][]float64 {
	res := make([][]float64, 26)
	for i := 0; i < 26; i++ {
		res[i] = make([]float64, 26)
		res[i][i] = 1
	}
	return res
}

func solve1(n int, K int, word []byte, mat [][]float64, words [][]byte) float64 {

	cache := make([][][]float64, 26)
	for i := 0; i < 26; i++ {
		cache[i] = make([][]float64, 26)
		for j := 0; j < 26; j++ {
			cache[i][j] = make([]float64, K)
			for k := 0; k < K; k++ {
				cache[i][j][k] = -1.0
			}
		}
	}

	var dfs func(cur int, dst int, k int, p float64) float64

	dfs = func(cur int, dst int, k int, p float64) float64 {
		if k == K {
			if cur == dst {
				return p
			}
			return 0.0
		}

		if cache[cur][dst][k] >= 0.0 {
			return cache[cur][dst][k]
		}

		if p == 0.0 {
			return 0.0
		}

		var ans float64

		for i := 0; i < 26; i++ {
			ans += dfs(i, dst, k+1, p*mat[cur][i])
		}
		cache[cur][dst][k] = ans

		return ans
	}

	cal := func(dst []byte) float64 {
		if len(dst) != len(word) {
			return 0.0
		}
		var res float64 = 1.0
		for i := 0; i < len(word); i++ {
			x := word[i]
			y := dst[i]
			if x < 'a' || x > 'z' || y < 'a' || y > 'z' {
				return 0.0
			}
			res *= dfs(int(x-'a'), int(y-'a'), 0, 1.0)
			if res == 0.0 {
				return 0.0
			}
		}
		return res
	}

	var res float64

	for i := 0; i < n; i++ {
		dstWord := words[i]
		res += cal(dstWord)
	}

	return res
}

type ByteSlices [][]byte

func (bs ByteSlices) Len() int {
	return len(bs)
}

func (bs ByteSlices) Less(i, j int) bool {
	return bytes.Compare(bs[i], bs[j]) < 0
}

func (bs ByteSlices) Swap(i, j int) {
	bs[i], bs[j] = bs[j], bs[i]
}
