https://codeforces.com/problemset/problem/178/F2


### ideas
1. 从n个里面选择k个字符串，使得f(.,.)的和最大
2. 计算贡献吗？
3. 通过前缀树，可以知道各个前缀有多少个字符串；
4. 最长的前缀是500
5. 假设选择了m个前缀，a[1], .... a[m]
6. 每一个选择了个数b[1], .... b[m]
7. 那么有 b[1] + ... + b[m] = k
8. 使得 sum a[1] * (b[1] - 1) * b[1] 最大
9. dp[i][j] 表示前i个前缀，共选择了j个字符串
10. dp[i][j] = max(dp[i][j], dp[i-1][j - x] + len(i) * x * (x - 1) / 2)
11. 好像不大对， 就是比如 abc, abc, ab 
12. 这里前缀ab的个数是3，abc的个数是2，也就是说是有重叠的部分的
13. 也就是说，虽然选择了3个字符串，但是贡献确实 2 * 3 * (3 - 1) + 3 * 2 * 1
14. 排好序后，肯定是连续的k个字符串
15. 上面的例子，可以这样看，a有3个字符串贡献，所以它的贡献是 1 * 3 * 2/2， b有3个，它的贡献也是3，c有两个，所以它的贡献是1
16. 所以，增加一个新的字符串时，看它已有的，删除的时候，同理
17. 不一定是连续的
18. dp[i][j] = x就是前i个里面选择j个字符串后，能够得到的最大值