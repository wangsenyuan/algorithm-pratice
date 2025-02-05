The sequence of integers 𝑎1,𝑎2,…,𝑎𝑘
 is called a good array if 𝑎1=𝑘−1
 and 𝑎1>0
. For example, the sequences [3,−1,44,0],[1,−99]
 are good arrays, and the sequences [3,7,8],[2,5,4,1],[0]
 — are not.

A sequence of integers is called good if it can be divided into a positive number of good arrays. Each good array should be a subsegment of sequence and each element of the sequence should belong to exactly one array. For example, the sequences [2,−3,0,1,4]
, [1,2,3,−3,−9,4]
 are good, and the sequences [2,−3,0,1]
, [1,2,3,−3−9,4,1]
 — are not.

For a given sequence of numbers, count the number of its subsequences that are good sequences, and print the number of such subsequences modulo 998244353.


### ideas
1. 不大对
2. 2, 1, 1
3. 假设a[i]被选中了，那么后面再跟a[i]个数后，可以和任意一个a[j]连接起来（前提是这个a[j]也能组成合法的组合）
4. dp[i]表示以i开始的有效的good序列的计数
5. dp[i] = sum nCr(j - i - 1, a[i]) * dp[j]，j - i - 1 >= a[i]
6. 但是这个式子貌似无法化简
7. r = a[i], nCr(r, r) * dp[j] + nCr(r + 1, r) * dp[j+1] + ...
8. nCr(n, r) = n! / (r! * (n-r)!)