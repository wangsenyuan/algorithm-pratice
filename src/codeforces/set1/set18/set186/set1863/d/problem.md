There is an 𝑛×𝑚
board divided into cells. There are also some dominoes on this board. Each domino covers two adjacent cells (that is,
two cells that share a side), and no two dominoes overlap.

Piet thinks that this board is too boring and it needs to be painted. He will paint the cells of the dominoes black and
white. He calls the painting beautiful if all of the following conditions hold:

1. for each domino, one of its cells is painted white and the other is painted black;
2. for each row, the number of black cells in this row equals the number of white cells in this row;
3. for each column, the number of black cells in this column equals the number of white cells in this column.

Note that the cells that are not covered by dominoes are not painted at all, they are counted as neither black nor
white.

Help Piet produce a beautiful painting or tell that it is impossible.

### thoughts

1. 这里面的不变量是什么？
2. 如果有一条垂直的domino,那么上下两行的奇偶性差1？
3. 除去这个domino, 假设上面有x个black cell, x + 1个白色
4. 那么下面一行，就有y个black cell, 且y-1个white cell
5. 让上面black，下面white，就可以保证上下两行满足条件
6. 此时，在上下间连一条线即可
7. 水平同样处理，然后进行dfs 黑/白着色
8. 上面的想法是错误的

### solution

Let's consider the requirement on the rows. Clearly, all horizontal dominoes (since each of them has 1
black and 1
white cell) do not influence the black-white balance for the rows. Thus, we are only interested in vertical dominoes.

Consider the first row and the vertical dominoes that intersect this row. Their number has to be even, otherwise, the
first row has an odd number of cells covered by dominoes and the solution is clearly impossible. But if there is an even
number of such dominoes, we have to paint half of them black-white, and half of them white-black. What's more, it
doesn't actually matter the exact order we paint them in, because vertical dominoes do not affect the columns' balance
and we will not influence the balance of the second row anyway. So we can freely paint them however we like. The same
logic applies for rows 2,…,𝑛
.

Now we turn to horizontal dominoes. In the first row, there once again has to be an even number of dominoes which
intersect this column. And we can paint half them black-white, and half of them white-black, and it doesn't matter which
exact way we choose. Do the same for columns 2,…,𝑚
.