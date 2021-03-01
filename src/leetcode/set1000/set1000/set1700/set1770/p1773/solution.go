package p1773

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	var res int

	for _, item := range items {
		if ruleKey == "type" {
			if item[0] == ruleValue {
				res++
			}
		} else if ruleKey == "color" {
			if item[1] == ruleValue {
				res++
			}
		} else {
			if item[2] == ruleValue {
				res++
			}
		}
	}

	return res
}
