Thomas is sailing around an island surrounded by the ocean. The ocean and island can be represented by a grid with 𝑛
 rows and 𝑚
 columns. The rows are numbered from 1
 to 𝑛
 from top to bottom, and the columns are numbered from 1
 to 𝑚
 from left to right. The position of a cell at row 𝑟
 and column 𝑐
 can be represented as (𝑟,𝑐)
. Below is an example of a valid grid.

Example of a valid grid
There are three types of cells: island, ocean and underwater volcano. Cells representing the island are marked with a '#', cells representing the ocean are marked with a '.', and cells representing an underwater volcano are marked with a 'v'. It is guaranteed that there is at least one island cell and at least one underwater volcano cell. It is also guaranteed that the set of all island cells forms a single connected component†
 and the set of all ocean cells and underwater volcano cells forms a single connected component. Additionally, it is guaranteed that there are no island cells at the edge of the grid (that is, at row 1
, at row 𝑛
, at column 1
, and at column 𝑚
).

Define a round trip starting from cell (𝑥,𝑦)
 as a path Thomas takes which satisfies the following conditions:

The path starts and ends at (𝑥,𝑦)
.
If Thomas is at cell (𝑖,𝑗)
, he can go to cells (𝑖+1,𝑗)
, (𝑖−1,𝑗)
, (𝑖,𝑗−1)
, and (𝑖,𝑗+1)
 as long as the destination cell is an ocean cell or an underwater volcano cell and is still inside the grid. Note that it is allowed for Thomas to visit the same cell multiple times in the same round trip.
The path must go around the island and fully encircle it. Some path 𝑝
 fully encircles the island if it is impossible to go from an island cell to a cell on the grid border by only traveling to adjacent on a side or diagonal cells without visiting a cell on path 𝑝
. In the image below, the path starting from (2,2)
, going to (1,3)
, and going back to (2,2)
 the other way does not fully encircle the island and is not considered a round trip.
Example of a path that does not fully encircle the island
The safety of a round trip is the minimum Manhattan distance‡
 from a cell on the round trip to an underwater volcano (note that the presence of island cells does not impact this distance).

You have 𝑞
 queries. A query can be represented as (𝑥,𝑦)
 and for every query, you want to find the maximum safety of a round trip starting from (𝑥,𝑦)
. It is guaranteed that (𝑥,𝑦)
 is an ocean cell or an underwater volcano cell.

†
A set of cells forms a single connected component if from any cell of this set it is possible to reach any other cell of this set by moving only through the cells of this set, each time going to a cell with a common side.

‡
Manhattan distance between cells (𝑟1,𝑐1)
 and (𝑟2,𝑐2)
 is equal to |𝑟1−𝑟2|+|𝑐1−𝑐2|
.

### ideas
1. 通过query去查，显然是不行了
2. 换个角度，看，距离(>=)d的情况下，是否能够把岛都围起来
3. 如果可以，然后检查是否覆盖到了关心的点
4. 然后放开d，直到处理完所有的点

### solution
For each non-island cell (𝑖,𝑗)
, let 𝑑𝑖,𝑗
 be the minimum Manhattan distance of cell (𝑖,𝑗)
 to an underwater volcano. We can find all 𝑑𝑖,𝑗
 with a multisource BFS from all underwater volcanos. The danger of a round trip is the smallest value of 𝑑𝑢,𝑣
 over all (𝑢,𝑣)
 in the path.

Consider any island cell. We can take inspiration from how we check whether a point is in a polygon — if a point is inside the polygon then a ray starting from the point and going in any direction will intersect the polygon an odd number of times. Draw an imaginary line along the top border of the cell and extend it all the way to the right of the grid.

Image for observation
We can observe that an optimal round trip will always cross the line an odd number of times.

Using this observation, we can let our state be (row,column,parity of the number of times we crossed the line)
. Naively, we can binary search for our answer and BFS to check if (𝑥,𝑦,0)
 and (𝑥,𝑦,1)
 are connected. This solves the easy version of the problem.

To fully solve this problem, we can add states (and their corresponding edges to already added states) one at a time from highest 𝑑
 to lowest 𝑑
. For each query (𝑥,𝑦)
, we want to find the first time when (𝑥,𝑦,0)
 and (𝑥,𝑦,1)
 become connected. This is a classic DSU with small to large merging problem. In short, we drop a token labeled with the index of the query at both (𝑥,𝑦,0)
 and (𝑥,𝑦,1)
. Each time we merge, we also merge the sets of tokens small to large and check if merging has caused two tokens of the same label to be in the same component. The time complexity of our solution is 𝑂(𝑛𝑚log(𝑛𝑚)+𝑞log2𝑞)
 with the log(𝑛𝑚)
 coming from sorting the states or edges. Note that there exists a 𝑂((𝑛𝑚⋅𝛼(𝑛𝑚)+𝑞)⋅log(𝑛+𝑚))
 parallel binary search solution as well as a 𝑂(𝑛𝑚log(𝑛𝑚)+𝑞log(𝑛𝑚))
 solution that uses LCA queries on the Kruskal's reconstruction tree or min path queries on the MSTs. In fact, with offline LCA queries, we can reduce the complexity to 𝑂(𝑛𝑚⋅𝛼(𝑛𝑚)+𝑞)
.