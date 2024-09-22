Polycarp and Vasiliy love simple logical games. Today they play a game with infinite chessboard and one pawn for each player. Polycarp and Vasiliy move in turns, Polycarp starts. In each turn Polycarp can move his pawn from cell (x, y) to (x - 1, y) or (x, y - 1). Vasiliy can move his pawn from (x, y) to one of cells: (x - 1, y), (x - 1, y - 1) and (x, y - 1). Both players are also allowed to skip move.

There are some additional restrictions — a player is forbidden to move his pawn to a cell with negative x-coordinate or y-coordinate or to the cell containing opponent's pawn The winner is the first person to reach cell (0, 0).

You are given the starting coordinates of both pawns. Determine who will win if both of them play optimally well.


### ideas
1. A 可以往左，或者往下移动
2. B 可以往左，往下、或者往左下移动
3. 所以，A的移动的最快速度 = x + y
4. B的移动的最快速度 = max(x, y)
5. 这里不清楚的地方在于，A能否阻挡B？
6. B只要移动到A的左下方，那么A就无法阻挡
7. 如果B正好在A的上方，且现在轮到A运行，那么A只要向左就可以了（或者右方也是如此）
8. 假设如果dx = dy, 那么肯定是A赢
9. dx < dy, 那么对B来说，就要最快的到达A的前方位置去
10. 这时候，好像也是A赢呐，A可以等到B出现在它的上方的时候，开始运动
11. 如果dx = 0, 但是dy > 0, A也可以获胜
12. 