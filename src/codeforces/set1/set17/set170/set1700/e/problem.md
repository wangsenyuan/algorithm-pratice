Little pirate Serega robbed a ship with puzzles of different kinds. Among all kinds, he liked only one, the hardest.

A puzzle is a table of 𝑛
 rows and 𝑚
 columns, whose cells contain each number from 1
 to 𝑛⋅𝑚
 exactly once.

To solve a puzzle, you have to find a sequence of cells in the table, such that any two consecutive cells are adjacent by the side in the table. The sequence can have arbitrary length and should visit each cell one or more times. For a cell containing the number 𝑖
, denote the position of the first occurrence of this cell in the sequence as 𝑡𝑖
. The sequence solves the puzzle, if 𝑡1<𝑡2<⋯<𝑡𝑛𝑚
. In other words, the cell with number 𝑥
 should be first visited before the cell with number 𝑥+1
 for each 𝑥
.

Let's call a puzzle solvable, if there exists at least one suitable sequence.

In one move Serega can choose two arbitrary cells in the table (not necessarily adjacent by the side) and swap their numbers. He would like to know the minimum number of moves to make his puzzle solvable, but he is too impatient. Thus, please tell if the minimum number of moves is 0
, 1
, or at least 2
. In the case, where 1
 move is required, please also find the number of suitable cell pairs to swap.

