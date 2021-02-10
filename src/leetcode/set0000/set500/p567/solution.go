package main

func checkInclusion(s1 string, s2 string) bool {
	m1 := make([]int, 26)

	for i := 0; i < len(s1); i++ {
		m1[int(s1[i]-'a')]++
	}

	m2 := make([]int, 26)

	j := 0
	for i := 0; i < len(s2); i++ {
		x := int(s2[i] - 'a')
		m2[x]++

		ok := true

		for j := 0; j < 26 && ok; j++ {
			ok = m1[j] == m2[j]
		}

		if ok {
			return true
		}

		for m2[x] > m1[x] && j <= i {
			m2[int(s2[j]-'a')]--
			j++
		}
	}
	return false
}
