There is a toy building consisting of 𝑛
towers. Each tower consists of several cubes standing on each other. The 𝑖
-th tower consists of ℎ𝑖
cubes, so it has height ℎ𝑖
.

Let's define operation slice on some height 𝐻
as following: for each tower 𝑖
, if its height is greater than 𝐻
, then remove some top cubes to make tower's height equal to 𝐻
. Cost of one "slice" equals to the total number of removed cubes from all towers.

Let's name slice as good one if its cost is lower or equal to 𝑘
(𝑘≥𝑛
).

### ideas

1. greedy sucks