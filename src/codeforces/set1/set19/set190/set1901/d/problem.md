Vasya is a sorcerer that fights monsters. Again. There are 𝑛
monsters standing in a row, the amount of health points of the 𝑖
-th monster is 𝑎𝑖
.

Vasya is a very powerful sorcerer who knows many overpowered spells. In this fight, he decided to use a chain lightning
spell to defeat all the monsters. Let's see how this spell works.

Firstly, Vasya chooses an index 𝑖
of some monster (1≤𝑖≤𝑛
) and the initial power of the spell 𝑥
. Then the spell hits monsters exactly 𝑛
times, one hit per monster. The first target of the spell is always the monster 𝑖
. For every target except for the first one, the chain lightning will choose a random monster who was not hit by the
spell and is adjacent to one of the monsters that already was hit. So, each monster will be hit exactly once. The first
monster hit by the spell receives 𝑥
damage, the second monster receives (𝑥−1)
damage, the third receives (𝑥−2)
damage, and so on.

Vasya wants to show how powerful he is, so he wants to kill all the monsters with a single chain lightning spell. The
monster is considered dead if the damage he received is not less than the amount of its health points. On the other
hand, Vasya wants to show he doesn't care that much, so he wants to choose the minimum initial power of the spell 𝑥
such that it kills all monsters, no matter which monster (among those who can get hit) gets hit on each step.

Of course, Vasya is a sorcerer, but the amount of calculations required to determine the optimal spell setup is way
above his possibilities, so you have to help him find the minimum spell power required to kill all the monsters.

Note that Vasya chooses the initial target and the power of the spell, other things should be considered random and
Vasya wants to kill all the monsters even in the worst possible scenario.

### thoughts

1. 对于i，如果第一个目标选在它的左边，那么它有可能会是第i个(下标从1开始), 初始的伤害至少是 a[i] + i - 1
2. 如果第一个目标出现在它的右边，那么初始的伤害至少是a[i] + (n - i)
3. 再翻转一下，就是如果第一个目标是i, 那么它左边需要的最大的伤害是 a[j] + n - j, 右边需要的伤害是 a[j] + j-1, 取其中的最大值
4. 然后选择其中的最小值