Gildong has bought a famous painting software cfpaint. The working screen of cfpaint is square-shaped consisting of 𝑛
 rows and 𝑛
 columns of square cells. The rows are numbered from 1
 to 𝑛
, from top to bottom, and the columns are numbered from 1
 to 𝑛
, from left to right. The position of a cell at row 𝑟
 and column 𝑐
 is represented as (𝑟,𝑐)
. There are only two colors for the cells in cfpaint — black and white.

There is a tool named eraser in cfpaint. The eraser has an integer size 𝑘
 (1≤𝑘≤𝑛
). To use the eraser, Gildong needs to click on a cell (𝑖,𝑗)
 where 1≤𝑖,𝑗≤𝑛−𝑘+1
. When a cell (𝑖,𝑗)
 is clicked, all of the cells (𝑖′,𝑗′)
 where 𝑖≤𝑖′≤𝑖+𝑘−1
 and 𝑗≤𝑗′≤𝑗+𝑘−1
 become white. In other words, a square with side equal to 𝑘
 cells and top left corner at (𝑖,𝑗)
 is colored white.

A white line is a row or a column without any black cells.

Gildong has worked with cfpaint for some time, so some of the cells (possibly zero or all) are currently black. He wants to know the maximum number of white lines after using the eraser exactly once. Help Gildong find the answer to his question.

#### ideas
1. 从某个位置（i, j)删除一个size = k的正方形区域后，这个区域外面的部分没有影响，可以先计算出来
2. 这个区域里面的部分，从行的角度看，要看所有的部分，是否在这些列中间
3. 