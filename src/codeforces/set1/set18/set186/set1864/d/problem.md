There is a matrix of size 𝑛×𝑛
which consists of 0s and 1s. The rows are numbered from 1
to 𝑛
from top to bottom, the columns are numbered from 1
to 𝑛
from left to right. The cell at the intersection of the 𝑥
-th row and the 𝑦
-th column is denoted as (𝑥,𝑦)
.

AquaMoon wants to turn all elements of the matrix to 0s. In one step she can perform the following operation:

Select an arbitrary cell, let it be (𝑖,𝑗)
, then invert the element in (𝑖,𝑗)
and also invert all elements in cells (𝑥,𝑦)
for 𝑥>𝑖
and 𝑥−𝑖≥|𝑦−𝑗|
. To invert a value means to change it to the opposite: 0 changes to 1, 1 changes to 0.
Help AquaMoon determine the minimum number of steps she need to perform to turn all elements of the matrix to 0s. We can
show that an answer always exists.

Input
Each test contains multiple test cases. The first line contains the number of test cases 𝑡
(1≤𝑡≤105
). The description of the test cases follows.

The first line of each test case contains an integer 𝑛
(2≤𝑛≤3000
).

The 𝑖
-th of the following 𝑛
lines contains a binary string only of characters 0 and 1, of length 𝑛
.

It is guaranteed that the sum of 𝑛2
over all test cases does not exceed 9000000
.

### thoughts

1. 操作(i, j) 相当于操作整个以(i, j)为起点的倒三角的区域
2. 在操作(i, j)的时候，需要将这整个区域都变成和(i,j)相同的值
3. 假设dp[i,j,0]表示将(i, j, 0)三角区域变成0的次数
4. dp[i, j, 0] = dp[i + 1, j, 1] + dp[i + 1, j + 1, 1] + 1 + dp[i+2, j+2, 1]
5. 这个表示不大对
6. 如果(i, j) = 0
6. dp[i, j, 0] = dp[i + 1, j, 0] + dp[i + 1, j + 1, 0] + 1 (因为(i + 2， j + 2)被翻转了2次，要翻转回来)
7. if (i, j) is 1, dp[i, j, 0] = 1 + dp[i+1, j, 1] + dp[i+1, j + 1, 1] + 1 (在i, j处要再翻转一次)
8. 