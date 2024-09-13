Bashar was practicing for the national programming contest. Because of sitting too much in front of the computer without doing physical movements and eating a lot Bashar became much fatter. Bashar is going to quit programming after the national contest and he is going to become an actor (just like his father), so he should lose weight.

In order to lose weight, Bashar is going to run for 𝑘
 kilometers. Bashar is going to run in a place that looks like a grid of 𝑛
 rows and 𝑚
 columns. In this grid there are two one-way roads of one-kilometer length between each pair of adjacent by side cells, one road is going from the first cell to the second one, and the other road is going from the second cell to the first one. So, there are exactly (4𝑛𝑚−2𝑛−2𝑚)
 roads.
 Bashar wants to run by these rules:

He starts at the top-left cell in the grid;
In one move Bashar may go up (the symbol 'U'), down (the symbol 'D'), left (the symbol 'L') or right (the symbol 'R'). More formally, if he stands in the cell in the row 𝑖
 and in the column 𝑗
, i.e. in the cell (𝑖,𝑗)
 he will move to:
in the case 'U' to the cell (𝑖−1,𝑗)
;
in the case 'D' to the cell (𝑖+1,𝑗)
;
in the case 'L' to the cell (𝑖,𝑗−1)
;
in the case 'R' to the cell (𝑖,𝑗+1)
;
He wants to run exactly 𝑘
 kilometers, so he wants to make exactly 𝑘
 moves;
Bashar can finish in any cell of the grid;
He can't go out of the grid so at any moment of the time he should be on some cell;
Bashar doesn't want to get bored while running so he must not visit the same road twice. But he can visit the same cell any number of times.
Bashar asks you if it is possible to run by such rules. If it is possible, you should tell him how should he run.

You should give him 𝑎
 steps to do and since Bashar can't remember too many steps, 𝑎
 should not exceed 3000
. In every step, you should give him an integer 𝑓
 and a string of moves 𝑠
 of length at most 4
 which means that he should repeat the moves in the string 𝑠
 for 𝑓
 times. He will perform the steps in the order you print them.

For example, if the steps are 2
 RUD, 3
 UUL then the moves he is going to move are RUD +
 RUD +
 UUL +
 UUL +
 UUL =
 RUDRUDUULUULUUL.

Can you help him and give him a correct sequence of moves such that the total distance he will run is equal to 𝑘
 kilometers or say, that it is impossible?



### ideas
1. 很懵的一个题目
2. 什么情况下，不能完成k千米？ 4 * n * m - 2 * n - 2 * m < k 时，=> no
3. 然后是不是一定有答案呢？
4. 好想还要减掉 min(n, m) - 2, 假设n比较小，能够跑最多的策略就是 横着过去，然后再回来，到下一排（但是没法回去）
5. 然后重复，似乎没有比这个更长的距离了
6. 4 * n * m - 2 * n - 2 * m - min(n - 1, m - 1) >= k
7. 然后就重复就可以了？
8. 不是的，还要减去一个 min(n-1, m - 1)因为最后的那列的路径没法使用？
9. 似乎也不对。似乎还有更好的答案
10. 讨厌这种构造性的题目
11. 