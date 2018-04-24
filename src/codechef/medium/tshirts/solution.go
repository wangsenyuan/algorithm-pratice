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

func readNums(scanner *bufio.Scanner) []int {
	res := make([]int, 0, 10)
	x := -1
	scanner.Scan()
	bs := scanner.Bytes()
	for x < len(bs) {
		var num int
		x = readInt(bs, x+1, &num)
		res = append(res, num)
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	t := readNum(scanner)

	for t > 0 {
		n := readNum(scanner)
		people := make([][]int, n)
		for i := 0; i < n; i++ {
			people[i] = readNums(scanner)
		}

		fmt.Println(solve(n, people))
		t--
	}
}

const MOD = 1000000007

func solve(n int, people [][]int) int {
	ts := make(map[int]map[int]bool)

	for i := 0; i < n; i++ {
		for _, s := range people[i] {
			if _, found := ts[s]; !found {
				ts[s] = make(map[int]bool)
			}
			ts[s][i] = true
		}
	}

	N := 1 << uint(n)
	dp := make([]int, N)
	dp[N-1] = 1

	for _, ps := range ts {
		for i := 0; i < N; i++ {
			for x := range ps {
				if i&(1<<uint(x)) > 0 {
					dp[i^(1<<uint(x))] = (dp[i^(1<<uint(x))] + dp[i]) % MOD
				}
			}
		}
	}
	return dp[0]
}

func solve1(n int, people [][]int) int {
	ts := make(map[int]map[int]bool)

	for i := 0; i < n; i++ {
		for _, s := range people[i] {
			if _, found := ts[s]; !found {
				ts[s] = make(map[int]bool)
			}
			ts[s][i] = true
		}
	}

	N := 1 << uint(n)
	dp := make([][]int, N, 102)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, 102)
		for j := 0; j < 102; j++ {
			dp[i][j] = -1
		}
	}

	var rec func(mask int, id int) int

	rec = func(mask int, id int) int {
		if id == 101 {
			if mask == N-1 {
				dp[mask][id] = 1
			} else {
				dp[mask][id] = 0
			}
			return dp[mask][id]
		}

		if dp[mask][id] >= 0 {
			return dp[mask][id]
		}
		ans := rec(mask, id+1)

		for p := range ts[id] {
			if mask&(1<<uint(p)) == 0 {
				ans = (ans + rec(mask|(1<<uint(p)), id+1)) % MOD
			}
		}
		dp[mask][id] = ans
		return ans
	}

	return rec(0, 1)
}
