package p690

/**
 * Definition for Employee.
 */
type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	mem := make(map[int]*Employee)
	for _, emp := range employees {
		mem[emp.Id] = emp
	}

	var dfs func(cur *Employee) int

	dfs = func(cur *Employee) int {
		if cur == nil {
			return 0
		}
		res := cur.Importance
		for _, sub := range cur.Subordinates {
			res += dfs(mem[sub])
		}

		return res
	}

	return dfs(mem[id])
}
