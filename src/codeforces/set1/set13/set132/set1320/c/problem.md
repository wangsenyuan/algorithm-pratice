Roma is playing a new expansion for his favorite game World of Darkraft. He made a new character and is going for his first grind.

Roma has a choice to buy exactly one of 𝑛
 different weapons and exactly one of 𝑚
 different armor sets. Weapon 𝑖
 has attack modifier 𝑎𝑖
 and is worth 𝑐𝑎𝑖
 coins, and armor set 𝑗
 has defense modifier 𝑏𝑗
 and is worth 𝑐𝑏𝑗
 coins.

After choosing his equipment Roma can proceed to defeat some monsters. There are 𝑝
 monsters he can try to defeat. Monster 𝑘
 has defense 𝑥𝑘
, attack 𝑦𝑘
 and possesses 𝑧𝑘
 coins. Roma can defeat a monster if his weapon's attack modifier is larger than the monster's defense, and his armor set's defense modifier is larger than the monster's attack. That is, a monster 𝑘
 can be defeated with a weapon 𝑖
 and an armor set 𝑗
 if 𝑎𝑖>𝑥𝑘
 and 𝑏𝑗>𝑦𝑘
. After defeating the monster, Roma takes all the coins from them. During the grind, Roma can defeat as many monsters as he likes. Monsters do not respawn, thus each monster can be defeated at most one.

Thanks to Roma's excessive donations, we can assume that he has an infinite amount of in-game currency and can afford any of the weapons and armor sets. Still, he wants to maximize the profit of the grind. The profit is defined as the total coins obtained from all defeated monsters minus the cost of his equipment. Note that Roma must purchase a weapon and an armor set even if he can not cover their cost with obtained coins.

Help Roma find the maximum profit of the grind.


### ideas
1. bruteforce的做法，就是在选定武器i, 防护j的情况下，检查可以干掉多少monster
2. 也就是那些x[u] < a[i] and y[u] < b[j]的部分
3. 感觉是个矩形。
4. 假设已经处理了 (1, ?)的记录
5. 现在处理(2, ?); 如果 a[1] < a[2], 那么显然那些能够被1干掉的部分，也能被2干掉
6. 且会增加一部分原来不能处理的（但是此时，就和具体的j有关系了）
7. 放过来考虑，对于每个monster，找到它在武器中的上界和防具中的上界
8. 有一个观察就是，如果武器按照a[i]排序，那么ca[i]也应该是递增的，否则的话，使用新的更便宜，但是更厉害的武器，更好
9. 现在假设已经处理好了，武器和防具，具备越贵，值越大
10. 假设在处理武器i的时候，计算出来使用武器i和某个防具j，它可以取得最大的收益f(i), 且记录了此时kill的怪物的集合
11. 现在处理i+1, 那么有两种情况需要考虑
12. 第一种，就是显然在i+1时，能够被杀掉的快五，此时也是可以的；
13. 第二种，有一些新的怪物，原来处理不了，现在可以处理的（因为 a[i+1] > a[i]), 
14. 这类又分两种情况，a， 仍然适用原来的防具j；b，必须使用新的防具k
15. 所以，这里感觉有这样一种结构，就是在从i到i+1的时候，把所有那些新的能够杀死的怪物(x[k] >= a[i] and x[k] < a[i+1]) 的部分，激活在区间(那些b[j] > y[k])
16. 然后找出整个区间上的最大值
17. 所以有区间update + query的segment tree