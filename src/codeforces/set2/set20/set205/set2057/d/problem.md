"T-Generation" has decided to purchase gifts for various needs; thus, they have 𝑛
 different sweaters numbered from 1
 to 𝑛
. The 𝑖
-th sweater has a size of 𝑎𝑖
. Now they need to send some subsegment of sweaters to an olympiad. It is necessary that the sweaters fit as many people as possible, but without having to take too many of them.

They need to choose two indices 𝑙
 and 𝑟
 (1≤𝑙≤𝑟≤𝑛
) to maximize the convenience equal to
max(𝑎𝑙,𝑎𝑙+1,…,𝑎𝑟)−min(𝑎𝑙,𝑎𝑙+1,…,𝑎𝑟)−(𝑟−𝑙),
that is, the range of sizes minus the number of sweaters.

Sometimes the sizes of the sweaters change; it is known that there have been 𝑞
 changes, in each change, the size of the 𝑝
-th sweater becomes 𝑥
.

Help the "T-Generation" team and determine the maximum convenience among all possible pairs (𝑙,𝑟)
 initially, as well as after each size change.

 ### ideas
 1. 先考虑没有变化时，如何求解
 2. 假设在区间l...r中间的最大值发生在a[i]处，最小值发生在a[j]处
 3. 这段的score = a[i] - a[j] - (r - l)
 4. 如果j != l, 且i != l, 那么此时增加l是一个更有的选择（不影响最大值，最小值，且size变小了）
 5. 所以，知道l变成了i或者j
 6. 同理，也可以得到r会变成i或者j
 7. 所以，可以推论出，在一个可能的最优score中, l...r正好是两个极值
 8. 中间的数，都必然处在a[l]...a[r]间
 9. 接下来怎么搞呢？
 10. 假设a[l]是最小值， a[r]是最大值
 11. score = a[r] - a[l] - (r - l) = (a[r] - r) - (a[l] - l)
 12. 也就是对于固定的r， 要找到最小的 a[l] - l 的值
 13. 单次是是没有问题的，每个节点维护它左边最小的l(a[l] - l 的最小值)
 14. 但是更新a[i]后，要怎么处理呢？
 15. 这里有两种情况，一种情况是，a[i]不是原来的最优值的一部分（那么这次有可能变成新的最优值）
 16. a[i]就是原来最优解的一部分。 这次要计算新的最优解
 17. 还是有些问题，比如在l...r，是一个能够得到最优解的情况(l是最小的 a[l] - l 的值)
 18. 但是当a[l]变化以后，那么对于r来说，它的最优解就变成了另外一个
 19. 所以还是得维护一个区间
 20. 不大对。 刚才那样子，潜在的复杂性，至少是n*n(a[1]是最小值，那么所有的r都会挂在这个节点上)
 21. 修改a[1], 会造成整个结构重新计算
 22. 修改a[i], 那么就是在右边找出最大的a[r] - r 的最大值