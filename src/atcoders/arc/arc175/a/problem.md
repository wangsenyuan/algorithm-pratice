N people are sitting around a round table, and numbered 
1 to 
N in a counterclockwise order. Each person has a dominant hand: left or right.

There are 
N spoons numbered 
1 to 
N on the round table, with one spoon placed between each pair of adjacent people. For each 
1≤i≤N, to the left and right of person 
i, there are spoons 
i and 
(i+1), respectively. Here, spoon 
(N+1) refers to spoon 
1.

Below is a diagram for 
N=4.



You are given a permutation 
(P 
1
​
 ,…,P 
N
​
 ) of 
(1,…,N). In the order 
i=1,…,N, person 
P 
i
​
  will act as follows:

If there is a spoon remaining on left or right side, they will take one of them.
If there are spoons remaining on both sides, they will take the spoon on the side of their dominant hand.
Otherwise, they do nothing.
You are also given a string 
S of length 
N consisting of L, R, and ?. Among the 
2 
N
  possible combinations of dominant hands, find how many satisfy all of the following conditions, modulo 
998244353:

If the 
i-th character of 
S is L, person 
i is left-handed.
If the 
i-th character of 
S is R, person 
i is right-handed.
When everyone has finished acting, everyone has taken a spoon.


### ideas
1. 不大理解这个题目， S的作用是什么呢？S[i] = 'L', P[i] 是左利的，
2. 所以，只需要考虑S[i] = '?'的情况吗？
3. 如果所有人都需要能拿到spoon，其实只有两种情况，一种是大家都拿左边的，一种是大家都拿右边的
4. 否则肯定有人拿不到 （2个人的时候比较特殊）
5. 所以第一个人拿到的比较重要
6. 依次处理，如果到了a[i], 