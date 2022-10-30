package p2454

func twoEditWords(queries []string, dictionary []string) []string {

	check := func(word string) bool {
		for _, dic := range dictionary {
			var cnt int
			for i := 0; i < len(word) && cnt <= 2; i++ {
				if dic[i] != word[i] {
					cnt++
				}
			}
			if cnt <= 2 {
				return true
			}
		}
		return false
	}

	var res []string

	for _, word := range queries {
		if check(word) {
			res = append(res, word)
		}
	}

	return res
}
