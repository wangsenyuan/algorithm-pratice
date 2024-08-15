This is an interactive problem.

This is an easy version of the problem. The difference from the hard version is that in the easy version 𝑡=1
 and the number of queries is limited to 20
.

Polycarp is playing a computer game. In this game, an array consisting of zeros and ones is hidden. Polycarp wins if he guesses the position of the 𝑘
-th zero from the left 𝑡
 times.

Polycarp can make no more than 20
 requests of the following type:

? 𝑙
 𝑟
 — find out the sum of all elements in positions from 𝑙
 to 𝑟
 (1≤𝑙≤𝑟≤𝑛
) inclusive.
In this (easy version) of the problem, this paragraph doesn't really make sense since 𝑡=1
 always. To make the game more interesting, each guessed zero turns into one and the game continues on the changed array. More formally, if the position of the 𝑘
-th zero was 𝑥
, then after Polycarp guesses this position, the 𝑥
-th element of the array will be replaced from 0
 to 1
. Of course, this feature affects something only for 𝑡>1
.

Help Polycarp win the game.

### ideas
1. binary search