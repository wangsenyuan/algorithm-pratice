package mbp

// FindPerfectMaximumMatch is used to find the perfect maximum match number
// let's say there are n applicants and m jobs, each applicant may intersted in
// sereval jobs; matrix reprent that relationship; and this method find the
// max number of requirements can be satisified
func FindPerfectMaximumMatch(matrix [][]int, n int, m int) int {
	//job assigned to whom
	assignedTo := make([]int, m)

	for i := 0; i < m; i++ {
		assignedTo[i] = -1
	}

	var bpm func(v int) bool

	//job visited in the path
	seen := make([]bool, m)

	bpm = func(v int) bool {
		for i := 0; i < m; i++ {
			// try every interested job
			if matrix[v][i] == 1 && !seen[i] {
				seen[i] = true
				// this job is not assiged to other yet or can allocate another job for the previous applicant
				if assignedTo[i] < 0 || bpm(assignedTo[i]) {
					assignedTo[i] = v
					return true
				}
			}
		}

		return false
	}

	var res int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			seen[j] = false
		}
		if bpm(i) {
			res++
		}
	}

	return res
}
