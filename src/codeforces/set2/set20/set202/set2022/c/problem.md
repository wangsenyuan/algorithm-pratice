Álvaro and José are the only candidates running for the presidency of Tepito, a rectangular grid of 2
 rows and 𝑛
 columns, where each cell represents a house. It is guaranteed that 𝑛
 is a multiple of 3
.

Under the voting system of Tepito, the grid will be split into districts, which consist of any 3
 houses that are connected∗
. Each house will belong to exactly one district.

Each district will cast a single vote. The district will vote for Álvaro or José respectively if at least 2
 houses in that district select them. Therefore, a total of 2𝑛3
 votes will be cast.

As Álvaro is the current president, he knows exactly which candidate each house will select. If Álvaro divides the houses into districts optimally, determine the maximum number of votes he can get.

### ideas
1. 这里划分有好几种方式
2. 必须考虑前2列的状态，且必须保证前3列都已经分配好了
3. dp[i][state], 0 表示前面没有任何空间了, 这里有3种选择进行下去
4.  1表示第一行的前一列为空，且第二列被填好了，这时候，有两种选择，一种是横着排3个，或者一个 T
5.  2表示第一行的前两列为空，且第二列被填好了。这时候，只有一种选择
6.  3表示第二行
7.  4表示第二行的前两列为空
8.  答案 = dp[n][0]