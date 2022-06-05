package p2296

import "bytes"

type Node struct {
	prev *Node
	next *Node
	v    byte
}

type TextEditor struct {
	head *Node
	end  *Node
	cur  *Node
}

func Constructor() TextEditor {
	head := new(Node)
	end := new(Node)
	head.next = end
	end.prev = head
	cur := end
	return TextEditor{head, end, cur}
}

func connect(a *Node, b *Node) {
	a.next = b
	b.prev = a
}

func (this *TextEditor) AddText(text string) {
	prev := this.cur.prev
	// insert text after prev
	for i := 0; i < len(text); i++ {
		x := text[i]
		node := new(Node)
		node.v = x
		connect(prev, node)
		connect(node, this.cur)
		prev = node
	}
}

func (this *TextEditor) DeleteText(k int) int {
	var cnt int
	node := this.cur.prev
	for k > 0 && node != this.head {
		prev := node.prev
		connect(prev, this.cur)
		node = prev
		cnt++
		k--
	}
	return cnt
}

func (this *TextEditor) CursorLeft(k int) string {
	for k > 0 && this.cur.prev != this.head {
		k--
		this.cur = this.cur.prev
	}
	return this.leftContext()
}

func (this *TextEditor) leftContext() string {
	node := this.cur.prev
	var buf bytes.Buffer
	for i := 0; i < 10 && node != this.head; i++ {
		buf.WriteByte(node.v)
		node = node.prev
	}
	arr := buf.Bytes()
	reverse(arr)
	return string(arr)
}

func reverse(arr []byte) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func (this *TextEditor) CursorRight(k int) string {
	for k > 0 && this.cur != this.end {
		this.cur = this.cur.next
		k--
	}
	return this.leftContext()
}

/**
 * Your TextEditor object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddText(text);
 * param_2 := obj.DeleteText(k);
 * param_3 := obj.CursorLeft(k);
 * param_4 := obj.CursorRight(k);
 */
