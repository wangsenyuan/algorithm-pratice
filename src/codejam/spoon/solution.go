package main

import (
	"fmt"
)

func main() {
	/*for i := 0; i < 26; i++ {
		fmt.Println(toLower(byte('A' + i)))
	}*/

	//fmt.Println(toLower('A'))
	//fmt.Println(toLower('a'))

	var t int
	fmt.Scanf("%d\n", &t)
	for t > 0 {
		t--
		var r, c int
		fmt.Scanf("%d %d\n", &r, &c)
		matrix := make([][]byte, r)
		for i := 0; i < r; i++ {
			var s string
			fmt.Scanf("%s\n", &s)
			matrix[i] = []byte(s)
		}

		res := findSpoon(matrix, r, c)

		if res {
			fmt.Println("There is a spoon!")
		} else {
			fmt.Println("There is indeed no spoon!")
		}
	}
}

const spoon = "spoon"

func findSpoon(matrix [][]byte, r, c int) bool {
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if !equalsIgnoreCase(matrix[i][j], 's') {
				continue
			}
			a := 1
			for k := j + 1; k < c && a < len(spoon); k++ {
				if !equalsIgnoreCase(matrix[i][k], spoon[a]) {
					break
				}
				a++
			}
			if a == len(spoon) {
				return true
			}

			a = 1
			for k := i + 1; k < r && a < len(spoon); k++ {
				if !equalsIgnoreCase(matrix[k][j], spoon[a]) {
					break
				}
				a++
			}
			if a == len(spoon) {
				return true
			}
		}
	}

	return false
}

func equalsIgnoreCase(a, b byte) bool {
	return toLower(a) == toLower(b)
}

func toLower(a byte) byte {
	//strings.ToLower("")
	if a >= 'A' && a <= 'Z' {
		return a - 'A' + 'a'
	}
	return a
}
