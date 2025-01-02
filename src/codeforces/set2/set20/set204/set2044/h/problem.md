Swing is opening a pancake factory! A good pancake factory must be good at flattening things, so Swing is going to test his new equipment on 2D matrices.

Swing is given an 𝑛×𝑛
 matrix 𝑀
 containing positive integers. He has 𝑞
 queries to ask you.

For each query, he gives you four integers 𝑥1
, 𝑦1
, 𝑥2
, 𝑦2
 and asks you to flatten the submatrix bounded by (𝑥1,𝑦1)
 and (𝑥2,𝑦2)
 into an array 𝐴
. Formally, 𝐴=[𝑀(𝑥1,𝑦1),𝑀(𝑥1,𝑦1+1),…,𝑀(𝑥1,𝑦2),𝑀(𝑥1+1,𝑦1),𝑀(𝑥1+1,𝑦1+1),…,𝑀(𝑥2,𝑦2)]
.

The following image depicts the flattening of a submatrix bounded by the red dotted lines. The orange arrows denote the direction that the elements of the submatrix are appended to the back of 𝐴
, and 𝐴
 is shown at the bottom of the image.


Afterwards, he asks you for the value of ∑|𝐴|𝑖=1𝐴𝑖⋅𝑖
 (sum of 𝐴𝑖⋅𝑖
 over all 𝑖
).

Input
The first line contains an integer 𝑡
 (1≤𝑡≤103
) — the number of test cases.

The first line of each test contains two integers 𝑛
 and 𝑞
 (1≤𝑛≤2000,1≤𝑞≤106
) — the length of 𝑀
 and the number of queries.

The following 𝑛
 lines contain 𝑛
 integers each, the 𝑖
'th of which contains 𝑀(𝑖,1),𝑀(𝑖,2),…,𝑀(𝑖,𝑛)
 (1≤𝑀(𝑖,𝑗)≤106
).

The following 𝑞
 lines contain four integers 𝑥1
, 𝑦1
, 𝑥2
, and 𝑦2
 (1≤𝑥1≤𝑥2≤𝑛,1≤𝑦1≤𝑦2≤𝑛
) — the bounds of the query.

It is guaranteed that the sum of 𝑛
 over all test cases does not exceed 2000
 and the sum of 𝑞
 over all test cases does not exceed 106
.

Output
For each test case, output the results of the 𝑞
 queries on a new line.

 ### ideas
 1. 如果是一维的情况，要怎么处理呢？
 2. a[1] * 1 + a[2] * 2 + ... + a[i] * i + ... + a[k] * k
 3. a[1] + a[2] + ... + a[k] + a[2] + .. + a[k]
 4. sum[1...k] + sum[2...k] + .... sum[k...k]
 5. pref[k] - pref[0] + ... pref[k] - pref[1] + ... + pref[k] - pref[k-1]
 6. k * pref[k] - (pref[0] + pref[1] + ... pref[k-1])
 7. 到二维的情况， k很好计算. pref[k] 也很好计算（sub - matrix的和）
 8. 如果是在sub-matrix种的第r行
 9. a[i] * i = a[j] * (r * w + j) = r * w * (... 这一行的情况)
 10. 如果有个结构可以很快的算出每一行的情况，然后有个结构再算出每一列的情况，似乎就很好处理了
 11. dp[i][j][w] = 以(i, j)结尾，在该行长度为w的一个子区间的和（一维）
 12. dp[i][j][0] = a[0]
 13. dp[i][j][1] = dp[i][j - 1][0] + 2 * dp[i][j][0] 
 14. dp[i][j][w] = dp[i][j - (1 << (w-1))][w-1] + (1 << w) ❌
 15. dp[i][j][w] = dp[i][j-(1 << (w-1))][w-1] + dp[i][j][w-1] + (1 <<(w - 1)) * sum( 后半段)
 16. 知道这个，要怎么计算矩阵呢？
 17. a[i-1][j] => a[i][j]
 18. a[i-1][j] * x, a[i][j] * (x + w) (w是子矩阵的宽度)
 19. a[i-1][j] * x + a[i][j] * x + a[i][j] * w
 20. x * (a[i-1][j] + a[i][j])
 21. 从高到低的考虑 a[1] * x + a[2] * (x + w) + a[3] * (x + 2 * w) ... + a[i] * (x + (i-1) * w)
 22. (a[1] + a[2] + ... + a[i]) * x + a[2] * w + a[3] *  2w + ... + a[i] * (i-1) * w
 23. sum * x + w * (a[2] + a[3] * 2 + a.... + a[i] * (i-1))
 24. 后面部分貌似是一个一维的子问题
 25. 是不是太复杂了？
 26. 