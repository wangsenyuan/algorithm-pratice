There are 𝑛
TV shows you want to watch. Suppose the whole time is split into equal parts called "minutes". The 𝑖
-th of the shows is going from 𝑙𝑖
-th to 𝑟𝑖
-th minute, both ends inclusive.

You need a TV to watch a TV show and you can't watch two TV shows which air at the same time on the same TV, so it is
possible you will need multiple TVs in some minutes. For example, if segments [𝑙𝑖,𝑟𝑖]
and [𝑙𝑗,𝑟𝑗]
intersect, then shows 𝑖
and 𝑗
can't be watched simultaneously on one TV.

Once you start watching a show on some TV it is not possible to "move" it to another TV (since it would be too
distracting), or to watch another show on the same TV until this show ends.

There is a TV Rental shop near you. It rents a TV for 𝑥
rupees, and charges 𝑦
(𝑦<𝑥
) rupees for every extra minute you keep the TV. So in order to rent a TV for minutes [𝑎;𝑏]
you will need to pay 𝑥+𝑦⋅(𝑏−𝑎)
.

You can assume, that taking and returning of the TV doesn't take any time and doesn't distract from watching other TV
shows. Find the minimum possible cost to view all shows. Since this value could be too large, print it modulo 109+7
.

Input
The first line contains integers 𝑛
, 𝑥
and 𝑦
(1≤𝑛≤105
, 1≤𝑦<𝑥≤109
) — the number of TV shows, the cost to rent a TV for the first minute and the cost to rent a TV for every subsequent
minute.

Each of the next 𝑛
lines contains two integers 𝑙𝑖
and 𝑟𝑖
(1≤𝑙𝑖≤𝑟𝑖≤109
) denoting the start and the end minute of the 𝑖
-th TV show.

Output
Print exactly one integer — the minimum cost to view all the shows taken modulo 109+7
.

### ideas

1. 假设到现在为止的最小cost是x
2. 遇到一个新的界面，需要使用电视，但是没有多余的电视
3. 此时有两种选择，一种是租一个新的电视，cost = x + (r - l) * y
4. 或者把之前(假装)返还的电视继续租用，
5. 此时最好使用在最近（晚）返还的电视， 此时的额外的cost = (r - rx) * y
6. 