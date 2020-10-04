package p1598

func minOperations(logs []string) int {
	var cur int

	for _, log := range logs {
		if log == "../" {
			cur--
			if cur < 0 {
				cur = 0
			}
		} else if log == "./" {
			// no change
		} else {
			cur++
		}
	}

	return cur
}
