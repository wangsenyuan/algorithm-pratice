package p2940

func findMinimumOperations(s1 string, s2 string, s3 string) int {
	var i int
	for i < len(s1) && i < len(s2) && i < len(s3) {
		if s1[i] == s2[i] && s2[i] == s3[i] {
			i++
			continue
		}
		break
	}

	if i == 0 {
		return -1
	}
	return len(s1) - i + len(s2) - i + len(s3) - i
}
