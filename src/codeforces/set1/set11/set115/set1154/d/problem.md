There is a robot staying at 𝑋=0
 on the 𝑂𝑥
 axis. He has to walk to 𝑋=𝑛
. You are controlling this robot and controlling how he goes. The robot has a battery and an accumulator with a solar panel.

The 𝑖
-th segment of the path (from 𝑋=𝑖−1
 to 𝑋=𝑖
) can be exposed to sunlight or not. The array 𝑠
 denotes which segments are exposed to sunlight: if segment 𝑖
 is exposed, then 𝑠𝑖=1
, otherwise 𝑠𝑖=0
.

The robot has one battery of capacity 𝑏
 and one accumulator of capacity 𝑎
. For each segment, you should choose which type of energy storage robot will use to go to the next point (it can be either battery or accumulator). If the robot goes using the battery, the current charge of the battery is decreased by one (the robot can't use the battery if its charge is zero). And if the robot goes using the accumulator, the current charge of the accumulator is decreased by one (and the robot also can't use the accumulator if its charge is zero).

If the current segment is exposed to sunlight and the robot goes through it using the battery, the charge of the accumulator increases by one (of course, its charge can't become higher than it's maximum capacity).

If accumulator is used to pass some segment, its charge decreases by 1 no matter if the segment is exposed or not.

You understand that it is not always possible to walk to 𝑋=𝑛
. You want your robot to go as far as possible. Find the maximum number of segments of distance the robot can pass if you control him optimally.


### ideas
1. accumlator 可以被充电，所以如果可能的话，应该优先使用acc
2. battery只能这么多
3. 假设能够到达位置x, 那么必须满足条件 b + a + z >= x (z是中间充电的段)
4. 且还有个条件，假设在段i的时候，accumlator是满电的，也不能充电
5. 如果到了位置i是，如果b已经降到了0，只能使用i
6. 想象一个贪心的情况，在任何一个地方都尽量使用acc。并且记录s[i] = 1的段
7. 如果到了某个地方i, 此时 acc 等于0了，那么选择某个s[j] = 1的段，使用它来交换使用电池
8. 在s[j] = 1,这个位置肯定是 acc < k的最远那个地方
9. 假设交换后 b - 1, 此时acc 怎么变化？首先 acc + 1 （没有消耗且加电了）