### description

There are 𝑛
programmers choosing their favorite algorithm amongst 𝑚
different choice options. Before the first round, all 𝑚
options are available. In each round, every programmer makes a vote for one of the remaining algorithms. After the
round, only the algorithms with the maximum number of votes remain. The voting process ends when there is only one
option left. Determine whether the voting process can continue indefinitely or no matter how people vote, they will
eventually choose a single option after some finite amount of rounds?

### solution

First we need to notice, that in order to keep some amount of options indefinetely, this number has to be at least 2
and divide 𝑛
. Let's find the smallest such number 𝑑
. Now, if 𝑑≤𝑚
, let's always vote for the first 𝑑
options evenly. In the other case (𝑑>𝑚)
each round would force us to to decrease the number of remaining options, so eventually it will become one. So the
answer is YES if and only if 𝑑>𝑚
.

Now on how to find the number 𝑑
fast. Since 𝑑
is a divisor of 𝑛
we can say, that 𝑑
is the smallest divisor of 𝑛
not equal to 1
.

We can find the number 𝑑
using different approaches. A more straightforward one, is checking all the numbers from 2
up to 𝑛√
. If no divisors found, then 𝑛
is prime and 𝑑=𝑛
. This results in a 𝑂(𝑡⋅𝑛√)
solution.

The solution presented before is good, but not fast enough in some languages, like Python. We've decided not to cut it
off to not make the problem heavy in IO. We can optimize it via finding the smallest divisor using the sieve of
Eratosthenes. This would result in 𝑂(𝑛𝑙𝑜𝑔𝑛)
or even faster precomputation and 𝑂(1)
to answer a test case, so the total time complexity is 𝑂(𝑛𝑙𝑜𝑔𝑛+𝑡)
.

### reference

[source](https://codeforces.com/problemset/problem/1826/C)