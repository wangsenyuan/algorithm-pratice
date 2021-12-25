package main

import (
	"bufio"
	"container/list"
	"fmt"
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

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		line := readNNums(scanner, 3)
		N, M, Q := line[0], line[1], line[2]
		challeges := make([][]int, M)
		for i := 0; i < M; i++ {
			challeges[i] = readNNums(scanner, 3)
		}
		combos := make([][]int, Q)
		for i := 0; i < Q; i++ {
			combos[i] = readNNums(scanner, 2)
		}
		res := solve(N, M, Q, challeges, combos)

		for i := 0; i < N; i++ {
			if i < N-1 {
				fmt.Printf("%d ", res[i])
			} else {
				fmt.Printf("%d\n", res[i])
			}
		}
	}
}

func solve(N, M, Q int, challeges [][]int, combos [][]int) []int64 {
	ques := make([]*list.List, N+1)

	for i := 1; i <= N; i++ {
		ques[i] = list.New()
	}

	for i, ch := range challeges {
		i++
		l, r := ch[0], ch[1]
		ques[l].PushFront(i)
		if r+1 <= N {
			ques[r+1].PushFront(-i)
		}
	}

	V := make([]int64, Q)

	res := make([]int64, N)
	var sum int64

	for p := 1; p <= N; p++ {

		for e := ques[p].Front(); e != nil; e = e.Next() {
			x := e.Value.(int)
			if x > 0 {
				for i := 0; i < Q; i++ {
					if combos[i][0] <= x && x <= combos[i][1] {
						if V[i] > 0 {
							sum -= V[i]
						}

						V[i] += int64(challeges[x-1][2])

						if V[i] > 0 {
							sum += V[i]
						}
					}
				}
			} else {
				x = -x
				for i := 0; i < Q; i++ {
					if combos[i][0] <= x && x <= combos[i][1] {
						if V[i] > 0 {
							sum -= V[i]
						}

						V[i] -= int64(challeges[x-1][2])

						if V[i] > 0 {
							sum += V[i]
						}
					}
				}
			}
		}

		res[p-1] = sum
	}

	return res
}
