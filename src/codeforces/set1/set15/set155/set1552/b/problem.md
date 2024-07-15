The Olympic Games have just started and Federico is eager to watch the marathon race.

There will be 𝑛
 athletes, numbered from 1
 to 𝑛
, competing in the marathon, and all of them have taken part in 5
 important marathons, numbered from 1
 to 5
, in the past. For each 1≤𝑖≤𝑛
 and 1≤𝑗≤5
, Federico remembers that athlete 𝑖
 ranked 𝑟𝑖,𝑗
-th in marathon 𝑗
 (e.g., 𝑟2,4=3
 means that athlete 2
 was third in marathon 4
).

Federico considers athlete 𝑥
 superior to athlete 𝑦
 if athlete 𝑥
 ranked better than athlete 𝑦
 in at least 3
 past marathons, i.e., 𝑟𝑥,𝑗<𝑟𝑦,𝑗
 for at least 3
 distinct values of 𝑗
.

Federico believes that an athlete is likely to get the gold medal at the Olympics if he is superior to all other athletes.

Find any athlete who is likely to get the gold medal (that is, an athlete who is superior to all other athletes), or determine that there is no such athlete.

### ideas
1. 假设 a > b, b > c， 能否得到 a > c?
2. 有没有传递性
3. a > b 表示在至少3次比赛中， a超过了b, 假设就是1， 2， 3
4. b > c => 假设， 3， 4， 5（只有一项重叠）那么 a 至少在比赛3中超过了c，但是在4，5红很可能不如c
5. 且c在1， 2中也超过了b，所以，此时很难确定a和c的关系
6. 但是，假设设定这样一个策略，就是目前的候选者是x, 如果遇到一个选手超过了x， 那么就把y作为新的候选者
7. 直到结束
8. 最后再检查一轮y是否是真正的winner？
9. 首选，如果y超过了x，那么x肯定不是winner的候选；这个对于中间的任何一个人都是成立的，且它们都有被打败的经历
10. 成立