On a chessboard with a width of 109
and a height of 109
, the rows are numbered from bottom to top from 1
to 109
, and the columns are numbered from left to right from 1
to 109
. Therefore, for each cell of the chessboard you can assign the coordinates (𝑥,𝑦)
, where 𝑥
is the column number and 𝑦
is the row number.

Every day there are fights between black and white pieces on this board. Today, the black ones won, but at what price?
Only the rook survived, and it was driven into the lower left corner — a cell with coordinates (1,1)
. But it is still happy, because the victory has been won and it's time to celebrate it! In order to do this, the rook
needs to go home, namely — on the upper side of the field (that is, in any cell that is in the row with number 109
).

Everything would have been fine, but the treacherous white figures put spells on some places of the field before the end
of the game. There are two types of spells:

Vertical. Each of these is defined by one number 𝑥
. Such spells create an infinite blocking line between the columns 𝑥
and 𝑥+1
.
Horizontal. Each of these is defined by three numbers 𝑥1
, 𝑥2
, 𝑦
. Such spells create a blocking segment that passes through the top side of the cells, which are in the row 𝑦
and in columns from 𝑥1
to 𝑥2
inclusive. The peculiarity of these spells is that it is impossible for a certain pair of such spells to have a common
point. Note that horizontal spells can have common points with vertical spells.

An example of a chessboard.
Let's recall that the rook is a chess piece that in one move can move to any point that is in the same row or column
with its initial position. In our task, the rook can move from the cell (𝑟0,𝑐0)
into the cell (𝑟1,𝑐1)
only under the condition that 𝑟1=𝑟0
or 𝑐1=𝑐0
and there is no blocking lines or blocking segments between these cells (For better understanding, look at the samples).

Fortunately, the rook can remove spells, but for this it has to put tremendous efforts, therefore, it wants to remove
the minimum possible number of spells in such way, that after this it can return home. Find this number!

Input
The first line contains two integers 𝑛
and 𝑚
(0≤𝑛,𝑚≤105
) — the number of vertical and horizontal spells.

Each of the following 𝑛
lines contains one integer 𝑥
(1≤𝑥<109
) — the description of the vertical spell. It will create a blocking line between the columns of 𝑥
and 𝑥+1
.

Each of the following 𝑚
lines contains three integers 𝑥1
, 𝑥2
and 𝑦
(1≤𝑥1≤𝑥2≤109
, 1≤𝑦<109
) — the numbers that describe the horizontal spell. It will create a blocking segment that passes through the top sides
of the cells that are in the row with the number 𝑦
, in columns from 𝑥1
to 𝑥2
inclusive.

It is guaranteed that all spells are different, as well as the fact that for each pair of horizontal spells it is true
that the segments that describe them do not have common points.

Output
In a single line print one integer — the minimum number of spells the rook needs to remove so it can get from the cell (
1,1)
to at least one cell in the row with the number 109

### ideas

1. vertical lines 将区域分成了完全不相关的m+1个区域
2. 且如果要移除vertical的区域，应该在最开始就移除
3. 从后往前的区域进行处理
4. 假设要在区域i中到达顶部，就需要找到这个区域[l...r]中最小的值（横向间隔的个数)
5. 数位压缩 + 差分数组 (或者 segment tree)
6. 😢，忽略了情况，是它不需要一次到达目的地，其实它可以到达某层以后，左右移动到某个位置继续
7. 所以只有那些从1开始，并在当前竖版（包括）后面的是必须要移除的