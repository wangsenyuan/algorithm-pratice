You are given
N strings
S
1
​
,S
2
​
,…,S
N
​
.

Find the minimum length of a string that contains all these strings as substrings.

Here, a string
S contains a string
T as a substring if
T can be obtained by deleting zero or more characters from the beginning and zero or more characters from the end of
S.

Constraints
N is an integer.
1≤N≤20
S
i
​
is a string consisting of lowercase English letters whose length is at least

1.

The total length of
S
1
​
,S
2
​
,…,S
N
​
is at most
2×10
5
.

### ideas

1. 几个观察，如果 a完全包含于b，那么可以省区a；或者a是b的后缀，c的前缀，也可能可以剩余a
2. 说可能是因为，只有 b + c才能包含a
3. 还有就是n比较小
4. dp[state][x] 表示以x结尾的，包含了字符串state的状态
5. 如果 dp[next][y] 表示y是最后一个
6. 那么加入y以后，所有被y已经包含的，但是未选择的状态，可以被更新，或者 x + y包含的那部分，也可以被更新掉
7. 但有个问题，y必须切好被y包含，或者是x+y的一部分，所以按照字符串长度倒序排
8. 现在主要的问题是 a + b 可能包含c，但是 a + b 不是简单的连接，它们可能有部分是重叠的
9. 即使是重叠后，仍然可能包含c
10. 比如 snuke + kensho => sunkensho 9
11. 但是如果存在 keke 这样的字符, 那么sunkekensho (11) 要好过 sunkensho + keke (13)
12. 也就是在重叠的时候，仍然会有一个要重叠多少的问题存在
13. 这个情况不用考虑的

### solution

First, if there exist
i and
j such that
S
i
​
is a substring of
S
j
​
, then the condition that “
S
i
​
is contained as a substring” is completely implied by the condition that “
S
j
​
is contained as a substring,” so one can ignore
S
i
​
. From now on, we assume there are no such
i and
j.

For two strings
X and
Y, we define
f(X,Y) as the maximum
k such that the last
k characters of
X coincide with the first
k characters of
Y.

Consider a string with the minimal length that contains all of
S
1
​
,S
2
​
,…,S
N
​
as substrings, and take a pair
(l
i
​
,r
i
​
) so that its
l
i
​
-th through
(r
i
​
−1)-th characters match
S
i
​
, for each
i.

Without loss of generality, we assume
l
1
​
≤l
2
​
≤⋯≤l
N
​
. Then, for all
i=1,2,…,N−1 it holds that:

l
i
​
<l
i+1
​
≤r
i
​
,
r
i
​
−l
i+1
​
=f(S
i
​
,S
i+1
​
).
Otherwise, it contradicts the assumption in the first paragraph or the minimality assumed in the second paragraph.

Therefore, when finding the minimum length of a string that contains all of
S
1
​
,S
2
​
,…,S
N
​
as substrings, we only have to consider the following type of strings:

for a permutation
(p
1
​
,p
2
​
,…,p
N
​
) of
(1,2,…,N), the string is a concatenation of
S
p
1
​

​
, the last
∣S
p
2
​

​
∣−f(S
p
1
​

​
,S
p
2
​

​
) characters of
S
p
2
​

​
, the last
∣S
p
3
​

​
∣−f(S
p
2
​

​
,S
p
3
​

​
) characters of
S
p
3
​

​
,
…, and the last
∣S
p
N
​

​
∣−f(S
p
N−1
​

​
,S
p
N
​

​
) characters of
S
p
N
​

​
.
Thus, the original problem is reduced to minimizing
i=1
∑
N−1
​
f(S
p
i
​

​
,S
p
i+1
​

​
) for a permutation
(p
1
​
,p
2
​
,…,p
N
​
) of
(1,2,…,N).

By precalculating
f
(
​
S
i
​
,S
j
​
) for all pairs
i,j using string algorithms such as Z algorithm, KMP (Knuth–Morris–Pratt) algorithm, or rolling hash, this problem can
be solved with bit DP (Dynamic Programming) just as the traveling salesman problem.

Hence, the problem has been solved in a total of about
O(N∑∣S
i
​
∣+2
N
N
2
).