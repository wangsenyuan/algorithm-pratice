Monocarp is playing a computer game. He starts the game being level 1
. He is about to fight 𝑛
 monsters, in order from 1
 to 𝑛
. The level of the 𝑖
-th monster is 𝑎𝑖
.

For each monster in the given order, Monocarp's encounter goes as follows:

if Monocarp's level is strictly higher than the monster's level, the monster flees (runs away);
otherwise, Monocarp fights the monster.
After every 𝑘
-th fight with a monster (fleeing monsters do not count), Monocarp's level increases by 1
. So, his level becomes 2
 after 𝑘
 monsters he fights, 3
 after 2𝑘
 monsters, 4
 after 3𝑘
 monsters, and so on.

You need to process 𝑞
 queries of the following form:

𝑖 𝑥
: will Monocarp fight the 𝑖
-th monster (or will this monster flee) if the parameter 𝑘
 is equal to 𝑥
?

### ideas
1. 对于i，如果k=2的时候，它逃走了，那么在k=1的时候，它是否也肯定逃走呢？
2. k=2的时候它逃走，表示在到达它之前，level已经升级到了至少a[i]+1
3. 那么也就是说用户至少消灭2 * a[i]个怪物前
4. 