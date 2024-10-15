A team of three programmers is going to play a contest. The contest consists of 𝑛
 problems, numbered from 1
 to 𝑛
. Each problem is printed on a separate sheet of paper. The participants have decided to divide the problem statements into three parts: the first programmer took some prefix of the statements (some number of first paper sheets), the third contestant took some suffix of the statements (some number of last paper sheets), and the second contestant took all remaining problems. But something went wrong — the statements were printed in the wrong order, so the contestants have received the problems in some random order.

The first contestant has received problems 𝑎1,1,𝑎1,2,…,𝑎1,𝑘1
. The second one has received problems 𝑎2,1,𝑎2,2,…,𝑎2,𝑘2
. The third one has received all remaining problems (𝑎3,1,𝑎3,2,…,𝑎3,𝑘3
).

The contestants don't want to play the contest before they redistribute the statements. They want to redistribute them so that the first contestant receives some prefix of the problemset, the third contestant receives some suffix of the problemset, and the second contestant receives all the remaining problems.

During one move, some contestant may give one of their problems to other contestant. What is the minimum number of moves required to redistribute the problems?

### ideas
1. 最后的结果，肯定是 x已经 (? <= x)的数属于第一个集合
2. y已经(? >= y)的数属于第三个集合，中间的数属于第二个集合