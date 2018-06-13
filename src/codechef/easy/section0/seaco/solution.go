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
		n, m := readTwoNums(scanner)

		cmds := make([][]int, m)
		for i := 0; i < m; i++ {
			cmds[i] = readNNums(scanner, 3)
		}

		res := solve(n, m, cmds)

		fmt.Printf("%d", res[0])
		for i := 1; i < n; i++ {
			fmt.Printf(" %d", res[i])
		}

		fmt.Println()
		t--
	}
}

const MOD = 1000000007

func solve(n, m int, cmds [][]int) []int64 {
	dcnt := make([]int64, m)
	cnt := make([]int64, m)
	cnt[m-1] = 1
	for i := m - 1; i >= 0; i-- {
		if i < m-1 {
			cnt[i] = cnt[i+1] + dcnt[i]
			if cnt[i] >= MOD {
				cnt[i] -= MOD
			}
		}
		left, right := cmds[i][1]-1, cmds[i][2]-1
		if cmds[i][0] == 2 {
			dcnt[right] += cnt[i]
			if dcnt[right] >= MOD {
				dcnt[right] -= MOD
			}
			if left > 0 {
				dcnt[left-1] -= cnt[i]
				if dcnt[left-1] < 0 {
					dcnt[left-1] += MOD
				}
			}
		}
	}
	da := make([]int64, n)

	for i := 0; i < m; i++ {
		if cmds[i][0] == 1 {
			left, right := cmds[i][1]-1, cmds[i][2]-1
			da[right] += cnt[i]
			if da[right] >= MOD {
				da[right] -= MOD
			}
			if left > 0 {
				da[left-1] -= cnt[i]
				if da[left-1] < 0 {
					da[left-1] += MOD
				}
			}
		}
	}
	a := make([]int64, n)
	for i := n - 1; i >= 0; i-- {
		a[i] = da[i]
		if i < n-1 {
			a[i] = (a[i] + a[i+1]) % MOD
		}
	}
	return a
}
