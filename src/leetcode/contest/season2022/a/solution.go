package a

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	check := func(r int) bool {
		i, j := 0, len(arr)-1
		for i < j {
			if i == r {
				i++
				continue
			}
			if j == r {
				j--
				continue
			}
			if arr[i] != arr[j] {
				return false
			}
			i++
			j--
		}
		return true
	}

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			// need to remove either i or j
			if check(i) || check(j) {
				return true
			}
			return false
		}
	}

	return true
}
