Moamen was drawing a grid of 𝑛
 rows and 109
 columns containing only digits 0
 and 1
. Ezzat noticed what Moamen was drawing and became interested in the minimum number of rows one needs to remove to make the grid beautiful.

A grid is beautiful if and only if for every two consecutive rows there is at least one column containing 1
 in these two rows.

Ezzat will give you the number of rows 𝑛
, and 𝑚
 segments of the grid that contain digits 1
. Every segment is represented with three integers 𝑖
, 𝑙
, and 𝑟
, where 𝑖
 represents the row number, and 𝑙
 and 𝑟
 represent the first and the last column of the segment in that row.

For example, if 𝑛=3
, 𝑚=6
, and the segments are (1,1,1)
, (1,7,8)
, (2,7,7)
, (2,15,15)
, (3,1,1)
, (3,15,15)
, then the grid is:


Your task is to tell Ezzat the minimum number of rows that should be removed to make the grid beautiful.


### ideas
1. beautiful = 任意相邻的两行，至少有一列，两行同时为1
2. 只知道哪一行被删除了不够，还必须知道哪一行被保留下来了
3. dp[i] = 在保留第i行的情况下,留下的最大行数量
4. dp[i] = dp[j] + 1 如果 (i, j)之间有重叠的区域
5. 这样子就比较好搞了