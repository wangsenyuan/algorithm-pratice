Nash designed an interesting yet simple board game where a player is simply required to follow instructions written on the cell where the player currently stands.

This board game is played on the 𝑛×𝑛
 board. Rows and columns of this board are numbered from 1
 to 𝑛
. The cell on the intersection of the 𝑟
-th row and 𝑐
-th column is denoted by (𝑟,𝑐)
.

Some cells on the board are called blocked zones. On each cell of the board, there is written one of the following 5
 characters  — 𝑈
, 𝐷
, 𝐿
, 𝑅
 or 𝑋
  — instructions for the player. Suppose that the current cell is (𝑟,𝑐)
. If the character is 𝑅
, the player should move to the right cell (𝑟,𝑐+1)
, for 𝐿
 the player should move to the left cell (𝑟,𝑐−1)
, for 𝑈
 the player should move to the top cell (𝑟−1,𝑐)
, for 𝐷
 the player should move to the bottom cell (𝑟+1,𝑐)
. Finally, if the character in the cell is 𝑋
, then this cell is the blocked zone. The player should remain in this cell (the game for him isn't very interesting from now on).

It is guaranteed that the characters are written in a way that the player will never have to step outside of the board, no matter at which cell he starts.

As a player starts from a cell, he moves according to the character in the current cell. The player keeps moving until he lands in a blocked zone. It is also possible that the player will keep moving infinitely long.

For every of the 𝑛2
 cells of the board Alice, your friend, wants to know, how will the game go, if the player starts in this cell. For each starting cell of the board, she writes down the cell that the player stops at, or that the player never stops at all. She gives you the information she has written: for each cell (𝑟,𝑐)
 she wrote:

a pair (𝑥
,𝑦
), meaning if a player had started at (𝑟,𝑐)
, he would end up at cell (𝑥
,𝑦
).
or a pair (−1
,−1
), meaning if a player had started at (𝑟,𝑐)
, he would keep moving infinitely long and would never enter the blocked zone.
It might be possible that Alice is trying to fool you and there's no possible grid that satisfies all the constraints Alice gave you. For the given information Alice provided you, you are required to decipher a possible board, or to determine that such a board doesn't exist. If there exist several different boards that satisfy the provided information, you can find any of them.


### ideas
1. 如果已知原始的grid，是很容易计算出来a的，从所有的x开始，反向bfs即可
2. 那么所有的不是(-1, -1)的都是x所在的地方
3. 假设(r, c)的目的地是(x, y)
4. 那么必然存在一条路径从(r, c) -> ... -> (x, y)
5. 所有这类的组成一个联通块（如果不连通， 则没有答案）
6. 然后(-1, -1)的必须存在一个圈，才行（2个点也是可以的）
7. got了