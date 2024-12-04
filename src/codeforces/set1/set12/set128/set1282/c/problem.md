Petya has come to the math exam and wants to solve as many problems as possible. He prepared and carefully studied the rules by which the exam passes.

The exam consists of 𝑛
 problems that can be solved in 𝑇
 minutes. Thus, the exam begins at time 0
 and ends at time 𝑇
. Petya can leave the exam at any integer time from 0
 to 𝑇
, inclusive.

All problems are divided into two types:

easy problems — Petya takes exactly 𝑎
 minutes to solve any easy problem;
hard problems — Petya takes exactly 𝑏
 minutes (𝑏>𝑎
) to solve any hard problem.
Thus, if Petya starts solving an easy problem at time 𝑥
, then it will be solved at time 𝑥+𝑎
. Similarly, if at a time 𝑥
 Petya starts to solve a hard problem, then it will be solved at time 𝑥+𝑏
.

For every problem, Petya knows if it is easy or hard. Also, for each problem is determined time 𝑡𝑖
 (0≤𝑡𝑖≤𝑇
) at which it will become mandatory (required). If Petya leaves the exam at time 𝑠
 and there is such a problem 𝑖
 that 𝑡𝑖≤𝑠
 and he didn't solve it, then he will receive 0
 points for the whole exam. Otherwise (i.e if he has solved all such problems for which 𝑡𝑖≤𝑠
) he will receive a number of points equal to the number of solved problems. Note that leaving at time 𝑠
 Petya can have both "mandatory" and "non-mandatory" problems solved.

For example, if 𝑛=2
, 𝑇=5
, 𝑎=2
, 𝑏=3
, the first problem is hard and 𝑡1=3
 and the second problem is easy and 𝑡2=2
. Then:

if he leaves at time 𝑠=0
, then he will receive 0
 points since he will not have time to solve any problems;
if he leaves at time 𝑠=1
, he will receive 0
 points since he will not have time to solve any problems;
if he leaves at time 𝑠=2
, then he can get a 1
 point by solving the problem with the number 2
 (it must be solved in the range from 0
 to 2
);
if he leaves at time 𝑠=3
, then he will receive 0
 points since at this moment both problems will be mandatory, but he will not be able to solve both of them;
if he leaves at time 𝑠=4
, then he will receive 0
 points since at this moment both problems will be mandatory, but he will not be able to solve both of them;
if he leaves at time 𝑠=5
, then he can get 2
 points by solving all problems.
Thus, the answer to this test is 2
.

Help Petya to determine the maximal number of points that he can receive, before leaving the exam.

### ideas
1. 这个题目费脑~~~
2. 假设Petya在时刻s离开，那么必须满足在时刻s前，所有的mandatory的题目都被解决了，否则还不如在时刻0离开
3. 离开的越晚，似乎成绩也好（所以可以二分离开的时间吗）
4. 确定s后，就可以知道在s前面有多少个easy/hard的mandatory的题目, 然后就可以看是否能够完成了
5. 但是这里有个严重的问题，可以在时刻1离开，不代表可以在时刻2离开