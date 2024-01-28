package p3022

func countKeyChanges(s string) int {
	var res int

	for i := 1; i < len(s); i++ {
		a := get(s[i-1])
		b := get(s[i])
		if a != b {
			res++
		}
	}
	return res
}

func get(x byte) byte {
	if x >= 'A' && x <= 'Z' {
		return byte(x - 'A' + 'a')
	}
	return x
}
