Bernard loves visiting Rudolf, but he is always running late. The problem is that Bernard has to cross the river on a ferry. Rudolf decided to help his friend solve this problem.

The river is a grid of 𝑛
 rows and 𝑚
 columns. The intersection of the 𝑖
-th row and the 𝑗
-th column contains the number 𝑎𝑖,𝑗
 — the depth in the corresponding cell. All cells in the first and last columns correspond to the river banks, so the depth for them is 0
.

The river may look like this.
Rudolf can choose the row (𝑖,1),(𝑖,2),…,(𝑖,𝑚)
 and build a bridge over it. In each cell of the row, he can install a support for the bridge. The cost of installing a support in the cell (𝑖,𝑗)
 is 𝑎𝑖,𝑗+1
. Supports must be installed so that the following conditions are met:

A support must be installed in cell (𝑖,1)
;
A support must be installed in cell (𝑖,𝑚)
;
The distance between any pair of adjacent supports must be no more than 𝑑
. The distance between supports (𝑖,𝑗1)
 and (𝑖,𝑗2)
 is |𝑗1−𝑗2|−1
.
Building just one bridge is boring. Therefore, Rudolf decided to build 𝑘
 bridges on consecutive rows of the river, that is, to choose some 𝑖
 (1≤𝑖≤𝑛−𝑘+1
) and independently build a bridge on each of the rows 𝑖,𝑖+1,…,𝑖+𝑘−1
. Help Rudolf minimize the total cost of installing supports.

### ideas.
1. 先计算每行的最小值，然后计算连续k行的最小值