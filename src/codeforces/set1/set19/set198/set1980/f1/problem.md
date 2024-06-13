Alice and Bob are dividing the field. The field is a rectangle of size 𝑛×𝑚
 (2≤𝑛,𝑚≤109
), the rows are numbered from 1
 to 𝑛
 from top to bottom, and the columns are numbered from 1
 to 𝑚
 from left to right. The cell at the intersection of row 𝑟
 and column 𝑐
 is denoted as (𝑟,𝑐
).

Bob has 𝑘
 (2≤𝑘≤2⋅105
) fountains, all of them are located in different cells of the field. Alice is responsible for dividing the field, but she must meet several conditions:

To divide the field, Alice will start her path in any free (without a fountain) cell on the left or top side of the field and will move, each time moving to the adjacent cell down or right. Her path will end on the right or bottom side of the field.
Alice's path will divide the field into two parts — one part will belong to Alice (this part includes the cells of her path), the other part — to Bob.
Alice will own the part that includes the cell (𝑛,1
).
Bob will own the part that includes the cell (1,𝑚
).
Alice wants to divide the field in such a way as to get as many cells as possible.

Bob wants to keep ownership of all the fountains, but he can give one of them to Alice. First, output the integer 𝛼
 — the maximum possible size of Alice's plot, if Bob does not give her any fountain (i.e., all fountains will remain on Bob's plot). Then output 𝑘
 non-negative integers 𝑎1,𝑎2,…,𝑎𝑘
, where:

𝑎𝑖=0
, if after Bob gives Alice the 𝑖
-th fountain, the maximum possible size of Alice's plot does not increase (i.e., remains equal to 𝛼
);
𝑎𝑖=1
, if after Bob gives Alice the 𝑖
-th fountain, the maximum possible size of Alice's plot increases (i.e., becomes greater than 𝛼
).
Input
The first line contains a single integer 𝑡
 (1≤𝑡≤104
) — the number of test cases.

The first line of each test case contains three integers 𝑛
, 𝑚
, and 𝑘
 (2≤𝑛,𝑚≤109
, 2≤𝑘≤2⋅105
) — the field sizes and the number of fountains, respectively.

Then follow 𝑘
 lines, each containing two numbers 𝑟𝑖
 and 𝑐𝑖
 (1≤𝑟𝑖≤𝑛
, 1≤𝑐𝑖≤𝑚
) — the coordinates of the cell with the 𝑖
-th fountain. It is guaranteed that all cells are distinct and none of them is (𝑛,1
).

It is guaranteed that the sum of 𝑘
 over all test cases does not exceed 2⋅105
.

Output
For each test case, first output 𝛼
 — the maximum size of the plot that can belong to Alice if Bob does not give her any of the fountains. Then output 𝑘
 non-negative integers 𝑎1,𝑎2,…,𝑎𝑘
, where:

𝑎𝑖=0
, if after Bob gives Alice the 𝑖
-th fountain, the maximum possible size of Alice's plot does not increase compared to the case when all 𝑘
 fountains belong to Bob;
𝑎𝑖=1
, if after Bob gives Alice the 𝑖
-th fountain, the maximum possible size of Alice's plot increases compared to the case when all 𝑘
 fountains belong to Bob.

 ### ideas
 1. 同一列里面，只有最底部的那个有关系，同一列里面，最左边的那个才有关系
 2. 所以，移除掉i后，如果它是这一列里面最下面的，且它是这一行最左边的，增加上去
 3. 如果是这样，岂不是有点太简单了？
 4. 确实不大对， 如果它前面的，下限已经在自己的下面了，也是不行的
 5. 要找那些在边界上的点