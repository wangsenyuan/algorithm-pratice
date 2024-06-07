Bob has a grid with 3
 rows and 𝑛
 columns, each of which contains either 𝑎𝑖
 or −𝑎𝑖
 for some integer 1≤𝑖≤𝑛
. For example, one possible grid for 𝑛=4
 is shown below:

⎡⎣⎢⎢𝑎1−𝑎4𝑎1−𝑎2𝑎4𝑎2−𝑎3−𝑎1−𝑎2−𝑎2−𝑎3𝑎4⎤⎦⎥⎥
Alice and Bob play a game as follows:

Bob shows Alice his grid.
Alice gives Bob an array 𝑎1,𝑎2,…,𝑎𝑛
 of her choosing, whose elements are all −1
 or 1
.
Bob substitutes these values into his grid to make a grid of −1
s and 1
s.
Bob sorts the elements of each column in non-decreasing order.
Alice wins if all the elements in the middle row are 1
; otherwise, Bob wins.
For example, suppose Alice gives Bob the array [1,−1,−1,1]
 for the grid above. Then the following will happen (colors are added for clarity):

⎡⎣⎢⎢𝑎1−𝑎4𝑎1−𝑎2𝑎4𝑎2−𝑎3−𝑎1−𝑎2−𝑎2−𝑎3𝑎4⎤⎦⎥⎥−→−−−−−−[1,−1,−1,1]⎡⎣⎢⎢1−1111−11−11111⎤⎦⎥⎥−→−−−−−−−−−sort each column⎡⎣⎢⎢−111−111−111111⎤⎦⎥⎥.
Since the middle row is all 1
, Alice wins.

Given Bob's grid, determine whether or not Alice can choose the array 𝑎
 to win the game.

### ideas
1. 考虑在同一列里面的3个数a, b, c
2. 如果a,b,c对应3个不同的i, j, k, 那么要保证其中至少有两个是1
3. 如果a, c都是 a[i], 那么 a[i] = 1, (如果它是-1，就没法保证中间的那个是1)
4. 如果 a, c 都是 -a[i], 那么 a[i] = -1
5. 如果其中两个是 a[i], -a[i], 那么另外一个b肯定是1 (a[i]没有限制)
6. 所以，a, b, c 对应3个不同的i, j, k时，如何保证其中至少有两个1呢？
7. -a -> b, -a -> c, -b -> a, -b -> c, -c -> a... 类似这样子？
8. a和-a之间连一条边，表示他们的符号必须相反