This is an interactive problem.

The jury has hidden a permutation†
𝑝
of length 𝑛
.

In one query, you can pick two integers 𝑙
and 𝑟
(1≤𝑙<𝑟≤𝑛
) by paying (𝑟−𝑙)2
coins. In return, you will be given the number of inversions‡
in the subarray [𝑝𝑙,𝑝𝑙+1,…𝑝𝑟]
.

Find the index of the maximum element in 𝑝
by spending at most 5⋅𝑛2
coins.

Note: the grader is not adaptive: the permutation is fixed before any queries are made.

†
A permutation of length 𝑛
is an array consisting of 𝑛
distinct integers from 1
to 𝑛
in arbitrary order. For example, [2,3,1,5,4]
is a permutation, but [1,2,2]
is not a permutation (2
appears twice in the array), and [1,3,4]
is also not a permutation (𝑛=3
but there is 4
in the array).

‡
The number of inversions in an array is the number of pairs of indices (𝑖,𝑗)
such that 𝑖<𝑗
and 𝑎𝑖>𝑎𝑗
. For example, the array [10,2,6,3]
contains 4
inversions. The inversions are (1,2),(1,3),(1,4)
, and (3,4)
.

Input
Each test contains multiple test cases. The first line of input contains a single integer 𝑡
(1≤𝑡≤100
) — the number of test cases.

The only line of each test case contains a single integer 𝑛
(2≤𝑛≤2000
) — the length of the hidden permutation 𝑝
.

It is guaranteed that the sum of 𝑛
over all test cases does not exceed 2000
.

Interaction
The interaction for each test case begins by reading the integer 𝑛
.

To make a query, output "? l r" (1≤𝑙<𝑟≤𝑛
) without quotes. Afterwards, you should read one single integer — the answer for your query.

If you receive the integer −1
instead of an answer or a valid value of 𝑛
, it means your program has made an invalid query, has exceed the limit of queries, or has given an incorrect answer on
the previous test case. Your program must terminate immediately to receive a Wrong Answer verdict. Otherwise you can get
an arbitrary verdict because your solution will continue to read from a closed stream.

When you are ready to give the final answer, output "! i" (1≤𝑖≤𝑛
) without quotes — the index of the maximum of the hidden permutation. After solving a test case, your program should
move to the next one immediately. After solving all test cases, your program should be terminated immediately.

After printing a query do not forget to output end of line and flush the output. Otherwise, you will get Idleness limit
exceeded. To do this, use:

fflush(stdout) or cout.flush() in C++;
System.out.flush() in Java;
flush(output) in Pascal;
stdout.flush() in Python;
see documentation for other languages.
Hacks

To hack, use the following format:

The first line contains an integer 𝑡
(1≤𝑡≤100
) — the number of test cases.

The first line of each test case contains a single integer 𝑛
(2≤𝑛≤2000
) — the length of the hidden permutation 𝑝
.

The second line of each test case contains 𝑛
integers 𝑝1,𝑝2,…,𝑝𝑛
(1≤𝑝𝑖≤𝑛
), 𝑝
must be a permutation.

The sum of 𝑛
over all test cases should not exceed 2000
.

### thoughts

1. cost = (r - l) ^ 2
2. 如果每次让(r - l) = 1, 那么可以通过n-1的cost，知道相邻位置的大小关系
3. 然后这样子，就能选出至少一半的较大的值
4. query的结果表示的是，在这个区间内 inversions 的数量
5. 如果知道 (l, r)的结果，和 (l, r + 1)的结果代表的差，表示的是什么呢？
6. 增加的数量，表示的是在区间(l, r)中，比 p[r+1]大的个数
7. 似乎也没啥用
8. 如果回到第3步，从剩余的峰值里再找出峰值
9. 但是这个cost怎么算呢？

### solution

Let 𝑞(𝑙,𝑟)
be be the number of inversions in the subarray [𝑝𝑙,𝑝𝑙+1,…𝑝𝑟]
. If 𝑙=𝑟
, we have 𝑞(𝑙,𝑟)=0
, otherwise, 𝑞(𝑙,𝑟)=
is equal to the result of the query "? l r".

Let 𝑓(𝑙,𝑟)
calculate the index of the maximum value in 𝑝𝑙,𝑝𝑙+1,…,𝑝𝑟
. If 𝑙=𝑟
, we have 𝑓(𝑙,𝑟)=𝑙
.

Suppose 𝑙<𝑟
. Let 𝑖:=𝑓(𝑙,𝑚)
and 𝑗:=𝑓(𝑚+1,𝑟)
, where 𝑚:=⌊𝑙+𝑟2⌋
. Let's compare 𝑞(𝑙,𝑗−1)
and 𝑞(𝑙,𝑗)
.

If they are equal, 𝑝𝑗
is greater than all the elements in the subarray [𝑝𝑙,𝑝𝑙+1,…,𝑝𝑚]
, so 𝑓(𝑙,𝑟)=𝑗
.

If 𝑞(𝑙,𝑗−1)<𝑞(𝑙,𝑗)
, 𝑝𝑗
is not greater that all the elements in the subarray [𝑝𝑙,𝑝𝑙+1,…,𝑝𝑚]
, and thus, the maximum on the whole subarray is 𝑝𝑖
, so 𝑓(𝑙,𝑟)=𝑖
.

Note that the case 𝑞(𝑙,𝑗−1)>𝑞(𝑙,𝑗)
is impossible, since all inversions in the subarray [𝑝𝑙,𝑝𝑙+1,…,𝑝𝑗−1]
remain as inversions for the subarray [𝑝𝑙,𝑝𝑙+1,…,𝑝𝑗]
.

To find the values of 𝑞(𝑙,𝑗−1)
and 𝑞(𝑙,𝑗)
, we will use (𝑗−1−𝑙)2+(𝑗−𝑙)2≤(𝑟−𝑙)2+(𝑟−𝑙)2=2⋅(𝑟−𝑙)2
coins.

Let 𝑔𝑛
be the number of coins needed to find the maximum on a subarray of length 𝑛
. We will prove by induction that 𝑔𝑛≤4⋅𝑛2
for all natural 𝑛
.

Base case: 𝑛=1
, 𝑔1:=0≤4⋅𝑛2=4
.

Induction step: let 𝑚:=⌈𝑛2⌉
. From the statements above, we have:
𝑔𝑛≤2⋅(𝑛−1)2+𝑔𝑚+𝑔𝑛−𝑚≤
2⋅(𝑛−1)2+4⋅(𝑚2+(𝑛−𝑚)2)=
6𝑛2+8𝑚2+2−8𝑛𝑚−4𝑛=
4𝑛2+2⋅(𝑛2−4𝑛𝑚+4𝑚2)+2−4𝑛=
4𝑛2+2⋅(𝑛−2𝑚)2+2−4𝑛≤
4𝑛2+2⋅1+2−4𝑛=
4𝑛2+4−4𝑛≤4𝑛2
Thus, to calculate 𝑓(1,𝑛)
, the answer to our problem, we will use 𝑔𝑛≤4⋅𝑛2
coins, which comfortably fits into the problem's 5⋅𝑛2
coin limit.

