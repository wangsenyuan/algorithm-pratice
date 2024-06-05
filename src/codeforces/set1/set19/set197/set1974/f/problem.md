Alice and Bob were playing a game again. They have a grid of size 𝑎×𝑏
 (1≤𝑎,𝑏≤109
), on which there are 𝑛
 chips, with at most one chip in each cell. The cell at the intersection of the 𝑥
-th row and the 𝑦
-th column has coordinates (𝑥,𝑦)
.

Alice made the first move, and the players took turns. On each move, a player could cut several (but not all) rows or columns from the beginning or end of the remaining grid and earn a point for each chip that was on the cut part of the grid. Each move can be described by the character 'U', 'D', 'L', or 'R' and an integer 𝑘
:

If the character is 'U', then the first 𝑘
 remaining rows will be cut;
If the character is 'D', then the last 𝑘
 remaining rows will be cut;
If the character is 'L', then the first 𝑘
 remaining columns will be cut;
If the character is 'R', then the last 𝑘
 remaining columns will be cut.
Based on the initial state of the grid and the players' moves, determine the number of points earned by Alice and Bob, respectively.

Input
The first line contains a single integer 𝑡
 (1≤𝑡≤104
) — the number of test cases.

The first line of each test case contains four integers 𝑎
, 𝑏
, 𝑛
, and 𝑚
 (2≤𝑎,𝑏≤109
, 1≤𝑛,𝑚≤2⋅105
) — the dimensions of the grid, the number of chips, and the number of moves.

Each of the next 𝑛
 lines contain two integers 𝑥𝑖
 and 𝑦𝑖
 (1≤𝑥𝑖≤𝑎
, 1≤𝑦𝑖≤𝑏
) — the coordinates of the chips. All pairs of coordinates are distinct.

Each of the next 𝑚
 lines contain a character 𝑐𝑗
 and an integer 𝑘𝑗
 — the description of the 𝑗
-th move. It is guaranteed that 𝑘
 is less than the number of rows/columns in the current grid. In other words, a player cannot cut the entire remaining grid on their move.

It is guaranteed that the sum of the values of 𝑛
 across all test cases in the test does not exceed 2⋅105
. It is guaranteed that the sum of the values of 𝑚
 across all test cases in the test does not exceed 2⋅105
.

### ideas
1. 通过模拟是否可行？
2. 好像是可以的，问题就是每步cut的时候，快速的找到哪些clip需要被移除掉
3. 在4个方向分别维护一个优先队列，比如在cut 上部时，使用top的队列