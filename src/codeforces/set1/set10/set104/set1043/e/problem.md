Zibi is a competitive programming coach. There are 𝑛
competitors who want to be prepared well. The training contests are quite unusual – there are two people in a team, two
problems, and each competitor will code exactly one of them. Of course, people in one team will code different problems.

Rules of scoring also aren't typical. The first problem is always an implementation problem: you have to implement some
well-known algorithm very fast and the time of your typing is rated. The second one is an awful geometry task and you
just have to get it accepted in reasonable time. Here the length and difficulty of your code are important. After that,
Zibi will give some penalty points (possibly negative) for each solution and the final score of the team is the sum of
them (the less the score is, the better).

We know that the 𝑖
-th competitor will always have score 𝑥𝑖
when he codes the first task and 𝑦𝑖
when he codes the second task. We can assume, that all competitors know each other's skills and during the contest
distribute the problems in the way that minimizes their final score. Remember that each person codes exactly one problem
in a contest.

Zibi wants all competitors to write a contest with each other. However, there are 𝑚
pairs of people who really don't like to cooperate and they definitely won't write a contest together. Still, the coach
is going to conduct trainings for all possible pairs of people, such that the people in pair don't hate each other. The
coach is interested for each participant, what will be his or her sum of scores of all teams he trained in?

### ideas

1. 教练想知道每个选手在所有他参加的训练队中的得分总和是多少。
2. 也就是对于i，来说，假设他可以和(j, k, l..) 配对，那么在和他们配对的情况下，最小值的sum是多少
3. 考虑(xi, yi) (xj, yj)
4. 如果 xi + yj <= yi + xj 那我们就取xi + yj
5. xi - yi <= xj - yj
6. 所以对于给定的 i, 分成两类 xj - yj >= xi - yi 部分，取 xi + yj
7. 或者 xj - yj < xi - yi, 取 yi + xj
8. 所以可以按照 xi - yi 排序， 这些设置y, 然后逐步的进行插入新的值
9. 但是如何处理hate pair呢？
10. got it