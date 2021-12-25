package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		n := readNum(scanner)
		scanner.Scan()
		s := scanner.Text()
		fmt.Println(solve(n, s))
	}
}

func solve(n int, s string) int64 {
	single := make([]int64, 26)
	pair := make([][][]Pair, 26)

	for i := 0; i < 26; i++ {
		pair[i] = make([][]Pair, 26)
		for j := 0; j < 26; j++ {
			pair[i][j] = make([]Pair, 0, 10)
		}
	}

	letters := make([]Pair, 0, n)

	for i, j := 1, 0; i <= n; i++ {
		if i == n || s[i] != s[j] {
			letter := Pair{int(s[j] - 'a'), i - j}
			letters = append(letters, letter)
			j = i
		}
	}

	for i := 0; i < len(letters); i++ {
		x := letters[i].first
		single[x] = int64(max(int(single[x]), letters[i].second))
		if i+1 < len(letters) {
			y := letters[i+1].first
			pair[x][y] = append(pair[x][y], Pair{letters[i].second, letters[i+1].second})
		}
	}

	calculate := func(i, j int) int64 {
		ps := pair[i][j]
		sort.Sort(Pairs(ps))

		vec := make([]Pair, 0, len(ps))

		for i := 0; i < len(ps); i++ {
			j := len(vec)
			for j > 0 && vec[j-1].second <= ps[i].second {
				j--
			}
			vec = append(vec[:j], ps[i])
			j++
		}
		var ans int64

		var last int

		for _, p := range vec {
			ans += int64(p.first-last) * int64(p.second)
			last = p.first
		}

		return ans
	}

	var ans int64

	for i := 0; i < 26; i++ {
		ans += single[i]
		for j := 0; j < 26; j++ {
			ans += calculate(i, j)
		}
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Pair struct {
	first  int
	second int
}

type Pairs []Pair

func (this Pairs) Len() int {
	return len(this)
}

func (this Pairs) Less(i, j int) bool {
	return this[i].first < this[j].first || (this[i].first == this[j].first && this[i].second < this[j].second)
}

func (this Pairs) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
