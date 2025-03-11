The football season has just ended in Berland. According to the rules of Berland football, each match is played between two teams. The result of each match is either a draw, or a victory of one of the playing teams. If a team wins the match, it gets 𝑤
 points, and the opposing team gets 0
 points. If the game results in a draw, both teams get 𝑑
 points.

The manager of the Berland capital team wants to summarize the results of the season, but, unfortunately, all information about the results of each match is lost. The manager only knows that the team has played 𝑛
 games and got 𝑝
 points for them.

You have to determine three integers 𝑥
, 𝑦
 and 𝑧
 — the number of wins, draws and loses of the team. If there are multiple answers, print any of them. If there is no suitable triple (𝑥,𝑦,𝑧)
, report about it.

Input
The first line contains four integers 𝑛
, 𝑝
, 𝑤
 and 𝑑
 (1≤𝑛≤1012,0≤𝑝≤1017,1≤𝑑<𝑤≤105)
 — the number of games, the number of points the team got, the number of points awarded for winning a match, and the number of points awarded for a draw, respectively. Note that 𝑤>𝑑
, so the number of points awarded for winning is strictly greater than the number of points awarded for draw.

Output
If there is no answer, print −1
.

Otherwise print three non-negative integers 𝑥
, 𝑦
 and 𝑧
 — the number of wins, draws and losses of the team. If there are multiple possible triples (𝑥,𝑦,𝑧)
, print any of them. The numbers should meet the following conditions:

𝑥⋅𝑤+𝑦⋅𝑑=𝑝
,
𝑥+𝑦+𝑧=𝑛
.


### ideas
1. 1 < d < w <1e5
2. 所以感觉可以迭代d or w？
3. 假设迭代d, p = x * w + y * d < x * w  + y * w = (x + y) * w
4. => (x + y) >= upper(p / w) => xy ... 的下界
5. p = x * w + y * d > x * d + y * d => (x + y)的上界
6. p / w, p / d
7. 使用 扩展gcd？