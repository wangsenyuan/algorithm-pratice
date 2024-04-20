Elections in Berland are coming. There are only two candidates — Alice and Bob.

The main Berland TV channel plans to show political debates. There are 𝑛
people who want to take part in the debate as a spectator. Each person is described by their influence and political
views. There are four kinds of political views:

supporting none of candidates (this kind is denoted as "00"),
supporting Alice but not Bob (this kind is denoted as "10"),
supporting Bob but not Alice (this kind is denoted as "01"),
supporting both candidates (this kind is denoted as "11").
The direction of the TV channel wants to invite some of these people to the debate. The set of invited spectators should
satisfy three conditions:

at least half of spectators support Alice (i.e. 2⋅𝑎≥𝑚
, where 𝑎
is number of spectators supporting Alice and 𝑚
is the total number of spectators),
at least half of spectators support Bob (i.e. 2⋅𝑏≥𝑚
, where 𝑏
is number of spectators supporting Bob and 𝑚
is the total number of spectators),
the total influence of spectators is maximal possible.
Help the TV channel direction to select such non-empty set of spectators, or tell that this is impossible.

Input
The first line contains integer 𝑛
(1≤𝑛≤4⋅105
) — the number of people who want to take part in the debate as a spectator.

These people are described on the next 𝑛
lines. Each line describes a single person and contains the string 𝑠𝑖
and integer 𝑎𝑖
separated by space (1≤𝑎𝑖≤5000
), where 𝑠𝑖
denotes person's political views (possible values — "00", "10", "01", "11") and 𝑎𝑖
— the influence of the 𝑖
-th person.

Output
Print a single integer — maximal possible total influence of a set of spectators so that at least half of them support
Alice and at least half of them support Bob. If it is impossible print 0 instead.

### ideas

1. 假设一个 11的人没有加入，此时公有m个人，其中a个支持Alice, b个支持Bob
2. 所以有 2 * a >= m, and 2 * b >= m
3. 那么加入那个 11 后 2 * (1 + a) = 2 + 2 * a >= m + 2 > m + 1
4. 所以 11 的人始终应该加入, 假设当前有m个人
5. 然后剩余的人假设加入了x个 10, y个01， z个00
6. n = m + x + y + z
7. 2 * (m + x) >= n = m + x + y + z => m + x >= y + z
8. 2 * (m + y) >= n => m + y >= x + z
9. 假设目前的配置是 m + x + y (不考虑z) 且x = y
10. 现在有一个10 也满足条件就有
11. m + x + 1 >= y
12. m + y >= x + 1
13. 两个条件合并 m >= 1
14. 如果有一个01， 那么它可以无条件加入
15. 所以，有一个结论是min(x, y) 的 人都可以加入进来