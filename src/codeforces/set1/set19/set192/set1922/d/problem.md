Monocarp is playing a computer game (yet again). Guess what is he doing? That's right, killing monsters.

There are 𝑛
 monsters in a row, numbered from 1
 to 𝑛
. The 𝑖
-th monster has two parameters: attack value equal to 𝑎𝑖
 and defense value equal to 𝑑𝑖
. In order to kill these monsters, Monocarp put a berserk spell on them, so they're attacking each other instead of Monocarp's character.

The fight consists of 𝑛
 rounds. Every round, the following happens:

first, every alive monster 𝑖
 deals 𝑎𝑖
 damage to the closest alive monster to the left (if it exists) and the closest alive monster to the right (if it exists);
then, every alive monster 𝑗
 which received more than 𝑑𝑗
 damage during this round dies. I. e. the 𝑗
-th monster dies if and only if its defense value 𝑑𝑗
 is strictly less than the total damage it received this round.
For each round, calculate the number of monsters that will die during that round.


### ideas
1. 能不能反过来计算每个monster能够坚持到第几轮？
2. 这样好像挺麻烦的，就是能够杀死当前monster的是它相邻的两个，但这两个有可能会提前挂掉。情况就变得复杂了
3. d[i]不会减少
4. d[i] < a[prev[i]] + a[next[i]] 这样的会被清理掉
5. 当d[i]被清理掉后，需要知道它新的前继和后继
6. 且此时需要更新的也就是前继和后继
7. 应该还是 o(n) * o(lg)