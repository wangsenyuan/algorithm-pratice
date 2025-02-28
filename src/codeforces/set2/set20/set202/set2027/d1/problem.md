You're given an array a of length n, and an array b of length m (bi>bi+1 for all 1≤i<m). Initially, the value of k is 1. Your aim is to make the array a empty by performing one of these two operations repeatedly:

Type 1 — If the value of k is less than m and the array a is not empty, you can increase the value of k by 1. This does not incur any cost.
Type 2 — You remove a non-empty prefix of array a, such that its sum does not exceed bk. This incurs a cost of m−k.
You need to minimize the total cost of the operations to make array a empty. If it's impossible to do this through any sequence of operations, output −1. Otherwise, output the minimum total cost of the operations.

### ideas
1. 虽然k越大，cost越小（m-k),但是b[k]越小，那么就越有可能无法删除一个前缀
2. 假设在位置i开始，如果可以在k=10时能删除掉
3. dp[i][k] = min dp[j][k] + m - k if sum[i...j] <= b[k]
4.          or dp[i][k+1]
5. n * m <= 1e5 那么就是ok的
6. 这里的问题就是要在i的后面，尽快的找出j满足条件
7. dp[j][k] 是最小的，且 sum[i...j-1] <= b[k] 要满足
8. 那么可以用个双端队列就可以了