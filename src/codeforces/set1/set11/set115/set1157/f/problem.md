There are 𝑛
 people in a row. The height of the 𝑖
-th person is 𝑎𝑖
. You can choose any subset of these people and try to arrange them into a balanced circle.

A balanced circle is such an order of people that the difference between heights of any adjacent people is no more than 1
. For example, let heights of chosen people be [𝑎𝑖1,𝑎𝑖2,…,𝑎𝑖𝑘]
, where 𝑘
 is the number of people you choose. Then the condition |𝑎𝑖𝑗−𝑎𝑖𝑗+1|≤1
 should be satisfied for all 𝑗
 from 1
 to 𝑘−1
 and the condition |𝑎𝑖1−𝑎𝑖𝑘|≤1
 should be also satisfied. |𝑥|
 means the absolute value of 𝑥
. It is obvious that the circle consisting of one person is balanced.

Your task is to choose the maximum number of people and construct a balanced circle consisting of all chosen people. It is obvious that the circle consisting of one person is balanced so the answer always exists.

### ideas
1. 一个观察是，假设最终的长度是k, 比如5，那么最大值和最小值的差，不会超过3
2. 1,2,3,3,2
3. 将a升序排列。dp[i]表示i前面可以得到的最长序列
4. dp[i] = dp[i-1] or 如果 a[i]是最大值时的序列，这里需要两个信息，就是它左边的序列，右边的序列，且左右两边序列的端点的差，不超过1
5. 这个序列里面，除了两端的，必须至少有两个，否则肯定没法闭环
6. 那就好办了