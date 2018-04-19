package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bs []byte, from int, val *int) int {
	tmp := 0
	i := from
	for i < len(bs) && bs[i] != ' ' {
		tmp = tmp*10 + int(bs[i]-'0')
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

func readNNums(scanner *bufio.Scanner, n int) []int {
	scanner.Scan()
	bs := scanner.Bytes()
	res := make([]int, n)
	x := -1
	for i := 0; i < n; i++ {
		x = readInt(bs, x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	t := readNum(scanner)

	for t > 0 {
		first := readNNums(scanner, 2)
		n, k := first[0], first[1]
		board := make([][]byte, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			board[i] = scanner.Bytes()
		}
		res := solve(n, k, board)
		if res {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

		t--
	}
}
func solve(n, K int, board [][]byte) bool {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// row, column diagonal, reverse-diagonal
			var rl, cl, ld, rd int
			var rc, cc, ac, bc int
			for k := 0; k < K; k++ {
				if i+k < n && (board[i+k][j] == 'X' || (board[i+k][j] == '.' && cc == 0)) {
					cl++
					if board[i+k][j] == '.' {
						cc++
					}
				}
				if j+k < n && (board[i][j+k] == 'X' || (board[i][j+k] == '.' && rc == 0)) {
					rl++
					if board[i][j+k] == '.' {
						rc++
					}
				}

				if i+k < n && j+k < n && (board[i+k][j+k] == 'X' || (board[i+k][j+k] == '.' && ac == 0)) {
					ld++
					if board[i+k][j+k] == '.' {
						ac++
					}
				}

				if i+k < n && j-k >= 0 && (board[i+k][j-k] == 'X' || (board[i+k][j-k] == '.' && bc == 0)) {
					rd++
					if board[i+k][j-k] == '.' {
						bc++
					}
				}
			}
			if cl == K || rl == K || ld == K || rd == K {
				return true
			}
		}
	}
	return false
}

func solve1(n, k int, board [][]byte) bool {

	flags := make([][][]int, n)
	for i := 0; i < n; i++ {
		flags[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			flags[i][j] = make([]int, 8)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				flags[i][j][0] = 1
				flags[i][j][1] = 1
				flags[i][j][2] = 1
				if i > 0 {
					flags[i][j][0] += flags[i-1][j][0]
				}
				if j > 0 {
					flags[i][j][1] += flags[i][j-1][1]
				}
				if i > 0 && j > 0 {
					flags[i][j][2] += flags[i-1][j-1][2]
				}
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if board[i][j] == 'X' {
				flags[i][j][3] = 1
				flags[i][j][4] = 1
				flags[i][j][5] = 1
				if i < n-1 {
					flags[i][j][3] += flags[i+1][j][3]
				}
				if j < n-1 {
					flags[i][j][4] += flags[i][j+1][4]
				}
				if i < n-1 && j < n-1 {
					flags[i][j][5] += flags[i+1][j+1][5]
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := n - 1; j >= 0; j-- {
			if board[i][j] == 'X' {
				flags[i][j][6] = 1
				if i > 0 && j < n-1 {
					flags[i][j][6] += flags[i-1][j+1][6]
				}
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				flags[i][j][7] = 1
				if i < n-1 && j > 0 {
					flags[i][j][7] += flags[i+1][j-1][7]
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '.' {
				var a, b, c, d int
				if i > 0 {
					a += flags[i-1][j][0]
				}
				if i < n-1 {
					a += flags[i+1][j][3]
				}
				if j > 0 {
					b += flags[i][j-1][1]
				}
				if j < n-1 {
					b += flags[i][j+1][4]
				}
				if i > 0 && j > 0 {
					c += flags[i-1][j-1][2]
				}
				if i < n-1 && j < n-1 {
					c += flags[i+1][j+1][5]
				}
				if i > 0 && j < n-1 {
					d += flags[i-1][j+1][6]
				}
				if i < n-1 && j > 0 {
					d += flags[i+1][j-1][7]
				}
				if a+1 >= k || b+1 >= k || c+1 >= k || d+1 >= k {
					return true
				}
			}
		}
	}

	return false
}
