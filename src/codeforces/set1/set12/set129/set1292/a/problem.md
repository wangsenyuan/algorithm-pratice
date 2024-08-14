NEKO#ΦωΦ has just got a new maze game on her PC!

The game's main puzzle is a maze, in the forms of a 2×𝑛
 rectangle grid. NEKO's task is to lead a Nekomimi girl from cell (1,1)
 to the gate at (2,𝑛)
 and escape the maze. The girl can only move between cells sharing a common side.

However, at some moments during the game, some cells may change their state: either from normal ground to lava (which forbids movement into that cell), or vice versa (which makes that cell passable again). Initially all cells are of the ground type.

After hours of streaming, NEKO finally figured out there are only 𝑞
 such moments: the 𝑖
-th moment toggles the state of cell (𝑟𝑖,𝑐𝑖)
 (either from ground to lava or vice versa).

Knowing this, NEKO wonders, after each of the 𝑞
 moments, whether it is still possible to move from cell (1,1)
 to cell (2,𝑛)
 without going through any lava cells.

Although NEKO is a great streamer and gamer, she still can't get through quizzes and problems requiring large amount of Brain Power. Can you help her?


### ideas
1. 记录上下两行，有多少位置是block的（block是指这个位置造成了，上下被封锁了）
2. 好像记录一行也是可以的。因为修改下面一行的状态时，能够被波及的位置，只有3个