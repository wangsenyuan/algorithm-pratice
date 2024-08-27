Koa the Koala is at the beach!

The beach consists (from left to right) of a shore, 𝑛+1
 meters of sea and an island at 𝑛+1
 meters from the shore.

She measured the depth of the sea at 1,2,…,𝑛
 meters from the shore and saved them in array 𝑑
. 𝑑𝑖
 denotes the depth of the sea at 𝑖
 meters from the shore for 1≤𝑖≤𝑛
.

Like any beach this one has tide, the intensity of the tide is measured by parameter 𝑘
 and affects all depths from the beginning at time 𝑡=0
 in the following way:

For a total of 𝑘
 seconds, each second, tide increases all depths by 1
.
Then, for a total of 𝑘
 seconds, each second, tide decreases all depths by 1
.
This process repeats again and again (ie. depths increase for 𝑘
 seconds then decrease for 𝑘
 seconds and so on ...).
Formally, let's define 0
-indexed array 𝑝=[0,1,2,…,𝑘−2,𝑘−1,𝑘,𝑘−1,𝑘−2,…,2,1]
 of length 2𝑘
. At time 𝑡
 (0≤𝑡
) depth at 𝑖
 meters from the shore equals 𝑑𝑖+𝑝[𝑡mod2𝑘]
 (𝑡mod2𝑘
 denotes the remainder of the division of 𝑡
 by 2𝑘
). Note that the changes occur instantaneously after each second, see the notes for better understanding.

At time 𝑡=0
 Koa is standing at the shore and wants to get to the island. Suppose that at some time 𝑡
 (0≤𝑡
) she is at 𝑥
 (0≤𝑥≤𝑛
) meters from the shore:

In one second Koa can swim 1
 meter further from the shore (𝑥
 changes to 𝑥+1
) or not swim at all (𝑥
 stays the same), in both cases 𝑡
 changes to 𝑡+1
.
As Koa is a bad swimmer, the depth of the sea at the point where she is can't exceed 𝑙
 at integer points of time (or she will drown). More formally, if Koa is at 𝑥
 (1≤𝑥≤𝑛
) meters from the shore at the moment 𝑡
 (for some integer 𝑡≥0
), the depth of the sea at this point  — 𝑑𝑥+𝑝[𝑡mod2𝑘]
  — can't exceed 𝑙
. In other words, 𝑑𝑥+𝑝[𝑡mod2𝑘]≤𝑙
 must hold always.
Once Koa reaches the island at 𝑛+1
 meters from the shore, she stops and can rest.
Note that while Koa swims tide doesn't have effect on her (ie. she can't drown while swimming). Note that Koa can choose to stay on the shore for as long as she needs and neither the shore or the island are affected by the tide (they are solid ground and she won't drown there).

Koa wants to know whether she can go from the shore to the island. Help her!

### ideas
1. 显然，如果有任何d[i] > l => false
2. d[i] + p[t % 2k] <= l
3. 且p是一个连续变化的序列，所以可以求出i的一个连续的安全时间段
4. 如果所有的点的，安全时间段，可以重叠移动，那么就是ok的