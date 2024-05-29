In a well-known school, a physical education lesson took place. As usual, everyone was lined up and asked to settle in "the first–𝑘
-th" position.

As is known, settling in "the first–𝑘
-th" position occurs as follows: the first 𝑘
 people have numbers 1,2,3,…,𝑘
, the next 𝑘−2
 people have numbers 𝑘−1,𝑘−2,…,2
, the next 𝑘
 people have numbers 1,2,3,…,𝑘
, and so on. Thus, the settling repeats every 2𝑘−2
 positions. Examples of settling are given in the "Note" section.

The boy Vasya constantly forgets everything. For example, he forgot the number 𝑘
 described above. But he remembers the position he occupied in the line, as well as the number he received during the settling. Help Vasya understand how many natural numbers 𝑘
 fit under the given constraints.

Note that the settling exists if and only if 𝑘>1
. In particular, this means that the settling does not exist for 𝑘=1
.

Input
Each test consists of multiple test cases. The first line contains a single integer 𝑡
 (1≤𝑡≤100
) — the number of test cases. This is followed by the description of the test cases.

The only line of each test case contains two integers 𝑛
 and 𝑥
 (1≤𝑥<𝑛≤109
) — Vasya's position in the line and the number Vasya received during the settling.

Output
For each test case, output a single integer — the number of different 𝑘
 that fit under the given constraints.

It can be proven that under the given constraints, the answer is finite.

### ideas
1. x <= k 
2. 如果x是在前半段 n % (2 * k - 2) = x
3. n = 10, x = 2, 如果 k = 3, 10 % 4 = 2, 
4.   k = 5, 10 % 8 = 2
5.   k = 6, 10 % 10 = 0 ? (它处于后半段)
6.   如果 x 处于后半段 n % (2 * k - 2) = 2 * k - x
7.   如果 x = k, 那么 n % (2 * k - 2) = k
8. a % b = c => (a - c) % b = 0