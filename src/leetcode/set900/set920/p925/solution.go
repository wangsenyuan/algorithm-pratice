package p925

func isLongPressedName(name string, typed string) bool {
	if len(name) > len(typed) {
		return false
	}

	var i, j int

	for i < len(name) && j < len(typed) {
		if name[i] != typed[j] {
			return false
		}
		cnt1 := i
		for cnt1 < len(name) && name[cnt1] == name[i] {
			cnt1++
		}
		cnt1 -= i
		cnt2 := j
		for cnt2 < len(typed) && typed[cnt2] == typed[j] {
			cnt2++
		}
		cnt2 -= j
		if cnt1 > cnt2 {
			return false
		}
		i += cnt1
		j += cnt2
	}

	return i == len(name) && j == len(typed)
}
