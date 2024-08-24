A binary matrix is called good if every even length square sub-matrix has an odd number of ones.

Given a binary matrix 𝑎
 consisting of 𝑛
 rows and 𝑚
 columns, determine the minimum number of cells you need to change to make it good, or report that there is no way to make it good at all.

All the terms above have their usual meanings — refer to the Notes section for their formal definitions.

Input
The first line of input contains two integers 𝑛
 and 𝑚
 (1≤𝑛≤𝑚≤106
 and 𝑛⋅𝑚≤106
)  — the number of rows and columns in 𝑎
, respectively.

The following 𝑛
 lines each contain 𝑚
 characters, each of which is one of 0 and 1. If the 𝑗
-th character on the 𝑖
-th line is 1, then 𝑎𝑖,𝑗=1
. Similarly, if the 𝑗
-th character on the 𝑖
-th line is 0, then 𝑎𝑖,𝑗=0
.

### ideas
1. 所有偶数长度的正方形，里面包含奇数个1
2. 这个条件里面的不变性是什么？
3. 假设所有长度为2的正方形，都满足条件，就是在每个区域内有奇数的1
4. 那么长度为4的正方形呢？
5. 肯定不成立的，因为长度为4的正方形，可以分成4个长度为2的正方形；而奇数 * 4 肯定是偶数
6. 所以，如果 m >= 4 且 n >= 4 => -1
7. 所以有 m < 4 or n < 4, 不妨设 n < 4 (即最多只有3行)
8. n = 1 => 0
9. n = 2 => 只需要依次检查dp[0/1/2/3]的状态即可
10. n = 3 => dp[0~7]的状态即可