Chouti is working on a strange math problem.

There was a sequence of 𝑛
positive integers 𝑥1,𝑥2,…,𝑥𝑛
, where 𝑛
is even. The sequence was very special, namely for every integer 𝑡
from 1
to 𝑛
, 𝑥1+𝑥2+...+𝑥𝑡
is a square of some integer number (that is, a perfect square).

Somehow, the numbers with odd indexes turned to be missing, so he is only aware of numbers on even positions, i.e.
𝑥2,𝑥4,𝑥6,…,𝑥𝑛
. The task for him is to restore the original sequence. Again, it's your turn to help him.

The problem setter might make mistakes, so there can be no possible sequence at all. If there are several possible
sequences, you can output any.

Input
The first line contains an even number 𝑛
(2≤𝑛≤105
).

The second line contains 𝑛2
positive integers 𝑥2,𝑥4,…,𝑥𝑛
(1≤𝑥𝑖≤2⋅105
).

Output
If there are no possible sequence, print "No".

Otherwise, print "Yes" and then 𝑛
positive integers 𝑥1,𝑥2,…,𝑥𝑛
(1≤𝑥𝑖≤1013
), where 𝑥2,𝑥4,…,𝑥𝑛
should be same as in input data. If there are multiple answers, print any.

Note, that the limit for 𝑥𝑖
is larger than for input data. It can be proved that in case there is an answer, there must be a possible sequence
satisfying 1≤𝑥𝑖≤1013
.

### ideas

1. 我们考虑 t, t + 1
2. pref(t) = a * a
3. pref(t + 1) = b * b
4. pref(t) - pref(t+1) = b * b - a * a = x[t+1]
5. x[t+1] = (b - a) * (b + a)
6. x[2], x[4]
7. x[1] 是一个平方数
8. x[1] + x[2] 也是一个平方数
9. x[2] =  (b - a) * (b + a)
10. 也就是说 b - a 和 b + a 是 x[2]的两个除数
11. 这样子，只能找到x[1] = a * a
12. 但是如何找到x[3]
13. 