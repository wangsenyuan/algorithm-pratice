package p5554

func canFormArray(arr []int, pieces [][]int) bool {
	mem := make(map[int]int)

	for i := 0; i < len(pieces); i++ {
		mem[pieces[i][0]] = i
	}

	for i := 0; i < len(arr); {
		cur := arr[i]
		if ii, found := mem[cur]; !found {
			return false
		} else {
			for j := 0; j < len(pieces[ii]); j++ {
				if arr[i] != pieces[ii][j] {
					return false
				}
				i++
			}
		}
	}
	return true
}
