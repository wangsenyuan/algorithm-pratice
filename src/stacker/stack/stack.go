// stack
package stack

import (
	"errors"
)

type Stack []interface{}

func (stack Stack) Len() int {
	return len(stack)
}

func (stack Stack) Cap() int {
	return cap(stack)
}

func (stack Stack) IsEmpty() bool {
	return stack.Len() == 0
}

func (stack *Stack) Push(x interface{}) {
	*stack = append(*stack, x)
}

func (stack Stack) Top() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("Can't Top an empty stack")
	}
	return stack[len(stack)-1], nil
}

func (stack *Stack) Pop() (interface{}, error) {
	theStack := *stack
	if len(theStack) == 0 {
		return nil, errors.New("Can't Pop an empty stack")
	}

	x := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return x, nil
}
