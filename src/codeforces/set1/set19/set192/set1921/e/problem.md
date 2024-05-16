Alice and Bob are playing a game on a checkered board. The board has ℎ
 rows, numbered from top to bottom, and 𝑤
 columns, numbered from left to right. Both players have a chip each. Initially, Alice's chip is located at the cell with coordinates (𝑥𝑎,𝑦𝑎)
 (row 𝑥𝑎
, column 𝑦𝑎
), and Bob's chip is located at (𝑥𝑏,𝑦𝑏)
. It is guaranteed that the initial positions of the chips do not coincide. Players take turns making moves, with Alice starting.

On her turn, Alice can move her chip one cell down or one cell down-right or down-left (diagonally). Bob, on the other hand, moves his chip one cell up, up-right, or up-left. It is not allowed to make moves that go beyond the board boundaries.

More formally, if at the beginning of Alice's turn she is in the cell with coordinates (𝑥𝑎,𝑦𝑎)
, then she can move her chip to one of the cells (𝑥𝑎+1,𝑦𝑎)
, (𝑥𝑎+1,𝑦𝑎−1)
, or (𝑥𝑎+1,𝑦𝑎+1)
. Bob, on his turn, from the cell (𝑥𝑏,𝑦𝑏)
 can move to (𝑥𝑏−1,𝑦𝑏)
, (𝑥𝑏−1,𝑦𝑏−1)
, or (𝑥𝑏−1,𝑦𝑏+1)
. The new chip coordinates (𝑥′,𝑦′)
 must satisfy the conditions 1≤𝑥′≤ℎ
 and 1≤𝑦′≤𝑤
.
