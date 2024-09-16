There are 𝑛
 monsters standing in a row numbered from 1
 to 𝑛
. The 𝑖
-th monster has ℎ𝑖
 health points (hp). You have your attack power equal to 𝑎
 hp and your opponent has his attack power equal to 𝑏
 hp.

You and your opponent are fighting these monsters. Firstly, you and your opponent go to the first monster and fight it till his death, then you and your opponent go the second monster and fight it till his death, and so on. A monster is considered dead if its hp is less than or equal to 0
.

The fight with a monster happens in turns.

You hit the monster by 𝑎
 hp. If it is dead after your hit, you gain one point and you both proceed to the next monster.
Your opponent hits the monster by 𝑏
 hp. If it is dead after his hit, nobody gains a point and you both proceed to the next monster.
You have some secret technique to force your opponent to skip his turn. You can use this technique at most 𝑘
 times in total (for example, if there are two monsters and 𝑘=4
, then you can use the technique 2
 times on the first monster and 1
 time on the second monster, but not 2
 times on the first monster and 3
 times on the second monster).

Your task is to determine the maximum number of points you can gain if you use the secret technique optimally.


### ideas
1. 检查有多少个，可以用0次作弊，1次作弊，2次作弊。。。
2. 然后按照作弊次数升序排