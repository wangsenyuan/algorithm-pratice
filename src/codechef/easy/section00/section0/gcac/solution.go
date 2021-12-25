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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var t int
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &t)

	for i := 0; i < t; i++ {
		var n, m int
		scanner.Scan()
		bs := scanner.Bytes()
		readInt(bs, readInt(bs, 0, &n)+1, &m)
		minSalary := make([]int, n)
		scanner.Scan()
		bs = scanner.Bytes()
		for j, x := 0, -1; j < n; j++ {
			x = readInt(bs, x+1, &minSalary[j])
		}
		offeredSalary, maxJobOffers := make([]int, m), make([]int, m)
		for j := 0; j < m; j++ {
			var a, b int
			scanner.Scan()
			bs = scanner.Bytes()
			readInt(bs, readInt(bs, 0, &a)+1, &b)
			offeredSalary[j] = a
			maxJobOffers[j] = b
		}
		qual := make([][]byte, n)
		for j := 0; j < n; j++ {
			scanner.Scan()
			qual[j] = scanner.Bytes()
		}
		employee, salary, employer := solve(n, minSalary, m, offeredSalary, maxJobOffers, qual)
		fmt.Printf("%d %d %d\n", employee, salary, employer)
	}
}

func solve(n int, minSalary []int, m int, offeredSalary []int, maxJobOffers []int, qual [][]byte) (int, int64, int) {
	var employeed int
	var totalSalary int64
	companies := make([]bool, m)

	for i := 0; i < n; i++ {
		selected := -1
		for j := 0; j < m; j++ {
			if qual[i][j] == '0' {
				continue
			}
			if minSalary[i] > offeredSalary[j] {
				continue
			}
			if maxJobOffers[j] == 0 {
				continue
			}
			if selected == -1 || offeredSalary[selected] < offeredSalary[j] {
				selected = j
			}
		}
		if selected >= 0 {
			companies[selected] = true
			maxJobOffers[selected]--
			employeed++
			totalSalary += int64(offeredSalary[selected])
		}
	}
	var totalEmployers int
	for i := 0; i < m; i++ {
		if companies[i] {
			totalEmployers++
		}
	}
	return employeed, totalSalary, m - totalEmployers
}
