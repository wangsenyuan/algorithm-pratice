package p082

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := new(ListNode)
	head.Val = arr[0]
	head.Next = NewList(arr[1:])
	return head
}

func (list ListNode) ToArray() []int {
	cur := &list
	var res []int
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	return res
}

func deleteDuplicates(head *ListNode) *ListNode {
	newHead := &ListNode{-1, head}
	cur := newHead
	for cur != nil {
		// cur should be kept
		tmp := cur.Next
		if tmp == nil {
			break
		}
		for tmp != nil {
			x := tmp.Val
			cnt := 1
			for tmp.Next != nil && tmp.Next.Val == x {
				cnt++
				tmp = tmp.Next
			}
			if cnt == 1 {
				cur = tmp
				break
			}
			cur.Next = tmp.Next
			tmp = tmp.Next
		}
	}
	return newHead.Next
}
