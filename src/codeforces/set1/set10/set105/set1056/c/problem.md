If you say that Arkady is a bit old-fashioned playing checkers, you won't be right. There is also a modern computer game
Arkady and his friends are keen on. We won't discuss its rules, the only feature important to this problem is that each
player has to pick a distinct hero in the beginning of the game.

There are 2
teams each having 𝑛
players and 2𝑛
heroes to distribute between the teams. The teams take turns picking heroes: at first, the first team chooses a hero in
its team, after that the second team chooses a hero and so on. Note that after a hero is chosen it becomes unavailable
to both teams.

The friends estimate the power of the 𝑖
-th of the heroes as 𝑝𝑖
. Each team wants to maximize the total power of its heroes. However, there is one exception: there are 𝑚
pairs of heroes that are especially strong against each other, so when any team chooses a hero from such a pair, the
other team must choose the other one on its turn. Each hero is in at most one such pair.

This is an interactive problem. You are to write a program that will optimally choose the heroes for one team, while the
jury's program will play for the other team. Note that the jury's program may behave inefficiently, in this case you
have to take the opportunity and still maximize the total power of your team. Formally, if you ever have chance to reach
the total power of 𝑞
or greater regardless of jury's program choices, you must get 𝑞
or greater to pass a test.

Input
The first line contains two integers 𝑛
and 𝑚
(1≤𝑛≤103
, 0≤𝑚≤𝑛
) — the number of players in one team and the number of special pairs of heroes.

The second line contains 2𝑛
integers 𝑝1,𝑝2,…,𝑝2𝑛
(1≤𝑝𝑖≤103
) — the powers of the heroes.

Each of the next 𝑚
lines contains two integer 𝑎
and 𝑏
(1≤𝑎,𝑏≤2𝑛
, 𝑎≠𝑏
) — a pair of heroes that are especially strong against each other. It is guaranteed that each hero appears at most once
in this list.

The next line contains a single integer 𝑡
(1≤𝑡≤2
) — the team you are to play for. If 𝑡=1
, the first turn is yours, otherwise you have the second turn.

### ideas

1. 考虑p=1时的情况，总的=sum，如果对手获得了x，那么自己获得sum-x
2. 为了得到更大的值，应该尽量的减小x，
3. 但是唯一的控制，就是m对，如果不优先使用，就有可能会让对方先手
4. 先使用m对，并获得每对中的最大值，然后在剩余的里面，使用最大值即可
5. 如果p=2,如果对手使用了pair，没有选择，只能选择pair，否则的话，应该尽快的转换为使用pair