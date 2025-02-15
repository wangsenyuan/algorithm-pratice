There are 𝑛
 enemies in a line. The 𝑖
'th enemy from the left has health ℎ𝑖
 and is currently at position 𝑥𝑖
. Xilonen has an attack damage of 𝑚
, and you are ready to defeat the enemies with her.

Xilonen has a powerful "ground stomp" attack. Before you perform any attacks, you select an integer 𝑝
 and position Xilonen there (𝑝
 can be any integer position, including a position with an enemy currently). Afterwards, for each attack, she deals 𝑚
 damage to an enemy at position 𝑝
 (if there are any), 𝑚−1
 damage to enemies at positions 𝑝−1
 and 𝑝+1
, 𝑚−2
 damage to enemies at positions 𝑝−2
 and 𝑝+2
, and so on. Enemies that are at least a distance of 𝑚
 away from Xilonen take no damage from attacks.

Formally, if there is an enemy at position 𝑥
, she will deal max(0,𝑚−|𝑝−𝑥|)
 damage to that enemy each hit. Note that you may not choose a different 𝑝
 for different attacks.

Over all possible 𝑝
, output the minimum number of attacks Xilonen must perform to defeat at least 𝑘
 enemies. If it is impossible to find a 𝑝
 such that eventually at least 𝑘
 enemies will be defeated, output −1
 instead. Note that an enemy is considered to be defeated if its health reaches 0
 or below.

