You are given a string 𝑠
 of length 𝑛
 consisting of lowercase Latin letters. You may apply some operations to this string: in one operation you can delete some contiguous substring of this string, if all letters in the substring you delete are equal. For example, after deleting substring bbbb from string abbbbaccdd we get the string aaccdd.

Calculate the minimum number of operations to delete the whole string 𝑠
.

### ideas
1. dp[l][r] 表示删除完l...r所需的最小操作数
2. dp[l][r] = 如果dp[l1][r1] + 1 [l..l1), (r1...r] 是相同的
3. 不大对
4. 先将l...r清理掉，再处理外围的，不一定是最优的方案
5. 比如在例子 abaca中，先处理b,c, 再处理aaa 是最优的方案
6. 但是有一点, abaca = min(aaca, abaa, abca, ...) + 1的最优解
7. dp[i][j] 表示将这个区域清理掉的最优解
8. dp[i][j] = dp[i][k] + dp[k+1][j] 两边分别处理
9. 或者是，将i。。。j中间保留一种字符（比如a，其他的在区间内分隔后更小区间的dp的和）
10. 假设最后删除a是更优的（那么久没有必要在内部删除a，否则的话，它肯定会造成额外的操作）
11. 也就是说，如果a是最后删除的，那么a肯定是最后一起删除（不能在之前删除）
12. 比如例子abaca中，如果a是最后删除的，那么中间的a，不应该被提前删除掉
13. dp[l][r][x] 表示在区间[l...r]中删除掉所有除x外的最小操作数
14. dp[l][r][x] = 在这个区间内，按照x划分出，更小的区间，（且在这些更小的区间中间，肯定不会有x）
15. 所以就不用flag去表示哪些被删除了
16. 