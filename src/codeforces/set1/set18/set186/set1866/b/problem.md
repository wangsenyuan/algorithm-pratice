On the trip to campus during the mid semester exam period, Chaneka thinks of two positive integers 𝑋
and 𝑌
. Since the two integers can be very big, both are represented using their prime factorisations, such that:

𝑋=𝐴𝐵11×𝐴𝐵22×…×𝐴𝐵𝑁𝑁
(each 𝐴𝑖
is prime, each 𝐵𝑖
is positive, and 𝐴1<𝐴2<…<𝐴𝑁
)
𝑌=𝐶𝐷11×𝐶𝐷22×…×𝐶𝐷𝑀𝑀
(each 𝐶𝑗
is prime, each 𝐷𝑗
is positive, and 𝐶1<𝐶2<…<𝐶𝑀
)
Chaneka ponders about these two integers for too long throughout the trip, so Chaneka's friend commands her "Gece,
deh!" (move fast) in order to not be late for the exam.

Because of that command, Chaneka comes up with a problem, how many pairs of positive integers 𝑝
and 𝑞
such that LCM(𝑝,𝑞)=𝑋
and GCD(𝑝,𝑞)=𝑌
. Since the answer can be very big, output the answer modulo 998244353
.

Notes:

LCM(𝑝,𝑞)
is the smallest positive integer that is simultaneously divisible by 𝑝
and 𝑞
.
GCD(𝑝,𝑞)
is the biggest positive integer that simultaneously divides 𝑝
and 𝑞
.
Input
The first line contains a single integer 𝑁
(1≤𝑁≤105
) — the number of distinct primes in the prime factorisation of 𝑋
.

The second line contains 𝑁
integers 𝐴1,𝐴2,𝐴3,…,𝐴𝑁
(2≤𝐴1<𝐴2<…<𝐴𝑁≤2⋅106
; each 𝐴𝑖
is prime) — the primes in the prime factorisation of 𝑋
.

The third line contains 𝑁
integers 𝐵1,𝐵2,𝐵3,…,𝐵𝑁
(1≤𝐵𝑖≤105
) — the exponents in the prime factorisation of 𝑋
.

The fourth line contains a single integer 𝑀
(1≤𝑀≤105
) — the number of distinct primes in the prime factorisation of 𝑌
.

The fifth line contains 𝑀
integers 𝐶1,𝐶2,𝐶3,…,𝐶𝑀
(2≤𝐶1<𝐶2<…<𝐶𝑀≤2⋅106
; each 𝐶𝑗
is prime) — the primes in the prime factorisation of 𝑌
.

The sixth line contains 𝑀
integers 𝐷1,𝐷2,𝐷3,…,𝐷𝑀
(1≤𝐷𝑗≤105
) — the exponents in the prime factorisation of 𝑌
.

### thoughts

1. 如果a[i] = c[j] 那么 b[i] >= d[j]
2. 假设有两个数u, v 包含c[j], 那么它们中有一个至少包含c[j] d[j]个， 而另一个包含b[i]个
3. 