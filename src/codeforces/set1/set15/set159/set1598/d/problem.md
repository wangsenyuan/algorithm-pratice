Monocarp is the coach of the Berland State University programming teams. He decided to compose a problemset for a
training session for his teams.

Monocarp has 𝑛
problems that none of his students have seen yet. The 𝑖
-th problem has a topic 𝑎𝑖
(an integer from 1
to 𝑛
) and a difficulty 𝑏𝑖
(an integer from 1
to 𝑛
). All problems are different, that is, there are no two tasks that have the same topic and difficulty at the same time.

Monocarp decided to select exactly 3
problems from 𝑛
problems for the problemset. The problems should satisfy at least one of two conditions (possibly, both):

the topics of all three selected problems are different;
the difficulties of all three selected problems are different.
Your task is to determine the number of ways to select three problems for the problemset.

### ideas

1. 那是不是限定一种条件会比较容易处理，
2. 假设限定必须有3个topics，在这个前提下，有多少ways呢？
3. 把problem按照topic分组，假设a[1]...a[n]是每组中的个数
4. dp[i, x] = 表示前面i组中选择x个topic的ways
5. dp[i, x] = dp[i - 1, x] + a[i] * dp[i-1, x - 1] (如果 x <= 3 and x > 0)
6. 这个可以按照difficulty再做一遍
7. 然后要 减去 同时选择3个topic，且3个level的方案数
8. 这个好像搞不定呐～
9. 如果以 {x, y}为key，从这样的里面随意选出3个，
10. {1, a}, {1, b}, {2, a}，什么情况下会出现不满足条件的情况呢？
11. 要么1出现了至少2次，且a出现了至少2次

### solution

There are many different ways to solve this problem, but, in my opinion, the easiest one is to count all possible
triples and subtract the number of bad triples.

The first part is easy — the number of ways to choose 3
elements out of 𝑛
is just 𝑛⋅(𝑛−1)⋅(𝑛−2)6
. The second part is a bit tricky.

What does it mean that the conditions in the statements are not fulfilled? There is a pair of problems with equal
difficulty, and there is a pair of problems with the same topic. Since all problems in the input are different, it means
that every bad triple has the following form: [(𝑥𝑎,𝑦𝑎),(𝑥𝑏,𝑦𝑎),(𝑥𝑎,𝑦𝑏)]
— i. e. there exists a problem such that it shares the difficulty with one of the other two problems, and the topic with
the remaining problem of the triple.

This observation allows us to calculate the number of bad triples as follows: we will iterate on the "central" problem (
the one that shares the topic with the second problem and the difficulty with the third problem). If we pick (𝑥𝑎,𝑦𝑎)
as the "central" problem, we need to choose the other two. Counting ways to choose the other problems is easy if we
precalculate the number of problems for each topic/difficulty: let 𝑐𝑛𝑡𝑇𝑥
be the number of problems with topic 𝑥
, and 𝑐𝑛𝑡𝐷𝑦
be the number of problems with difficulty 𝑦
; then, if we pick the problem (𝑥,𝑦)
as the "central one", there are 𝑐𝑛𝑡𝑇𝑥−1
ways to choose a problem that shares the topic with it, and 𝑐𝑛𝑡𝐷𝑦−1
ways to choose a problem that has the same difficulty — so, we have to subtract (𝑐𝑛𝑡𝑇𝑥−1)(𝑐𝑛𝑡𝐷𝑦−1)
from the answer for every problem (𝑥,𝑦)
.