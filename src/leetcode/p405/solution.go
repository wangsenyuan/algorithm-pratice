package main

import "fmt"

var hex = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}

func toHex(num int) string {

	i := uint(0)

	var res string
	for i < 32 {
		a := (num >> i) & 15
		x := hex[a]
		res = x + res
		i += 4
	}
	j := 0
	for j < len(res) {
		if res[j] != '0' {
			break
		}
		j++
	}
	if j == len(res) {
		return "0"
	}
	return res[j:]
}

func main() {
	fmt.Println(toHex(1000))
	fmt.Println(toHex(26))
	fmt.Println(toHex(-1))
	fmt.Println(toHex(1))
}
