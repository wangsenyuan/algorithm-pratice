Seryozha conducts a course dedicated to building a map of heights of Stepanovo recreation center. He laid a rectangle grid of size 𝑛×𝑚
 cells on a map (rows of grid are numbered from 1
 to 𝑛
 from north to south, and columns are numbered from 1
 to 𝑚
 from west to east). After that he measured the average height of each cell above Rybinsk sea level and obtained a matrix of heights of size 𝑛×𝑚
. The cell (𝑖,𝑗)
 lies on the intersection of the 𝑖
-th row and the 𝑗
-th column and has height ℎ𝑖,𝑗
.

Seryozha is going to look at the result of his work in the browser. The screen of Seryozha's laptop can fit a subrectangle of size 𝑎×𝑏
 of matrix of heights (1≤𝑎≤𝑛
, 1≤𝑏≤𝑚
). Seryozha tries to decide how the weather can affect the recreation center — for example, if it rains, where all the rainwater will gather. To do so, he is going to find the cell having minimum height among all cells that are shown on the screen of his laptop.

Help Seryozha to calculate the sum of heights of such cells for all possible subrectangles he can see on his screen. In other words, you have to calculate the sum of minimum heights in submatrices of size 𝑎×𝑏
 with top left corners in (𝑖,𝑗)
 over all 1≤𝑖≤𝑛−𝑎+1
 and 1≤𝑗≤𝑚−𝑏+1
.

Consider the sequence 𝑔𝑖=(𝑔𝑖−1⋅𝑥+𝑦)mod𝑧
. You are given integers 𝑔0
, 𝑥
, 𝑦
 and 𝑧
. By miraculous coincidence, ℎ𝑖,𝑗=𝑔(𝑖−1)⋅𝑚+𝑗−1
 ((𝑖−1)⋅𝑚+𝑗−1
 is the index).


### ideas
1. 考虑一个高度为b的条，当它从做往右移动的过程中，可以更新这个区域内的最小值