package p1803

func countPairs(nums []int, low int, high int) int {
	root := new(Trie)
	var res int
	for _, num := range nums {
		a := find(root, num, high)
		b := find(root, num, low-1)
		res += a - b
		addNum(root, num)
	}

	return res
}

type Trie struct {
	children [2]*Trie
	cnt      int
}

const H = 20

func addNum(node *Trie, num int) {
	for i := H - 1; i >= 0; i-- {
		node.cnt++
		x := (num >> i) & 1
		if node.children[x] == nil {
			node.children[x] = new(Trie)
		}
		node = node.children[x]
	}
	node.cnt++
}

func find(root *Trie, num int, limit int) int {
	// find pairs some tmp xor num <= limit
	node := root
	var res int
	for i := H - 1; i >= 0 && node != nil; i-- {
		x := (num >> i) & 1

		y := (limit >> i) & 1

		if y == 0 {
			if node.children[x] == nil {
				node = nil
				break
			}
			node = node.children[x]
			continue
		}
		// y == 1
		if node.children[x] != nil {
			res += node.children[x].cnt
		}
		if node.children[1-x] == nil {
			node = nil
			break
		}
		node = node.children[1-x]
	}
	if node != nil {
		res += node.cnt
	}
	return res
}
