You play a computer game. In this game, you lead a party of 𝑚
 heroes, and you have to clear a dungeon with 𝑛
 monsters. Each monster is characterized by its power 𝑎𝑖
. Each hero is characterized by his power 𝑝𝑖
 and endurance 𝑠𝑖
.

The heroes clear the dungeon day by day. In the beginning of each day, you choose a hero (exactly one) who is going to enter the dungeon this day.

When the hero enters the dungeon, he is challenged by the first monster which was not defeated during the previous days (so, if the heroes have already defeated 𝑘
 monsters, the hero fights with the monster 𝑘+1
). When the hero fights the monster, there are two possible outcomes:

if the monster's power is strictly greater than the hero's power, the hero retreats from the dungeon. The current day ends;
otherwise, the monster is defeated.
After defeating a monster, the hero either continues fighting with the next monster or leaves the dungeon. He leaves the dungeon either if he has already defeated the number of monsters equal to his endurance during this day (so, the 𝑖
-th hero cannot defeat more than 𝑠𝑖
 monsters during each day), or if all monsters are defeated — otherwise, he fights with the next monster. When the hero leaves the dungeon, the current day ends.

Your goal is to defeat the last monster. What is the minimum number of days that you need to achieve your goal? Each day you have to use exactly one hero; it is possible that some heroes don't fight the monsters at all. Each hero can be used arbitrary number of times.

### ideas
1. 假设消灭完前i个monster需要dp[i]天
2. 然后在一天那现在消灭接下来的w个， 那么必须满足两个条件
3.  必须存在一个hero，他的持续时间是w（如果是到最后的话，>= w)
4.  这个here的power，至少>= 这段的最大值a[?]
5. 如果每个hero都检查一遍，那么是 n * m 的复杂性，TLE
6. 有个观察是, 如果有个heor，更持久，且更强力，那么应该尽量的使用他
7. 这样子hero就有个排序，越强力的，持久力越短
8. 但是这样子，还是n * m 
9. a[i] 那么久应该选择刚好是 pow[j] >= a[i]的，然后尽量的远
10. 然后到达下个位置，再选择更大的，