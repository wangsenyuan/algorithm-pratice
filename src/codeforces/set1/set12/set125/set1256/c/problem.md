There is a river of width 𝑛
. The left bank of the river is cell 0
 and the right bank is cell 𝑛+1
 (more formally, the river can be represented as a sequence of 𝑛+2
 cells numbered from 0
 to 𝑛+1
). There are also 𝑚
 wooden platforms on a river, the 𝑖
-th platform has length 𝑐𝑖
 (so the 𝑖
-th platform takes 𝑐𝑖
 consecutive cells of the river). It is guaranteed that the sum of lengths of platforms does not exceed 𝑛
.

You are standing at 0
 and want to reach 𝑛+1
 somehow. If you are standing at the position 𝑥
, you can jump to any position in the range [𝑥+1;𝑥+𝑑]
. However you don't really like the water so you can jump only to such cells that belong to some wooden platform. For example, if 𝑑=1
, you can jump only to the next position (if it belongs to the wooden platform). You can assume that cells 0
 and 𝑛+1
 belong to wooden platforms.

You want to know if it is possible to reach 𝑛+1
 from 0
 if you can move any platform to the left or to the right arbitrary number of times (possibly, zero) as long as they do not intersect each other (but two platforms can touch each other). It also means that you cannot change the relative order of platforms.

Note that you should move platforms until you start jumping (in other words, you first move the platforms and then start jumping).