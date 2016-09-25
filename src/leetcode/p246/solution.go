package main

import "fmt"

var pairs map[string]string

func init() {
	pairs = make(map[string]string)
	pairs["8"] = "8"
	pairs["1"] = "1"
	pairs["0"] = "0"
	pairs["6"] = "9"
	pairs["9"] = "6"

}

func isStrobogrammatic(num string) bool {
	i, j := 0, len(num)
	for i < j {
		a := num[i : i+1]
		b := num[j-1 : j]
		if x, ok := pairs[a]; ok && x == b {
			i++
			j--
			continue
		}
		return false
	}

	return true
}

func main() {
	//fmt.Println(isStrobogrammatic("88"))
	//fmt.Println(isStrobogrammatic("818"))
	fmt.Println(isStrobogrammatic("818"))
	fmt.Println(isStrobogrammatic("828"))
	fmt.Println(isStrobogrammatic("69"))
	fmt.Println(isStrobogrammatic("96"))
	fmt.Println(isStrobogrammatic("60809"))
	fmt.Println(isStrobogrammatic("6"))
	fmt.Println(isStrobogrammatic("1"))

}
