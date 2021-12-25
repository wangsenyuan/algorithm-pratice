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

	P, S := readTwoNums(scanner)

	tasks := make([][]int, 2*P)
	for i := 0; i < P; i++ {
		tasks[i<<1] = readNNums(scanner, S)
		tasks[i<<1|1] = readNNums(scanner, S)
	}
	res := solve(P, S, tasks)

	for _, ans := range res {
		fmt.Println(ans)
	}
}

func solve(P, S int, tasks [][]int) []int {
	nis := make(Pairs, P)

	count := func(index int) int {
		ps := make(Pairs, S)
		for i := 0; i < S; i++ {
			ps[i] = Pair{tasks[index<<1][i], tasks[index<<1|1][i]}
		}
		sort.Sort(ps)
		var cnt int
		for i := 0; i < S-1; i++ {
			if ps[i].second > ps[i+1].second {
				cnt++
			}
		}
		return cnt
	}

	for i := 0; i < P; i++ {
		n := count(i)
		nis[i] = Pair{n, i}
	}

	sort.Sort(nis)
	res := make([]int, P)
	for i := 0; i < P; i++ {
		res[i] = nis[i].second + 1
	}
	return res
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first || (ps[i].first == ps[j].first && ps[i].second < ps[j].second)
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
