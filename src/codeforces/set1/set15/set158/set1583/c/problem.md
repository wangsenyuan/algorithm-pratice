Consider a grid in which some cells are empty and some cells are filled. Call a cell in this grid exitable if, starting
at that cell, you can exit the grid by moving up and left through only empty cells. This includes the cell itself, so
all filled in cells are not exitable. Note that you can exit the grid from any leftmost empty cell (cell in the first
column) by going left, and from any topmost empty cell (cell in the first row) by going up.

Let's call a grid determinable if, given only which cells are exitable, we can exactly determine which cells are filled
in and which aren't.

You are given a grid 𝑎
of dimensions 𝑛×𝑚
, i. e. a grid with 𝑛
rows and 𝑚
columns. You need to answer 𝑞
queries (1≤𝑞≤2⋅105
). Each query gives two integers 𝑥1,𝑥2
(1≤𝑥1≤𝑥2≤𝑚
) and asks whether the subgrid of 𝑎
consisting of the columns 𝑥1,𝑥1+1,…,𝑥2−1,𝑥2
is determinable.

### ideas

1. 先不考虑子序列的问题， 如果存在一个区域，它不可离开，且这个区域的宽带/高度 >= 3， 那么就无法确定
2. 如果exitable + * 的面积 = 这块区域的面积，就是确定的
3. 如果某个空位（i， j)在整个区域就可以逃离，那么它一直都是可逃离的
4. 只有那些不可逃离的区域，删除掉左边后，它就有可能逃离掉，
5. 所以记录每个空区域能够到达的最左边的位置，当这个位置被剔除后，看是否能解救那些区域
6. 是不是我理解题目不大对～
7. 如果告诉一个区域那所有可逃离的区域，那么这些区域必然是空的，剩下的区域, 只能确定边缘是被堵住了
8. 但是内部就不确定了，所以有一个区域如果都是 X，其实是判断不出来的
9. 如果知道区域那所有不可确定的部分，然后不断的变得确定后，似乎就可以了

### solution

First notice that in a determinable grid, for any cell, it can't be that both the cell above it and the cell to its left
are filled. If that were the case, then the cell wouldn't be exitable regardless of whether it was filled or not, and so
we couldn't determine whether it was filled.

Now notice that in any grid with the above property, namely that from each cell you can move either up or to the left
into an empty cell (or both), every empty cell must be exitable — just keep moving either up or to the left, whichever
is possible, until you exit the grid.

It follows that for any grid satisfying that property, given only which cells are exitable, starting from the outermost
cells you will be able to determine that the nonexitable cells are filled, which implies that the next cells satisfy the
property, which further implies that the nonexitable ones there are filled, and so on. This allows you to determine the
entire grid (since the exitable cells are obviously empty).

Therefore, a grid being determinable is equivalent to all of its cells having an empty cell immediately above and/or to
the left of it. You can check this for arbitrary subgrids by precomputing two dimensional prefix sums of the cells that
violate this property, then checking whether the sum for a given subgrid is 0
. This solution is 𝑂(𝑛𝑚+𝑞)
.

The actual problem only asked for subgrids that contained every row, which allows for a somewhat simpler implementation.
