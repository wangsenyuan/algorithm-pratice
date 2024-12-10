You are given 𝑛
 arrays that can have different sizes. You also have a table with 𝑤
 columns and 𝑛
 rows. The 𝑖
-th array is placed horizontally in the 𝑖
-th row. You can slide each array within its row as long as it occupies several consecutive cells and lies completely inside the table.

You need to find the maximum sum of the integers in the 𝑗
-th column for each 𝑗
 from 1
 to 𝑤
 independently.


 ### ideas
 1. 给定column是可以在O(n)时间内处理出来吗？
 2. 这里有两个限制，一个是这一行不能到table的外面
 3. 似乎是可以的