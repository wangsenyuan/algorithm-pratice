The teachers of the Summer Informatics School decided to plant 𝑛
trees in a row, and it was decided to plant only oaks and firs. To do this, they made a plan, which can be represented
as a binary string 𝑠
of length 𝑛
. If 𝑠𝑖=0
, then the 𝑖
-th tree in the row should be an oak, and if 𝑠𝑖=1
, then the 𝑖
-th tree in the row should be a fir.

The day of tree planting is tomorrow, and the day after tomorrow an inspector will come to the School. The inspector
loves nature very much, and he will evaluate the beauty of the row as follows:

First, he will calculate 𝑙0
as the maximum number of consecutive oaks in the row (the maximum substring consisting of zeros in the plan 𝑠
). If there are no oaks in the row, then 𝑙0=0
.
Then, he will calculate 𝑙1
as the maximum number of consecutive firs in the row (the maximum substring consisting of ones in the plan 𝑠
). If there are no firs in the row, then 𝑙1=0
.
Finally, he will calculate the beauty of the row as 𝑎⋅𝑙0+𝑙1
for some 𝑎
— the inspector's favourite number.
The teachers know the value of the parameter 𝑎
, but for security reasons they cannot tell it to you. They only told you that 𝑎
is an integer from 1
to 𝑛
.

Since the trees have not yet been planted, the teachers decided to change the type of no more than 𝑘
trees to the opposite (i.e., change 𝑠𝑖
from 0
to 1
or from 1
to 0
in the plan) in order to maximize the beauty of the row of trees according to the inspector.

For each integer 𝑗
from 1
to 𝑛
answer the following question independently:

What is the maximum beauty of the row of trees that the teachers can achieve by changing the type of no more than 𝑘
trees if the inspector's favourite number 𝑎
is equal to 𝑗
?

### thoughts

1. a = 1, beauty = l0 + l1, 这时候，应该尽可能的增大l0和l1
2. 可以计算在修改k次的情况下，从当前位置r能够到达的最远的l是多少
3. 这个可以用滑动窗口处理
4. 这样不大对， 因为通过修改0到1的时候，虽然获得了更大的l1， 却有可能减少了l0
5. 可以确定的一点时l0和l1是肯定分在两边的。倒是可以计算dp[i][j][0]表示在前i位中，使用j次修改，获得的最长的l0是多少
6. dp[i][j][1]表示获得最长的l1,
7. fp[i][j][0]/[1]表示获得的最长的后缀l0, l1
8. dp[i][j][0] + fp[i+1][k - j][1]
9. 对于a来说，感觉上时0越多越好，如果l0 = 1, l1 = n - 1
10. 假设a = 3, k次操作中，用了x次把1转化为0， y次把0转化为1
11. x + y = k
12. 假设就是将l1中的1转化为了l0中的0
13. 一次转化，相当于带来了2的收益

### solution

There are many various dynamic programming solutions of this problem. We will describe one of them. Let's calculate the
dynamics 𝑝𝑟𝑒𝑓𝑖, 𝑗
= the length of the longest subsegment of zeros that can be obtained on the prefix up to 𝑖
, which ends at index 𝑖
and costs exactly 𝑗
operations. Similarly, 𝑠𝑢𝑓𝑖, 𝑗
is the length of the longest subsegment of zeros on the suffix starting at 𝑖
, which starts at index 𝑖
and costs exactly 𝑗
operations. Such dynamics can be easily computed:
𝑝𝑟𝑒𝑓𝑖, 𝑗=⎧⎩⎨⎪⎪𝑝𝑟𝑒𝑓𝑖−1, 𝑗+1𝑝𝑟𝑒𝑓𝑖−1, 𝑗−1+10ififotherwise𝑠𝑖=0𝑠𝑖=1and𝑗>0
In the first case, we simply prolong the existing subsegment of zeros, in the second case, we change the current 1
to 0
, spending one operation on it (so if have 0
operations left (𝑗=0
), we cannot do anything and the value of the dynamics is 0
, meaning the segment has ended). 𝑠𝑢𝑓𝑖,𝑗
can be calculated similarly. Let's update both dynamics in such a way that 𝑝𝑟𝑒𝑓𝑖,𝑗
will mean the maximum length of a subsegment of zeros that ends no later than 𝑖
and costs no more than 𝑗
operations. This can be easily done by updating 𝑝𝑟𝑒𝑓𝑖,𝑗
with the value of 𝑝𝑟𝑒𝑓𝑖−1,𝑗
, and then with 𝑝𝑟𝑒𝑓𝑖,𝑗−1
. Similarly, we update the second dynamics.

Now let's consider a subsegment [𝑙, 𝑟]
that we want to convert into a segment of ones. We can easily calculate the number of operations 𝑥
that we will need (we'll just need to calculate the number of zeros in such a segment). Now, calculate the new dynamics
𝑑𝑝𝑙𝑒𝑛
for the length 𝑙𝑒𝑛=𝑟−𝑙+1
of the segment of ones, which equals the maximum length of a subsegment of zeros that we can obtain. Update this value
with max(𝑝𝑟𝑒𝑓𝑙−1,𝑘−𝑥,𝑠𝑢𝑓𝑟+1,𝑘−𝑥)
.

Then, to answer the question for a fixed number 𝑎
, we can iterate over the length 𝑙𝑒𝑛
of the segment of ones that will be in our answer and update the answer with the value 𝑎⋅𝑑𝑝𝑙𝑒𝑛+𝑙𝑒𝑛
, if there exists a value for 𝑙𝑒𝑛
in the dynamics 𝑑𝑝
.

The complexity is 𝑂(𝑛𝑘+𝑛2)
. Solutions with complexity 𝑂(𝑛𝑘log𝑛)
and 𝑂(𝑛𝑘)
using various optimizations of the dynamics (convex hull trick, D&C) also exist.