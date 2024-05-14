Kilani is playing a game with his friends. This game can be represented as a grid of size 𝑛×𝑚
, where each cell is either empty or blocked, and every player has one or more castles in some cells (there are no two castles in one cell).

The game is played in rounds. In each round players expand turn by turn: firstly, the first player expands, then the second player expands and so on. The expansion happens as follows: for each castle the player owns now, he tries to expand into the empty cells nearby. The player 𝑖
 can expand from a cell with his castle to the empty cell if it's possible to reach it in at most 𝑠𝑖
 (where 𝑠𝑖
 is player's expansion speed) moves to the left, up, right or down without going through blocked cells or cells occupied by some other player's castle. The player examines the set of cells he can expand to and builds a castle in each of them at once. The turned is passed to the next player after that.

The game ends when no player can make a move. You are given the game field and speed of the expansion for each player. Kilani wants to know for each player how many cells he will control (have a castle their) after the game ends.

[problem link](https://codeforces.com/problemset/problem/1105/D)

### ideas
1. 最多有9个 player
2. 考虑一个player的处理过程
3. 从它一开始的位置开始，每次从它已经占据的位置，向四个方向最多移动s[i]距离，遇到block，界外，或者其他player的区域
4. 考虑一个格子，如果它离player i的距离（经过几次后）是x, 而离player j的距离是
5. 如果j到到它的距离 >= i，那么它只能被i占据
6. 所以，如果对于每个位置，能计算出它哪个player占据掉，
7. 进而，也可以计算出它们相邻的位置
8. 这里理解错了，是距离现在的位置 <= s[i] 的地方