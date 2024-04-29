Max and Lucas are playing the game. Max goes first, then Lucas, then Max again and so on. Each player has a marble,
initially located at some vertex. Each player in his/her turn should move his/her marble along some edge (a player can
move the marble from vertex v to vertex u if there's an outgoing edge from v to u). If the player moves his/her marble
from vertex v to vertex u, the "character" of that round is the character written on the edge from v to u. There's one
additional rule; the ASCII code of character of round i should be greater than or equal to the ASCII code of character
of round i - 1 (for i>1). The rounds are numbered for both players together, i. e. Max goes in odd numbers, Lucas goes
in even numbers. The player that can't make a move loses the game. The marbles may be at the same vertex at the same
time.

Since the game could take a while and Lucas and Max have to focus on finding Dart, they don't have time to play. So they
asked you, if they both play optimally, who wins the game?

You have to determine the winner of the game for all initial positions of the marbles.

### ideas

1. 假设当前Alice， Bob所在的节点是(i, j), 且前一轮的字符是x, 且当前轮由谁move 0/1
2. dp[i][j][x][0/1] = dp[u][v][y][1/0] 如果能够move
3. 如果轮到alice move , 那么要存在(i -> u, y) 的一条路径
4. 如果轮到bob move 要么要存在(j -> v, y) 的一条路径
5. dp[a][b][c][d] = 1表示 Alice获胜， -1, 表示bob获胜
6. 