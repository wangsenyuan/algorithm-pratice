You are given an array of integers 𝑎
 of length 𝑛
.

In one operation, you will perform the following two-step process:

Choose an index 𝑖
 such that 1≤𝑖<|𝑎|
 and 𝑎𝑖=𝑖
.
Remove 𝑎𝑖
 and 𝑎𝑖+1
 from the array and concatenate the remaining parts.
Find the maximum number of times that you can perform the operation above.

### ideas
1. a[i] = i, remove a[i] and a[i+1]
2. 第一个感觉是先删除后面的。但这里可能会存在一个情况，如果删除前面的某个，其实可以删除a[i+1]的，但是现在删除不了了？
3. 即便先删除后面的成立，也还是无法推进。因为总要去删除前面的某个i，它删除后，会造成后面原来不能删除的，变得可删除了。
4. 那这个状态就不稳定了。
5. 我们不去考虑顺序，如果要删除a[i] (+a[i+1])
6. 那么必须满足a[i] = i(它一开始可能不满足)，那么就必须在它的前面删除a[i] - i 个数才行（a[i] - i 必须是偶数)
7. dp[i][j] 表示在操作i时，至少有j次操作（删掉了2 * j个数时，是否ok）
8. dp[i][j] = dp[i-1][j]不考虑删除a[i]
9. dp[i][j] = dp[i-1][j-1] 成立，且 a[i] - i = 2 * j
10. 这里这个j似乎是确定的，没法迭代
11. dp[i][j] = 是否可以在i前面有j次删除
12. dp[i][0] = true (没有任何删除，肯定是成立的)
13. dp[i][j] = dp[i-1][j] (如果在i-1前面有j次删除) or dp[i-1][j-1] and a[i] - i = 2 * j
14. fp[i][j] 表示在i的前面进行j次删除后，在i开始的suf上能得到的最大操作次数
15. fp[n+1][?] = 0
16.  fp[i][j] = max(fp[i+1][j], fp[i+2][j+1] + 1 if dp[i][j])
17.  搞不定～～～
18.  再想想。。。每个数字只能往前移，下标只会变小
19.  所以，按照需要移动的距离（有些是负值，需要往后移的，可以忽略掉，还有需要移动奇数位的，也可以排除掉）
20.  然后按照移动距离排序，同样距离时，先处理右边的
21.  