You're given an array a of length n, and an array b of length m (bi>bi+1 for all 1≤i<m). Initially, the value of k is 1. Your aim is to make the array a empty by performing one of these two operations repeatedly:

Type 1 — If the value of k is less than m and the array a is not empty, you can increase the value of k by 1. This does not incur any cost.
Type 2 — You remove a non-empty prefix of array a, such that its sum does not exceed bk. This incurs a cost of m−k.
You need to minimize the total cost of the operations to make array a empty. If it's impossible to do this through any sequence of operations, output −1. Otherwise, output the minimum total cost of the operations, and the number of sequences of operations which yield this minimum cost modulo 109+7.

Two sequences of operations are considered different if you choose a different type of operation at any step, or the size of the removed prefix is different at any step.


#### ideas
1. 如何计算取到最小值的方法数呢？
2. dp[i][j] 将后缀i清理掉，且在i处的k不小于j时的state
3. state = {min ops, ways to get the min_ops}
4. dp[i][j] = dp[r][j] + m - j要对这些相同最小操作数的后缀进行求和 
5. 是不是不需要保留比 dp[i][j]更差的结果的？
6. 假设从i开始能够到达j， dp[i] = dp[j] + m - k
7. 如果上面的j对i成立, 即 sum[i] - sum[j+1] <= k
8. 那么 sum[i+1] - sum[j+1] <= k 肯定也成立
9. 是不是双指针的情况下，再加一个指针就可以了？
10. 