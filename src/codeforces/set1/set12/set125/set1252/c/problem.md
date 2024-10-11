Pathfinding is a task of finding a route between two points. It often appears in many problems. For example, in a GPS navigation software where a driver can query for a suggested route, or in a robot motion planning where it should find a valid sequence of movements to do some tasks, or in a simple maze solver where it should find a valid path from one point to another point. This problem is related to solving a maze.

The maze considered in this problem is in the form of a matrix of integers 𝐴
 of 𝑁×𝑁
. The value of each cell is generated from a given array 𝑅
 and 𝐶
 of 𝑁
 integers each. Specifically, the value on the 𝑖𝑡ℎ
 row and 𝑗𝑡ℎ
 column, cell (𝑖,𝑗)
, is equal to 𝑅𝑖+𝐶𝑗
. Note that all indexes in this problem are from 1
 to 𝑁
.

A path in this maze is defined as a sequence of cells (𝑟1,𝑐1),(𝑟2,𝑐2),…,(𝑟𝑘,𝑐𝑘)
 such that |𝑟𝑖−𝑟𝑖+1|+|𝑐𝑖−𝑐𝑖+1|=1
 for all 1≤𝑖<𝑘
. In other words, each adjacent cell differs only by 1
 row or only by 1
 column. An even path in this maze is defined as a path in which all the cells in the path contain only even numbers.

Given a tuple ⟨𝑟𝑎,𝑐𝑎,𝑟𝑏,𝑐𝑏⟩
 as a query, your task is to determine whether there exists an even path from cell (𝑟𝑎,𝑐𝑎)
 to cell (𝑟𝑏,𝑐𝑏)
. To simplify the problem, it is guaranteed that both cell (𝑟𝑎,𝑐𝑎)
 and cell (𝑟𝑏,𝑐𝑏)
 contain even numbers.

For example, let 𝑁=5
, 𝑅={6,2,7,8,3}
, and 𝐶={3,4,8,5,1}
. The following figure depicts the matrix 𝐴
 of 5×5
 which is generated from the given array 𝑅
 and 𝐶
.

### ideas
1. 偶数+偶数 => 偶数
2. 奇数+奇数 => 奇数
3. 奇数+偶数 => 奇数
4. row 和 col 可以直接变成0/1
5. 每次移动，要么沿着行移动，要么沿着列移动
6. 假设row[i]和row[i+1]的奇偶性不同，那么就不可能从沿着列从i移动到i+1
7. 对列也是同理
8. 那么对于行i来说，能够在这一行左右移动的列，是可以确定的
9. 如果在这一行移动到了某列j，能否继续转变成按照列移动呢？
10. row[i] = 0, row[i+1] = 1, 
11. 那么在行i左右移动的列，比如是col[?] = 0 的部分，也就是说，它永远不可能从row[i] => row[i+1]
12. 除非它们一开始就是可以移动的