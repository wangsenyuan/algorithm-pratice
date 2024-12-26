Mount ICPC can be represented as a grid of 𝑅
 rows (numbered from 1
 to 𝑅
) and 𝐶
 columns (numbered from 1
 to 𝐶
). The cell located at row 𝑟
 and column 𝑐
 is denoted as (𝑟,𝑐)
 and has a height of 𝐻𝑟,𝑐
. Two cells are adjacent to each other if they share a side. Formally, (𝑟,𝑐)
 is adjacent to (𝑟−1,𝑐)
, (𝑟+1,𝑐)
, (𝑟,𝑐−1)
, and (𝑟,𝑐+1)
, if any exists.

You can move only between adjacent cells, and each move comes with a penalty. With an aura of an odd positive integer 𝑋
, moving from a cell with height ℎ1
 to a cell with height ℎ2
 gives you a penalty of (ℎ1−ℎ2)𝑋
. Note that the penalty can be negative.

You want to answer 𝑄
 independent scenarios. In each scenario, you start at the starting cell (𝑅𝑠,𝐶𝑠)
 and you want to go to the destination cell (𝑅𝑓,𝐶𝑓)
 with minimum total penalty. In some scenarios, the total penalty might become arbitrarily small; such a scenario is called invalid. Find the minimum total penalty to move from the starting cell to the destination cell, or determine if the scenario is invalid.

 ### ideas
 1. 当x = 1, 答案 = 两个位置的高度差
 2. x = 2, 考虑从h1 -> h2, 这个时候它的penalty肯定是正数，所以存在一个最小值
 3. x = 3时，从h1->h2, 如果h1 < h2, 那么这个penatity 是个负数，然后选择一个h3, h1 < h3 < h2
 4. 因为 (h2 - h1)**3 > (h2 - h3) ** 3 + (h3 - h1) ** 3 (成立吗？)
 5. 就可以陷入一个无限变小的环，只要存在这样的情况， 任何x > 3 的都是无效的
 6. 搞错了， x就是odd数