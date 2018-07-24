package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
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

	f, err := os.Open("./A-large-practice.in")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	tc := readNum(scanner)

	for x := 1; x <= tc; x++ {
		N := readNum(scanner)
		mixtures := make([][]string, N)
		for i := 0; i < N; i++ {
			scanner.Scan()
			S := scanner.Text()
			mixtures[i] = strings.Split(S, " ")
		}
		res := solve(N, mixtures)
		fmt.Printf("Case #%d: %d\n", x, res)
	}
}

func solve(N int, mixtures [][]string) int {
	outs := make(map[string]map[string]bool)

	for _, mixture := range mixtures {
		u := mixture[0]
		outs[u] = make(map[string]bool)
		for i := 2; i < len(mixture); i++ {
			outs[u][mixture[i]] = true
		}
	}

	var dfs func(u string) int

	dfs = func(u string) int {
		if isIngredient(u) {
			return 0
		}
		// only need first max 2
		cnt := make([]int, len(outs[u]))
		var i int
		for v := range outs[u] {
			tmp := dfs(v)
			cnt[i] = tmp
			i++
		}

		sort.Ints(cnt)
		if cnt[i-1] == 0 {
			// all ingedients
			return 1
		}
		// at least cnt[i-1] bowls
		have := cnt[i-1] - 1
		ans := cnt[i-1]
		for j := i - 2; j >= 0 && cnt[j] > 0; j-- {
			if cnt[j] > have {
				ans++
				have = cnt[j]
			}
			have--
		}
		if have > 0 {
			return ans
		}
		return ans + 1
	}
	return dfs(mixtures[0][0])
}

func isIngredient(u string) bool {
	return u[0] >= 'a' && u[0] <= 'z'
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
