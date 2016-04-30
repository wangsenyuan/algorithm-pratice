package main

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
)

func play(rows [][]int, n int) []int {
	counts := make(map[int]int)

	for _, row := range rows {
		for _, x := range row {
			counts[x]++
		}
	}

	missing := make([]int, 0, n)

	for x, cnt := range counts {
		if cnt%2 == 1 {
			missing = append(missing, x)
		}
	}

	sort.Ints(missing)

	return missing
}

func concat(ints []int, sep string) string {
	var buffer bytes.Buffer
	_sep := ""
	for _, x := range ints {
		buffer.WriteString(_sep)
		buffer.WriteString(strconv.Itoa(x))
		_sep = sep
	}
	return buffer.String()
}

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 1; i <= t; i++ {
		var n int
		fmt.Scanf("%d", &n)
		rows := make([][]int, 0, 2*n-1)

		for j := 0; j < 2*n-1; j++ {
			row := make([]int, 0, n)
			for k := 0; k < n; k++ {
				var x int
				fmt.Scanf("%d", &x)
				row = append(row, x)
			}
			rows = append(rows, row)
		}

		missing := play(rows, n)
		fmt.Printf("Case #%d: %s\n", i, concat(missing, " "))
	}
}
