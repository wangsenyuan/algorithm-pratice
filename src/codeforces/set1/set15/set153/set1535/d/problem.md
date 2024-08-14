2𝑘
 teams participate in a playoff tournament. The tournament consists of 2𝑘−1
 games. They are held as follows: first of all, the teams are split into pairs: team 1
 plays against team 2
, team 3
 plays against team 4
 (exactly in this order), and so on (so, 2𝑘−1
 games are played in that phase). When a team loses a game, it is eliminated, and each game results in elimination of one team (there are no ties). After that, only 2𝑘−1
 teams remain. If only one team remains, it is declared the champion; otherwise, 2𝑘−2
 games are played: in the first one of them, the winner of the game "1
 vs 2
" plays against the winner of the game "3
 vs 4
", then the winner of the game "5
 vs 6
" plays against the winner of the game "7
 vs 8
", and so on. This process repeats until only one team remains.

Let the string 𝑠
 consisting of 2𝑘−1
 characters describe the results of the games in chronological order as follows:

if 𝑠𝑖
 is 0, then the team with lower index wins the 𝑖
-th game;
if 𝑠𝑖
 is 1, then the team with greater index wins the 𝑖
-th game;
if 𝑠𝑖
 is ?, then the result of the 𝑖
-th game is unknown (any team could win this game).
Let 𝑓(𝑠)
 be the number of possible winners of the tournament described by the string 𝑠
. A team 𝑖
 is a possible winner of the tournament if it is possible to replace every ? with either 1 or 0 in such a way that team 𝑖
 is the champion.

You are given the initial state of the string 𝑠
. You have to process 𝑞
 queries of the following form:

𝑝
 𝑐
 — replace 𝑠𝑝
 with character 𝑐
, and print 𝑓(𝑠)
 as the result of the query.

 ### ideas
 1. 假设s中没有？，都是0/1, 那么i是否能够获胜的条件 = i
 2. 对于给定的一个i, 在[0:2**(k-1))中判断它能否获胜, 如果它是偶数s[i] = 0/?, else s[i] = 1/?
 3. 然后在下一个区间，判断，然后。。。
 4. 所以，搞成一棵树， 0的时候=左子树获胜的数量，1右子树获胜的数量, ? 左+右
 5. 因为高度为k，修改的时候，从该层开始，修改上去，即可