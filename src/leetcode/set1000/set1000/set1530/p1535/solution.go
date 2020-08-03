package p1535

func getWinner(arr []int, k int) int {
	n := len(arr)
	if k >= n || n == 2 {
		return findMax(arr)
	}

	// k < n
	nodes := make([]*Node, n)
	for i := 0; i < n; i++ {
		node := new(Node)
		node.val = arr[i]
		if i > 0 {
			nodes[i-1].next = node
			node.prev = nodes[i-1]
		}
		nodes[i] = node
	}

	tail := nodes[n-1]
	cur := nodes[0]

	for cur.win < k {
		next := cur.next
		if cur.val > next.val {
			cur.win++
			tmp := next.next
			cur.next = tmp
			tmp.prev = cur

			tail.next = next
			next.prev = tail
			next.next = nil
			next.win = 0
			tail = next
			continue
		}
		// cur.val < next.val
		next.win++
		tail.next = cur
		cur.prev = tail
		cur.next = nil
		cur.win = 0
		tail = cur
		cur = next
	}

	return cur.val
}

type Node struct {
	val  int
	win  int
	prev *Node
	next *Node
}

func findMax(arr []int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > res {
			res = arr[i]
		}
	}
	return res
}
