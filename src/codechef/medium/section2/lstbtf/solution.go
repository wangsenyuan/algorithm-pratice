package main

import (
	"bufio"
	"bytes"
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n := readNum(scanner)
		fmt.Println(solve(n))
	}
}

const XX = 29

var can [][][]bool

var squares []int

func init() {
	can = make([][][]bool, 11)
	for i := 0; i < 11; i++ {
		can[i] = make([][]bool, XX)
		for j := 0; j < XX; j++ {
			can[i][j] = make([]bool, 81*XX)
		}
	}

	can[10][0][0] = true

	for sm := 0; sm < 81*XX; sm++ {
		for first := 9; first > 0; first-- {
			for cnt := 0; cnt < XX; cnt++ {
				if cnt == 0 {
					can[first][cnt][sm] = sm == 0
				} else {
					can[first][cnt][sm] = can[first+1][cnt][sm]

					for t := 1; t <= cnt && t*first*first <= sm; t++ {
						can[first][cnt][sm] = can[first][cnt][sm] || can[first+1][cnt-t][sm-t*first*first]
					}
				}
			}
		}
	}

	squares = make([]int, 10001)

	for i := 1; i <= len(squares); i++ {
		squares[i-1] = i * i
	}
}

func solve(n int) string {
	var res bytes.Buffer

	var curSum int

	for i := 1; n > 0 && i <= 9; i++ {
		cnt := n
		var success bool

		for !success {
			for j := 0; j < len(squares) && !success; j++ {
				x := squares[j]
				tmpSum := curSum + cnt*i*i

				if tmpSum <= x && x-tmpSum < XX*81 {
					success = can[i+1][n-cnt][x-tmpSum]
				}
			}
			if !success {
				cnt--
			}
		}

		n -= cnt

		curSum += cnt * i * i

		for cnt > 0 {
			res.WriteByte(byte('0' + i))
			cnt--
		}
	}

	return res.String()
}
