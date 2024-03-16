Casimir has a rectangular piece of paper with a checkered field of size 𝑛×𝑚
. Initially, all cells of the field are white.

Let us denote the cell with coordinates 𝑖
vertically and 𝑗
horizontally by (𝑖,𝑗)
. The upper left cell will be referred to as (1,1)
and the lower right cell as (𝑛,𝑚)
.

Casimir draws ticks of different sizes on the field. A tick of size 𝑑
(𝑑>0
) with its center in cell (𝑖,𝑗)
is drawn as follows:

First, the center cell (𝑖,𝑗)
is painted black.
Then exactly 𝑑
cells on the top-left diagonally to the center and exactly 𝑑
cells on the top-right diagonally to the center are also painted black.
That is all the cells with coordinates (𝑖−ℎ,𝑗±ℎ)
for all ℎ
between 0
and 𝑑
are painted. In particular, a tick consists of 2𝑑+1
black cells.
An already painted cell will remain black if painted again. Below you can find an example of the 4×9
box, with two ticks of sizes 2
and 3
.

You are given a description of a checkered field of size 𝑛×𝑚
. Casimir claims that this field came about after he drew some (possibly 0
) ticks on it. The ticks could be of different sizes, but the size of each tick is at least 𝑘
(that is, 𝑑≥𝑘
for all the ticks).

Determine whether this field can indeed be obtained by drawing some (possibly none) ticks of sizes 𝑑≥𝑘
or not.

### ideas

1. 题目理解错了，不是说每个d只能出现一次，而是出现的d >= k