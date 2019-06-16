package main

import "fmt"

func main() {
	result := generate(4)

	for i := range result {
		fmt.Printf("%v\n", result[i])
	}
}

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	if numRows == 0 {
		return result
	}

	result[0] = []int{1}

	for i := 1; i < numRows; i++ {
		lastRow := result[i-1]
		thisRow := make([]int, len(lastRow)+1)
		for j := range thisRow {
			if j == 0 {
				thisRow[j] = lastRow[j]
			} else if j == len(lastRow) {
				thisRow[j] = lastRow[j-1]
			} else {
				thisRow[j] = lastRow[j-1] + lastRow[j]
			}
		}
		result[i] = thisRow
	}

	return result
}
