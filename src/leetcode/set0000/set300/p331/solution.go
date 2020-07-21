package p331

import "strings"

func isValidSerialization(preorder string) bool {
	nodes := strings.Split(preorder, ",")

	nulls := 0

	for i := len(nodes) - 1; i >= 0; i-- {
		if nodes[i] == "#" {
			nulls++
		} else if nulls < 2 {
			return false
		} else {
			nulls--
		}
	}

	return nulls == 1
}
