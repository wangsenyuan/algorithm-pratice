package main

import "fmt"

func main() {
	/*words := []string{
		"abcd",
		"bnrt",
		"crmy",
		"dtye",
	}*/
	/*words := []string{
		"ball",
		"area",
		"read",
		"lady",
	}*/

	words := []string{
		"abcd",
		"bnrt",
		"crm",
		"dt",
	}
	fmt.Println(validWordSquare(words))

	words = []string{
		"ball",
		"asee",
		"let",
		"lep",
	}
	fmt.Println(validWordSquare(words))
	words = []string{
		"abc",
		"b",
		"c",
	}
	fmt.Println(validWordSquare(words))
}

func validWordSquare(words []string) bool {
	if len(words) == 0 {
		return true
	}

	var cols [][]byte
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			if j == len(cols) {
				cols = append(cols, make([]byte, 0, len(words)))
			}
			cols[j] = append(cols[j], words[i][j])
		}
	}
	if len(words) != len(cols) {
		return false
	}
	for i := 0; i < len(words); i++ {
		if words[i] != string(cols[i]) {
			return false
		}
	}

	return true
}
