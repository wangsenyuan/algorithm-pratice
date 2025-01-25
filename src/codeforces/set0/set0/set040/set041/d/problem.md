On some square in the lowest row of a chessboard a stands a pawn. It has only two variants of moving: upwards and leftwards or upwards and rightwards. The pawn can choose from which square of the lowest row it can start its journey. On each square lay from 0 to 9 peas. The pawn wants to reach the uppermost row having collected as many peas as possible. As there it will have to divide the peas between itself and its k brothers, the number of peas must be divisible by k + 1. Find the maximal number of peas it will be able to collect and which moves it should make to do it.

The pawn cannot throw peas away or leave the board. When a pawn appears in some square of the board (including the first and last square of the way), it necessarily takes all the peas.

### ideas
1. dp[i][j][x] = 到达位置i,j，且 sum % (k+1) = x时的最大的sum