package p364

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger interface {
	IsInteger() bool
	GetInteger() int
	GetList() []*NestedInteger
}

func depthSumInverse(nestedList []*NestedInteger) int {
	return sum(nestedList, maxDepth(nestedList))
}

func sum(nestedList []*NestedInteger, depth int) int {
	s := 0

	for _, nlp := range nestedList {
		nl := *nlp
		if nl.IsInteger() {
			s += nl.GetInteger() * depth
			continue
		}

		s += sum(nl.GetList(), depth-1)
	}

	return s
}

func maxDepth(nestedList []*NestedInteger) int {
	depth := 1

	for _, nlp := range nestedList {
		nl := *nlp
		if nl.IsInteger() {
			continue
		}
		depth = max(depth, 1+maxDepth(nl.GetList()))
	}

	return depth
}
func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
