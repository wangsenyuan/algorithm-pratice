package p1441

func buildArray(target []int, n int) []string {
	var buf []string

	var j int
	var i int = 1

	for j < len(target) {
		x := target[j]

		if i == x {
			buf = append(buf, "Push")
			i++
			j++
			continue
		}
		buf = append(buf, "Push")
		buf = append(buf, "Pop")

		i++
	}

	return buf
}
