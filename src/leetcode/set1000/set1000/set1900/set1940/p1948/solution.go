package p1948

const N = 501

const MOD1 = 1000000009
const MOD2 = 1000000007

const PP = 1000000013

var B1 [N]int64
var B2 [N]int64

func init() {
	B1[0] = 1
	B2[0] = 1
	for i := 1; i < N; i++ {
		B1[i] = B1[i-1] * PP % MOD1
		B2[i] = B2[i-1] * PP % MOD2
	}
}

func deleteDuplicateFolder(paths [][]string) [][]string {
	root := NewNode("/")
	for _, path := range paths {
		AddPath(root, path)
	}

	cnt := make(map[Pair]int)

	var dfs func(node *Node) (int, Pair)
	dfs = func(node *Node) (int, Pair) {
		var d int

		key := calculateKey(node.label)

		if len(node.children) > 0 {
			var x, y int64
			for _, child := range node.children {
				cd, tmp := dfs(child)
				if cd > d {
					d = cd
				}
				x = (x + tmp.first) % MOD1
				y = (y + tmp.second) % MOD2
			}
			node.key = Pair{x, y}
			cnt[node.key]++

			d++
			x += key.first * B1[d] % MOD1
			y += key.second * B2[d] % MOD2
			x %= MOD1
			y %= MOD2
			key = Pair{x, y}
		}

		return d, key
	}

	dfs(root)

	res := make([][]string, 0, len(paths))

	hold := make([]string, N)

	var dfs2 func(node *Node, d int, path []string)

	dfs2 = func(node *Node, d int, path []string) {
		if cnt[node.key] > 1 {
			// duplicate
			return
		}
		path[d] = node.label
		if d > 0 {
			res = append(res, copyPath(path[1:d+1]))
		}
		for _, child := range node.children {
			dfs2(child, d+1, path)
		}
	}

	dfs2(root, 0, hold)

	return res
}

func copyPath(arr []string) []string {
	res := make([]string, len(arr))
	copy(res, arr)
	return res
}

type Pair struct {
	first  int64
	second int64
}

type Node struct {
	children map[string]*Node
	label    string
	key      Pair
}

func AddPath(root *Node, path []string) {

	node := root

	for i := 0; i < len(path); i++ {
		cur := path[i]
		if node.children[cur] == nil {
			node.children[cur] = NewNode(cur)
		}
		node = node.children[cur]
	}
}

func NewNode(label string) *Node {
	node := new(Node)
	node.children = make(map[string]*Node)
	node.label = label
	node.key = Pair{0, 0}
	return node
}

func calculateKey(label string) Pair {
	var x, y int64

	for i := 0; i < len(label); i++ {
		z := int64(label[i]-'a') + 1
		x = (x*27 + z) % MOD1
		y = (y*27 + z) % MOD2
	}

	return Pair{x, y}
}
