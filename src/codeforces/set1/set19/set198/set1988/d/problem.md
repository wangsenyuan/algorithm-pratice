You, the monster killer, want to kill a group of monsters. The monsters are on a tree with 𝑛
 vertices. On vertex with number 𝑖
 (1≤𝑖≤𝑛
), there is a monster with 𝑎𝑖
 attack points. You want to battle with monsters for 10100
 rounds.

In each round, the following happens in order:

All living monsters attack you. Your health decreases by the sum of attack points of all living monsters.
You select some (possibly all or none) monsters and kill them. After being killed, the monster will not be able to do any attacks in the future.
There is a restriction: in one round, you cannot kill two monsters that are directly connected by an edge.

If you choose what monsters to attack optimally, what is the smallest health decrement you can have after all rounds?

Input
Each test contains multiple test cases. The first line contains the number of test cases 𝑡
 (1≤𝑡≤104
). Description of the test cases follows.

The first line of each test case contains an integer 𝑛
 (1≤𝑛≤3⋅105
).

The second line of each test case contains 𝑛
 integers 𝑎1,…,𝑎𝑛
 (1≤𝑎𝑖≤1012
).

The following 𝑛−1
 lines each contain two integers 𝑥,𝑦
 (1≤𝑥,𝑦≤𝑛
), denoting an edge on the tree connecting vertex 𝑥
 and 𝑦
.

It is guaranteed that the sum of 𝑛
 over all test cases does not exceed 3⋅105
.

Output
For each test case, print one integer: the minimum possible health decrement.

### ideas
1. 首先，有可能在两轮内干掉所有的monster，奇数高度的一次，偶数高度一次
2. 但是如果这么简单的话，肯定不会到2000分。所以应该有更优的方案
3. 同一层的，肯定是可以一起处理的，所以可以搞成一个数
4. 那么现在就变成一条直线了，不能同时删除两个相临节点时的最小值是多少？
5. 假设现在在x轮次后完成了，那么就必须1, 2, 3, 4这些节点分配一个轮次
6. 它们的贡献 = a[1] * x[1] + a[2] * x[2] + a[3] * x[3] + ... + a[i] * x[i] + ...
7. 且不能有两个连续的x[i], x[i+1]有相同的值
8. 但怎么处理呢？
9. 比如  3, 1, 2 显然 3/2配置1是合理的
10. 但是如果是 1, 3, 1, 那么3配置1是合理的
11. 感觉最多是3，不可能超过3
12. 之所以出现3是因为，会有 5, 1, 1, 4 这种组合出来
13. dp[i][j] 表示i配置为j时的最优解
14. dp[i][1] = min(dp[i-1][2], dp[i-1][3]) + a[i]
15. dp[i][2] = min(dp[i-1][1], dp[i-1][3]) + 2 * a[i]
16. dp[i][3] = min(dp[i-1][1], dp[i-1][2]) + 3 * a[i]