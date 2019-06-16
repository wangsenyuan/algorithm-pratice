package main

import (
	"math"
	"fmt"
)

func findRestaurant(list1 []string, list2 []string) []string {

	find := func(short, long []string) []string {
		mi := make(map[string]int)
		for i, s := range short {
			mi[s] = i
		}

		var sum = math.MaxInt32
		var res []string
		for i, s := range long {
			if j, found := mi[s]; found {
				if i+j < sum {
					sum = i + j
					res = []string{s}
				} else if i+j == sum {
					res = append(res, s)
				}
			}
		}
		return res
	}

	if len(list1) < len(list2) {
		return find(list1, list2)
	}
	return find(list2, list1)
}

func main() {
	list1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	list2 := []string{"KFC", "Shogun", "Burger King"}

	fmt.Println(findRestaurant(list1, list2))
}
