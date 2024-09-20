Petya has a rectangular Board of size 𝑛×𝑚
. Initially, 𝑘
 chips are placed on the board, 𝑖
-th chip is located in the cell at the intersection of 𝑠𝑥𝑖
-th row and 𝑠𝑦𝑖
-th column.

In one action, Petya can move all the chips to the left, right, down or up by 1
 cell.

If the chip was in the (𝑥,𝑦)
 cell, then after the operation:

left, its coordinates will be (𝑥,𝑦−1)
;
right, its coordinates will be (𝑥,𝑦+1)
;
down, its coordinates will be (𝑥+1,𝑦)
;
up, its coordinates will be (𝑥−1,𝑦)
.
If the chip is located by the wall of the board, and the action chosen by Petya moves it towards the wall, then the chip remains in its current position.

Note that several chips can be located in the same cell.

For each chip, Petya chose the position which it should visit. Note that it's not necessary for a chip to end up in this position.

Since Petya does not have a lot of free time, he is ready to do no more than 2𝑛𝑚
 actions.

You have to find out what actions Petya should do so that each chip visits the position that Petya selected for it at least once. Or determine that it is not possible to do this in 2𝑛𝑚
 actions.


### ideas
1. 考虑每一个chip，它需要的最短移动路径
2. 一旦它移动到了目的地，那么就不需要再考虑它了。
3. 假设chip移动的路径= xLyTuRvD
4. (L 和 R不会同时出现, T 和 D也不会同时出现)
5. 如果不考虑2 * n * m 的限制
6. 先考虑处理掉往左上移动的那部分，可以选择最短的L，然后移动最短的T，交替移动（知道把这部分都处理掉）
7. 那么这样子，可以在最长L+T的时刻内完成（这部分操作）
8. 但是这里有个问题，有一些chip会远离原来的位置（比如那些向右下移动的部分）
9. 如何维护这个状态呢？
10. 如果每次移动后，都去更新每个chip的状态, 那么就是 (n + m) * k； 似乎也ok 