package main

import (
	"bufio"
	"fmt"
	"os"
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
		G := make([][]int, n)
		for i := 0; i < n; i++ {
			G[i] = readNNums(scanner, n)
		}

		res := solve(n, G)

		for i := 0; i < n; i++ {
			if i < n-1 {
				fmt.Printf("%d ", res[i])
			} else {
				fmt.Printf("%d\n", res[i])
			}
		}
	}
}

func solve(n int, G [][]int) []int {

	PP := make([]PS, n)
	for i := 0; i < n; i++ {
		PP[i] = make(PS, n)
		for j := 0; j < n; j++ {
			PP[i][j] = P{G[i][j], j}
		}
		sort.Sort(PP[i])
	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1
	}
	pos := make([]int, n)

	que := make([]int, n*n)
	var front, end int
	for i := 0; i < n; i++ {
		que[end] = i
		end++
	}

	for front < end {
		cur := que[front]
		front++

		p := PP[cur][pos[cur]]
		pos[cur]++
		// house index
		i := p.second
		if res[i] < 0 {
			res[i] = cur
		} else if G[cur][i] > G[res[i]][i] {
			que[end] = res[i]
			end++
			res[i] = cur
		} else {
			//try again
			front--
		}
	}
	for i := 0; i < n; i++ {
		res[i]++
	}
	return res
}

type P struct {
	first, second int
}

type PS []P

func (ps PS) Len() int {
	return len(ps)
}

func (ps PS) Less(i, j int) bool {
	return ps[i].first < ps[j].first
}

func (ps PS) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
