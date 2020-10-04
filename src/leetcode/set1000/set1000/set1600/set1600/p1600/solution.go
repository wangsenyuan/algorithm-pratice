package p1600

type ThroneInheritance struct {
	parent   map[string]string
	children map[string][]string
	dead     map[string]bool
	king     string
}

func Constructor(kingName string) ThroneInheritance {
	parent := make(map[string]string)
	children := make(map[string][]string)
	dead := make(map[string]bool)
	return ThroneInheritance{parent, children, dead, kingName}
}

func (this *ThroneInheritance) Birth(parentName string, childName string) {
	this.parent[childName] = parentName
	if len(this.children[parentName]) == 0 {
		this.children[parentName] = make([]string, 0, 10)
	}
	this.children[parentName] = append(this.children[parentName], childName)
}

func (this *ThroneInheritance) Death(name string) {
	this.dead[name] = true
}

func (this *ThroneInheritance) GetInheritanceOrder() []string {
	cur := make([]string, 0, len(this.parent))
	if !this.dead[this.king] {
		cur = append(cur, this.king)
	}
	pos := make(map[string]int)

	var dfs func(x string)

	dfs = func(x string) {

		if pos[x] == len(this.children[x]) {
			if x == this.king {
				return
			}
			dfs(this.parent[x])
		} else {
			pos[x]++
			y := this.children[x][pos[x]-1]
			if !this.dead[y] {
				cur = append(cur, y)
			}
			dfs(y)
		}
	}

	dfs(this.king)

	return cur
}

/**
 * Your ThroneInheritance object will be instantiated and called as such:
 * obj := Constructor(kingName);
 * obj.Birth(parentName,childName);
 * obj.Death(name);
 * param_3 := obj.GetInheritanceOrder();
 */
