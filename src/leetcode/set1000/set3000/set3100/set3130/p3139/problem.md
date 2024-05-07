给你一个整数数组 nums 和两个整数 cost1 和 cost2 。你可以执行以下 任一 操作 任意 次：

从 nums 中选择下标 i 并且将 nums[i] 增加 1 ，开销为 cost1。
选择 nums 中两个 不同 下标 i 和 j ，并且将 nums[1] 和 nums[2] 都 增加 1 ，开销为 cost2 。
你的目标是使数组中所有元素都 相等 ，请你返回需要的 最小开销 之和。

由于答案可能会很大，请你将它对 109 + 7 取余 后返回。

### ideas

1. 首先可以确定的是，通过操作1，始终可以使得所有的数字相等
2. 那什么情况下使用操作2呢？如果2 * cost1 <= cost2, 那么只需要使用操作1
3. 只有到 2 * cost1 > cost2时，才应该使用操作2
4. 假设最后的相等的值等于w, sum = sum of（w - num[i]) for every i
5. 那么最小的cost = x * cost1 + y * cost2 (x次操作1，y次操作2)
6. x + 2 * y = sum
7. 然后这里有个关键的信息，就是w = max(num[i]) 或者 +1, 因为更多的数会更差
8. 