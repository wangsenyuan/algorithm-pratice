After a long day, Alice and Bob decided to play a little game. The game board consists of 𝑛
cells in a straight line, numbered from 1
to 𝑛
, where each cell contains a number 𝑎𝑖
between 1
and 𝑛
. Furthermore, no two cells contain the same number.

A token is placed in one of the cells. They take alternating turns of moving the token around the board, with Alice
moving first. The current player can move from cell 𝑖
to cell 𝑗
only if the following two conditions are satisfied:

the number in the new cell 𝑗
must be strictly larger than the number in the old cell 𝑖
(i.e. 𝑎𝑗>𝑎𝑖
), and
the distance that the token travels during this turn must be a multiple of the number in the old cell (i.e.
|𝑖−𝑗|mod𝑎𝑖=0
).
Whoever is unable to make a move, loses. For each possible starting position, determine who wins if they both play
optimally. It can be shown that the game is always finite, i.e. there always is a winning strategy for one of the
players.

Input
The first line contains a single integer 𝑛
(1≤𝑛≤105
) — the number of numbers.

The second line contains 𝑛
integers 𝑎1,𝑎2,…,𝑎𝑛
(1≤𝑎𝑖≤𝑛
). Furthermore, there are no pair of indices 𝑖≠𝑗
such that 𝑎𝑖=𝑎𝑗
.

Output
Print 𝑠
— a string of 𝑛
characters, where the 𝑖
-th character represents the outcome of the game if the token is initially placed in the cell 𝑖
. If Alice wins, then 𝑠𝑖
has to be equal to "A"; otherwise, 𝑠𝑖
has to be equal to "B".

### ideas

1. 按照数字a[i]降序处理, a[i] = n肯定是Bob胜利（因为alice无处可去）
2. 然后依次处理，不过要考虑a[i]的倍数所在的位置