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
		n := readNum(scanner)
		W := make([][]int, n)
		for i := 0; i < n; i++ {
			W[i] = readNNums(scanner, n+1)
		}
		M := make([][]int, n)
		for i := 0; i < n; i++ {
			M[i] = readNNums(scanner, n+1)
		}

		res := solve(n, W, M)

		for i := 1; i <= n; i++ {
			fmt.Printf("%d %d\n", i, res[i-1])
		}

		t--
	}
}

func solve(n int, W [][]int, M [][]int) []int {
	ranks := make([][]int, n+1)
	for i := 0; i < n; i++ {
		id := W[i][0]
		ranks[id] = make([]int, n+1)
		for j := 1; j < len(W[i]); j++ {
			mid := W[i][j]
			ranks[id][mid] = j
		}
	}

	MP := make([][]int, n+1)
	for i := 0; i < n; i++ {
		id := M[i][0]
		MP[id] = make([]int, n+1)
		MP[id] = M[i]
		MP[id][0] = 1
	}

	pair := make([]int, n+1)
	rp := make([]int, n+1)
	que := make([]int, n)
	front, rear := 0, 0
	for i := 1; i <= n; i++ {
		que[rear] = i
		rear++
	}

	for front < rear {
		man := que[front]
		prefs := MP[man]
		i := prefs[0]
		prefs[0]++
		woman := prefs[i]
		if rp[woman] == 0 {
			// w is not married
			rp[woman] = man
			pair[man] = woman
			front++
		} else {
			anotherMan := rp[woman]
			if ranks[woman][anotherMan] > ranks[woman][man] {
				// current man is preferable to anotherman
				// change
				rp[woman] = man
				pair[man] = woman
				pair[anotherMan] = 0
				que[front] = anotherMan
				// poor another man, have to change gire frient
			}

			// poor man, try propose to the next
		}
	}

	return pair[1:]
}
