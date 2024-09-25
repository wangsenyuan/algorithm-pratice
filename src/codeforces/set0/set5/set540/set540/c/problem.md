You play a computer game. Your character stands on some level of a multilevel ice cave. In order to move on forward, you need to descend one level lower and the only way to do this is to fall through the ice.

The level of the cave where you are is a rectangular square grid of n rows and m columns. Each cell consists either from intact or from cracked ice. From each cell you can move to cells that are side-adjacent with yours (due to some limitations of the game engine you cannot make jumps on the same place, i.e. jump from a cell to itself). If you move to the cell with cracked ice, then your character falls down through it and if you move to the cell with intact ice, then the ice on this cell becomes cracked.

Let's number the rows with integers from 1 to n from top to bottom and the columns with integers from 1 to m from left to right. Let's denote a cell on the intersection of the r-th row and the c-th column as (r, c).

You are staying in the cell (r1, c1) and this cell is cracked because you've just fallen here from a higher level. You need to fall down through the cell (r2, c2) since the exit to the next level is there. Can you do this?

### ideas
1. 如果(r, c)是cracked，走到上面后，是掉下水去，而不是掉到下一层