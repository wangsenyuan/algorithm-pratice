A boy called Smilo is playing a new game! In the game, there are 𝑛
hordes of monsters, and the 𝑖
-th horde contains 𝑎𝑖
monsters. The goal of the game is to destroy all the monsters. To do this, you have two types of attacks and a combo
counter 𝑥
, initially set to 0
:

The first type: you choose a number 𝑖
from 1
to 𝑛
, such that there is at least one monster left in the horde with the number 𝑖
. Then, you kill one monster from horde number 𝑖
, and the combo counter 𝑥
increases by 1
.
The second type: you choose a number 𝑖
from 1
to 𝑛
, such that there are at least 𝑥
monsters left in the horde with number 𝑖
. Then, you use an ultimate attack and kill 𝑥
monsters from the horde with number 𝑖
. After that, 𝑥
is reset to zero.
Your task is to destroy all of the monsters, meaning that there should be no monsters left in any of the hordes. Smilo
wants to win as quickly as possible, so he wants to the minimum number of attacks required to win the game.

### thoughts

1. sort array a
2. 对于右端来说，使用x来destroy，似乎收益更好
3. 小于x的，也无法使用x来摧毁
4. x最对摧毁x，所以余出来的部分，也不能使用x
5. x是自己决定的～
6. 第一种攻击的收益是1，第二种攻击的收益是多少呢？
7. 假设有一个horde有y个monster，如果一次准备到y时，再发起攻击，x = y, 那么意味着至少要攻击y个额外的monster
8. 也就是x+1次的收益是 2 * x,
9. 也就是说x越大收益越高
10. 尽量的使得右边的被消灭掉