package p1790

func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	var i int
	for i < len(s1) && s1[i] == s2[i] {
		i++
	}
	// s1[i] != s2[i]
	first := i
	i++
	for i < len(s1) && s1[i] == s2[i] {
		i++
	}
	if i == len(s1) {
		return false
	}
	second := i
	if s1[first] != s2[second] {
		return false
	}
	i++
	for i < len(s1) && s1[i] == s2[i] {
		i++
	}
	return i == len(s1)
}
