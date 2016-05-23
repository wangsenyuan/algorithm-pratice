package preorder

func fastIsPreOrder(nums []int) bool {
  stack := make([]int, 0, len(nums))

  root := math.Int_MIN

  for _, x := range nums {
    if x < root {
      return false
    }

    for len(stack) > 0 && stack[len(stack) - 1] < x {
      stack = stack[0:len(stack) - 1]
    }

    stack = append(stack, x)
  }

  return true
}

func slowIsPreOrder(nums []int) bool {
    if len(nums) <= 1 {
      return true
    }

    x := nums[0]

    for i := 1; i < len(nums); i++ {
      if nums[i] > x {
        for j := i + 1; j < len(nums); j++ {
          if nums[j] < x {
            return false
          }
        }

        return slowIsPreOrder(nums[1:i]) && slowIsPreOrder(nums[i:])
      }
    }
    return slowIsPreOrder(nums[1:])
}
