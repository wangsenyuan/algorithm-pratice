package p339

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 */
type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool {
	return false
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() int {
	return 0
}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []*NestedInteger {
	return nil
}

func depthSum(nestedList []*NestedInteger) int {

	var sum func(nestedList []*NestedInteger, depth int) int
	sum = func(nestedList []*NestedInteger, depth int) int {
		res := 0
		for _, each := range nestedList {
			if each.IsInteger() {
				res += each.GetInteger() * depth
			} else {
				res += sum(each.GetList(), depth+1)
			}
		}
		return res
	}

	return sum(nestedList, 1)
}
