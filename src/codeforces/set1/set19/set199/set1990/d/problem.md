You are given an array 𝑎
 of size 𝑛
.

There is an 𝑛×𝑛
 grid. In the 𝑖
-th row, the first 𝑎𝑖
 cells are black and the other cells are white. In other words, note (𝑖,𝑗)
 as the cell in the 𝑖
-th row and 𝑗
-th column, cells (𝑖,1),(𝑖,2),…,(𝑖,𝑎𝑖)
 are black, and cells (𝑖,𝑎𝑖+1),…,(𝑖,𝑛)
 are white.

You can do the following operations any number of times in any order:

Dye a 2×2
 subgrid white;
Dye a whole row white. Note you can not dye a whole column white.
Find the minimum number of operations to dye all cells white.

### ideas
1. 如果存在一个区域可以在操作1中，完成就使用操作1，否则就使用操作2
2. 感觉太简单了吧？
3. 