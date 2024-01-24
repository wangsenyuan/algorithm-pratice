A portal of dark forces has opened at the border of worlds, and now the whole world is under a terrible threat. To close
the portal and save the world, you need to defeat 𝑛
monsters that emerge from the portal one after another.

Only the sorceress Vika can handle this. She possesses two magical powers — water magic and fire magic. In one second,
Vika can generate 𝑤
units of water mana and 𝑓
units of fire mana. She will need mana to cast spells. Initially Vika have 0
units of water mana and 0
units of fire mana.

Each of the 𝑛
monsters that emerge from the portal has its own strength, expressed as a positive integer. To defeat the 𝑖
-th monster with strength 𝑠𝑖
, Vika needs to cast a water spell or a fire spell of at least the same strength. In other words, Vika can spend at
least 𝑠𝑖
units of water mana on a water spell, or at least 𝑠𝑖
units of fire mana on a fire spell.

Vika can create and cast spells instantly. Vika can cast an unlimited number of spells every second as long she has
enough mana for that.

The sorceress wants to save the world as quickly as possible, so tell her how much time she will need.

### thoughts

1. 注意到n 比较小(<= 100), n * s <= 1e6
2. 如果有足够的水魔法，或者火魔法，就可以一次性消灭掉所有的monster
3. 假设经过了时间t， 此时有W = w * t, F = f * t
4. 然后应该有一个集合 s1 sum(of s) <= W (使用水魔法攻击)
5. s2 sum(of s) <= F（使用火魔法攻击)
6. s2 = sum of all - s1
7. 是否存在一个dp[s1] s1 <= W and sum - s2 <= F
8. 使用bitset