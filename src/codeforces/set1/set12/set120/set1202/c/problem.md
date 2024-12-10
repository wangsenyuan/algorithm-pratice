You have a string 𝑠
 — a sequence of commands for your toy robot. The robot is placed in some cell of a rectangular grid. He can perform four commands:

'W' — move one cell up;
'S' — move one cell down;
'A' — move one cell left;
'D' — move one cell right.
Let 𝐺𝑟𝑖𝑑(𝑠)
 be the grid of minimum possible area such that there is a position in the grid where you can place the robot in such a way that it will not fall from the grid while running the sequence of commands 𝑠
. For example, if 𝑠=DSAWWAW
 then 𝐺𝑟𝑖𝑑(𝑠)
 is the 4×3
 grid:

you can place the robot in the cell (3,2)
;
the robot performs the command 'D' and moves to (3,3)
;
the robot performs the command 'S' and moves to (4,3)
;
the robot performs the command 'A' and moves to (4,2)
;
the robot performs the command 'W' and moves to (3,2)
;
the robot performs the command 'W' and moves to (2,2)
;
the robot performs the command 'A' and moves to (2,1)
;
the robot performs the command 'W' and moves to (1,1)
.

You have 4
 extra letters: one 'W', one 'A', one 'S', one 'D'. You'd like to insert at most one of these letters in any position of sequence 𝑠
 to minimize the area of 𝐺𝑟𝑖𝑑(𝑠)
.

What is the minimum area of 𝐺𝑟𝑖𝑑(𝑠)
 you can achieve?


 ### ideas
 1. 一个s所圈定的面积 = 它的宽 * 它的高
 2. 它的宽= 左右的净移动量？还是最大连续移动量？
 3. 貌似是后者。比如连续10次A，那么宽度不能少于10，然后再5次S，净移动量是5，但是宽度是11（10 + 1）
 4. 所以，在任何一个位置，需要知道连续最多的前缀和后缀，然后通过添加不同的字符去计算