package main

import (
	"fmt"
	"geeks/permutations/iterate"
	"geeks/permutations/recursive"

)

func main() {

	fmt.Println("iterate way ########")

	ps := iterate.Permutations("ABCD")

	for _, p := range ps {
		fmt.Println(p)
	}

	fmt.Println("recursive way ######")
	ps = recursive.Permutations("ABCD")


	for _, p := range ps {
		fmt.Println(p)
	}
}
