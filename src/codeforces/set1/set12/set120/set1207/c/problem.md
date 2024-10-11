You are responsible for installing a gas pipeline along a road. Let's consider the road (for simplicity) as a segment [0,𝑛]
 on 𝑂𝑋
 axis. The road can have several crossroads, but for simplicity, we'll denote each crossroad as an interval (𝑥,𝑥+1)
 with integer 𝑥
. So we can represent the road as a binary string consisting of 𝑛
 characters, where character 0 means that current interval doesn't contain a crossroad, and 1 means that there is a crossroad.

Usually, we can install the pipeline along the road on height of 1
 unit with supporting pillars in each integer point (so, if we are responsible for [0,𝑛]
 road, we must install 𝑛+1
 pillars). But on crossroads we should lift the pipeline up to the height 2
, so the pipeline won't obstruct the way for cars.

We can do so inserting several zig-zag-like lines. Each zig-zag can be represented as a segment [𝑥,𝑥+1]
 with integer 𝑥
 consisting of three parts: 0.5
 units of horizontal pipe + 1
 unit of vertical pipe + 0.5
 of horizontal. Note that if pipeline is currently on height 2
, the pillars that support it should also have length equal to 2
 units.

 ### ideas
 1. dp[i][1/2] 表示到i为止，pipeline i 的高度是2时的最优解