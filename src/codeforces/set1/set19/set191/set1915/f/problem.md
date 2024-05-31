There are 𝑛
 people on the number line; the 𝑖
-th person is at point 𝑎𝑖
 and wants to go to point 𝑏𝑖
. For each person, 𝑎𝑖<𝑏𝑖
, and the starting and ending points of all people are distinct. (That is, all of the 2𝑛
 numbers 𝑎1,𝑎2,…,𝑎𝑛,𝑏1,𝑏2,…,𝑏𝑛
 are distinct.)

All the people will start moving simultaneously at a speed of 1
 unit per second until they reach their final point 𝑏𝑖
. When two people meet at the same point, they will greet each other once. How many greetings will there be?

Note that a person can still greet other people even if they have reached their final point.

Input
The first line of the input contains a single integer 𝑡
 (1≤𝑡≤104
) — the number of test cases. The description of test cases follows.

The first line of each test case contains a single integer 𝑛
 (1≤𝑛≤2⋅105
) — the number of people.

Then 𝑛
 lines follow, the 𝑖
-th of which contains two integers 𝑎𝑖
 and 𝑏𝑖
 (−109≤𝑎𝑖<𝑏𝑖≤109
) — the starting and ending positions of each person.

For each test case, all of the 2𝑛
 numbers 𝑎1,𝑎2,…,𝑎𝑛,𝑏1,𝑏2,…,𝑏𝑛
 are distinct.

The sum of 𝑛
 over all test cases does not exceed 2⋅105
.

### ideas
1. 显然，如果两个人都在移动（他们方向一样，从左到右，速度一样），他们肯定是没法greet的
2. 所以只有当前面的那个人停下来了，他们才有机会greet
3. 所以，就是，在[l,r]中间其他的r值，
4. 这个不对，假设对于两个人a、b
5. 如果a[l] < b[l] 且 a[r] < b[r], 那么他们是不可能相遇的
6. 只有当 a[l] > b[l], a[r] < b[r]
7. 或者 a[l] < b[l], b[r] < a[r]时，他们才有机会相遇
8. 所以必须要以l降序处理