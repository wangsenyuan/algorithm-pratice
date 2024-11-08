You are playing a computer card game called Splay the Sire. Currently you are struggling to defeat the final boss of the game.

The boss battle consists of 𝑛
 turns. During each turn, you will get several cards. Each card has two parameters: its cost 𝑐𝑖
 and damage 𝑑𝑖
. You may play some of your cards during each turn in some sequence (you choose the cards and the exact order they are played), as long as the total cost of the cards you play during the turn does not exceed 3
. After playing some (possibly zero) cards, you end your turn, and all cards you didn't play are discarded. Note that you can use each card at most once.

Your character has also found an artifact that boosts the damage of some of your actions: every 10
-th card you play deals double damage.

What is the maximum possible damage you can deal during 𝑛
 turns?

 ### ideas
 1. 一共n个block，每个block中，有少于1e5张牌
 2. 但是最多只能使用3张牌，（sum c <= 3)
 3. 第10，20...次攻击产生暴击效果
 4. 最大伤害
 5. dp[i][x] 表示第i个block完成后，且第i个block的最后一击的次数 mod 10 = x时的最大值
 6. 然后就是如何对当前block进行处理
 7. 如果选择c = 3, 的， 那么就是选择最大的1个，
 8. 如果选择c = (1, 2)的，也是最大的2个
 9. 如果是c = (1, 1, 1)的，还是最大的3个
 10. good