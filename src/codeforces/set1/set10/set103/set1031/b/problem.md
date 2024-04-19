When Masha came to math classes today, she saw two integer sequences of length 𝑛−1
on the blackboard. Let's denote the elements of the first sequence as 𝑎𝑖
(0≤𝑎𝑖≤3
), and the elements of the second sequence as 𝑏𝑖
(0≤𝑏𝑖≤3
).

Masha became interested if or not there is an integer sequence of length 𝑛
, which elements we will denote as 𝑡𝑖
(0≤𝑡𝑖≤3
), so that for every 𝑖
(1≤𝑖≤𝑛−1
) the following is true:

𝑎𝑖=𝑡𝑖|𝑡𝑖+1
(where |
denotes the bitwise OR operation) and
𝑏𝑖=𝑡𝑖&𝑡𝑖+1
(where &
denotes the bitwise AND operation).
The question appeared to be too difficult for Masha, so now she asked you to check whether such a sequence 𝑡𝑖
of length 𝑛
exists. If it exists, find such a sequence. If there are multiple such sequences, find any of them.

Input
The first line contains a single integer 𝑛
(2≤𝑛≤105
) — the length of the sequence 𝑡𝑖
.

The second line contains 𝑛−1
integers 𝑎1,𝑎2,…,𝑎𝑛−1
(0≤𝑎𝑖≤3
) — the first sequence on the blackboard.

The third line contains 𝑛−1
integers 𝑏1,𝑏2,…,𝑏𝑛−1
(0≤𝑏𝑖≤3
) — the second sequence on the blackboard.

Output
In the first line print "YES" (without quotes), if there is a sequence 𝑡𝑖
that satisfies the conditions from the statements, and "NO" (without quotes), if there is no such sequence.

If there is such a sequence, on the second line print 𝑛
integers 𝑡1,𝑡2,…,𝑡𝑛
(0≤𝑡𝑖≤3
) — the sequence that satisfies the statements conditions.

If there are multiple answers, print any of them.

### ideas

1. a[n-1] = t[n] | t[n-1]
2. b[n-1] = t[n] & t[n-1]
3. 如果 b[n-1] 有a[n-1] 没有的位 => false
4. b[0] = t[0] & t[1]
5. b[1] = t[1] & t[2]
6. dp[i][x] = t[i] = x, 且满足 i...n的要求