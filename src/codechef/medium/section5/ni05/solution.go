package main

import (
	"bufio"
	"os"
	"fmt"
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
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		scanner.Scan()
		want := readNNums(scanner, 10)
		fmt.Println(solve(nil, want))
	}
}

var matrix [][]int
var mem [][][][][][][][][][]int

func init() {
	matrix = [][]int{
		{55, 60, 4, 25, 18, 10, 12, 8, 11, 50},
		{60, 45, 75, 23, 27, 20, 24, 7, 33, 12},
		{4, 75, 78, 32, 36, 30, 36, 6, 12, 65},
		{25, 23, 32, 15, 45, 40, 48, 5, 14, 23},
		{18, 27, 36, 45, 54, 50, 60, 4, 15, 12},
		{10, 20, 30, 40, 50, 60, 72, 3, 32, 34},
		{12, 24, 36, 48, 60, 72, 84, 2, 23, 34},
		{8, 7, 6, 5, 4, 3, 2, 1, 34, 123},
		{11, 33, 12, 14, 15, 32, 23, 34, 65, 48},
		{50, 12, 65, 23, 12, 34, 34, 123, 48, 71},
	}

	mem = make([][][][][][][][][][]int, 3)

	for i0 := 0; i0 < 3; i0++ {
		mem[i0] = make([][][][][][][][][]int, 3)
		for i1 := 0; i1 < 3; i1++ {
			mem[i0][i1] = make([][][][][][][][]int, 3)
			for i2 := 0; i2 < 3; i2++ {
				mem[i0][i1][i2] = make([][][][][][][]int, 3)
				for i3 := 0; i3 < 3; i3++ {
					mem[i0][i1][i2][i3] = make([][][][][][]int, 3)
					for i4 := 0; i4 < 3; i4++ {
						mem[i0][i1][i2][i3][i4] = make([][][][][]int, 3)
						for i5 := 0; i5 < 3; i5++ {
							mem[i0][i1][i2][i3][i4][i5] = make([][][][]int, 3)
							for i6 := 0; i6 < 3; i6++ {
								mem[i0][i1][i2][i3][i4][i5][i6] = make([][][]int, 3)
								for i7 := 0; i7 < 3; i7++ {
									mem[i0][i1][i2][i3][i4][i5][i6][i7] = make([][]int, 3)
									for i8 := 0; i8 < 3; i8++ {
										mem[i0][i1][i2][i3][i4][i5][i6][i7][i8] = make([]int, 3)
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func dp(a []int) int {
	if mem[a[0]][a[1]][a[2]][a[3]][a[4]][a[5]][a[6]][a[7]][a[8]][a[9]] == 0 {
		var sum int
		for i := 0; i < 10; i++ {
			sum += a[i]
		}
		var ans int
		if sum == 2 {
			j, k := -1, -1
			for i := 0; i < 10; i++ {
				if a[i] == 0 {
					continue
				}
				if j < 0 {
					j = i
				} else {
					k = i
				}
			}
			if k < 0 {
				//only one
				ans = matrix[j][j]
			} else {
				ans = matrix[j][k]
			}
		} else {
			b := make([]int, 10)
			copy(b, a)
			for i := 0; i < 10; i++ {
				if b[i] > 0 {
					b[i]--
					for j := 0; j < 10; j++ {
						if b[j] > 0 {
							b[j]--
							ans = max(ans, dp(b)+matrix[i][j])
							b[j]++
						}
					}
					b[i]++
				}
			}
		}
		mem[a[0]][a[1]][a[2]][a[3]][a[4]][a[5]][a[6]][a[7]][a[8]][a[9]] = ans
	}

	return mem[a[0]][a[1]][a[2]][a[3]][a[4]][a[5]][a[6]][a[7]][a[8]][a[9]]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func solve(count []int, want []int) int {
	a := make([]int, 10)
	for i := 0; i < 10; i++ {
		if want[i] == 1 {
			a[i] = 2
		}
	}
	return dp(a)
}
