Some numbers, carved in chaotic order, immediately attracted his attention. Imp rapidly proposed a guess that they are
the remainders of division of a number n by all integers i from 1 to k. Unfortunately, there are too many integers to
analyze for Imp.

Imp wants you to check whether all these remainders are distinct. Formally, he wants to check, if all , 1 ≤ i ≤ k, are
distinct, i. e. there is no such pair (i, j) that:

1 ≤ i<j ≤ k,
, where is the remainder of division x by y.

### ideas

1. 如果n是一个质数, 是不是有所有的rem都是不一样的？
2. 如果 n <= k => false (n % 1 = 0, n % n = 0)
3. n > k  (n是质数)
4. 如果要所有的rem不一样, 比如有 n % i = i - 1
5. n % 2 = 1 (n % 1 = 0)

### solution

Consider the way remainders are obtained.

Remainder k - 1 can be obtained only when n is taken modulo k.
Remainder k - 2 can either be obtained when taken modulo k - 1 or k. Since the remainder modulo k is already fixed, the
only opportunity left is k - 1.
Proceeding this way, we come to a conclusion that if answer exists, then = i - 1 holds.
This condition is equal to (n + 1) mod i = 0, i. e. (n + 1) should be divisible by all numbers between 1 and k. In other
words, (n + 1) must be divisible by their LCM. Following the exponential growth of LCM, we claim that when k is huge
enough, the answer won't exists (more precisely at k ≥ 43). And for small k we can solve the task naively.