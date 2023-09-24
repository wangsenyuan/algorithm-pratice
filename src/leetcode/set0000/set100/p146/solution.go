package p146

type Node struct {
	prev  *Node
	next  *Node
	key   int
	value int
}

func NewNode(key int, value int) *Node {
	node := new(Node)
	node.key = key
	node.value = value
	return node
}

type LRUCache struct {
	front    *Node
	tail     *Node
	cache    map[int]*Node
	capacity int
}

func Constructor(capacity int) LRUCache {
	this := new(LRUCache)
	this.front = NewNode(-1, -1)
	this.tail = NewNode(-2, -2)
	this.front.next = this.tail
	this.tail.prev = this.front
	this.cache = make(map[int]*Node)
	this.capacity = capacity
	return *this
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		// need to move it to front
		this.moveNodeToFront(node)
		return node.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		this.moveNodeToFront(node)
		node.value = value
		return
	}
	node := NewNode(key, value)
	node.value = value
	this.cache[key] = node
	a := this.front.next
	this.front.next = node
	node.prev = this.front
	node.next = a
	a.prev = node

	if len(this.cache) > this.capacity {
		this.removeLast()
	}
}

func (this *LRUCache) removeLast() {
	a := this.tail.prev
	delete(this.cache, a.key)
	b := a.prev
	b.next = this.tail
	this.tail.prev = b
}

func (this *LRUCache) moveNodeToFront(node *Node) {
	a := node.prev
	b := node.next

	if a == this.front {
		// already in front
		return
	}
	// remove node out
	b.prev = a
	a.next = b

	c := this.front.next
	node.next = c
	c.prev = node

	node.prev = this.front
	this.front.next = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
