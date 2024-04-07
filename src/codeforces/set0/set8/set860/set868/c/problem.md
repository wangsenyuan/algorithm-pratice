Snark and Philip are preparing the problemset for the upcoming pre-qualification round for semi-quarter-finals. They
have a bank of n problems, and they want to select any non-empty subset of it as a problemset.

k experienced teams are participating in the contest. Some of these teams already know some of the problems. To make the
contest interesting for them, each of the teams should know at most half of the selected problems.

Determine if Snark and Philip can make an interesting problemset!

Input
The first line contains two integers n, k (1 ≤ n ≤ 105, 1 ≤ k ≤ 4) — the number of problems and the number of
experienced teams.

Each of the next n lines contains k integers, each equal to 0 or 1. The j-th number in the i-th line is 1 if j-th team
knows i-th problem and 0 otherwise.

Output
Print "YES" (quotes for clarity), if it is possible to make an interesting problemset, and "NO" otherwise.

You can print each character either upper- or lowercase ("YeS" and "yes" are valid when the answer is "YES").

### ideas

1. 先考虑任何一个人， 如果他知道的题目超过了一半，显然答案是no
2. 那貌似就一定可以啊， 只要把全部的都放进去就可以了
3. 考虑到样例2，前面的想法是错误的。因为可以不选那些大家都知道的
4. 居然一点想法都没得～
5. 如果一个题目，大家都知道，显然加入它没有好处
6. 如果存在一个题目，大家都不知道，用这个题目就好了
7. 如果不存在这样的一个题目， 那么就找到两个题目，如果能满足，也可以
8. 如果不存在就找3个题目
9. 那最多是不是就是k个题目呢？
10. 如果不存在k个题目，是不是就一定不存在答案呢？
11. 似乎是可以的，因为新加入的题目，肯定至少有一个人知道
12. 但是即使是从n个题目中选出k个，也是非常慢的

### solution

Let us show that if a solution exists, then there is always a solution that uses at most two problems. First, if there
is a problem not known to any of the teams, that we can just take this only problem in the set. Next, suppose that there
is a problem known only to one of the teams. If there is a problem this team doesn't know, then these two problems make
a good set. Otherwise, the team knows all the problems, hence we cannot find a good set.

In the rest case, each problem is known to at least two of the teams. Now, if there is a good set of problems, then each
of the problems in the set must be known to exactly two of the teams. Indeed, let pi be the number of teams that knows
the problem. If a good set contains k problems, then we must have , since otherwise we would have a team that knows more
than half of the problems by pigeonhole principle. We also have pi ≥ 2, hence , and only the case pi = 2 is possible.

At this point, if we can find a pair of problems with pi = 2 and non-intersecting set of teams, then we are done.
Otherwise, we can show that a good set does not exist by case analysis.

To avoid O(n2) solution, we can leave at most 24 problems with unique types (sets of teams) and do pairwise checking on
them. This solution has O(n) complexity.