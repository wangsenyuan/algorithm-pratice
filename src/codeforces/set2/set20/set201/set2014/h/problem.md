Sheriff of Nottingham has organized a tournament in archery. It's the final round and Robin Hood is playing against Sheriff!

There are 𝑛
 targets in a row numbered from 1
 to 𝑛
. When a player shoots target 𝑖
, their score increases by 𝑎𝑖
 and the target 𝑖
 is destroyed. The game consists of turns and players alternate between whose turn it is. Robin Hood always starts the game, then Sheriff and so on. The game continues until all targets are destroyed. Both players start with score 0
.

At the end of the game, the player with most score wins and the other player loses. If both players have the same score, it's a tie and no one wins or loses. In each turn, the player can shoot any target that wasn't shot before. Both play optimally to get the most score possible.

Sheriff of Nottingham has a suspicion that he might lose the game! This cannot happen, you must help Sheriff. Sheriff will pose 𝑞
 queries, each specifying 𝑙
 and 𝑟
. This means that the game would be played only with targets 𝑙,𝑙+1,…,𝑟
, as others would be removed by Sheriff before the game starts.

For each query 𝑙
, 𝑟
, determine whether the Sheriff can not lose the game when only considering the targets 𝑙,𝑙+1,…,𝑟
.

### ideas
1. will not lose only when the length is even, and every item has even size equals values
2. 1 1 2 2 3 3 3 3 4 4 这样子的才有可能
3. 否则肯定是会输的
4. 