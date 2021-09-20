package p2010

func finalValueAfterOperations(operations []string) int {
	var x int
	for _, op := range operations {
		if op[1] == '+' {
			x++
		} else {
			x--
		}
	}
	return x
}
