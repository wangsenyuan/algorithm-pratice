ZS the Coder and Chris the Baboon has explored Udayland for quite some time. They realize that it consists of n towns
numbered from 1 to n.

There are n directed roads in the Udayland. i-th of them goes from town i to some other town ai (ai ≠ i). ZS the Coder
can flip the direction of any road in Udayland, i.e. if it goes from town A to town B before the flip, it will go from
town B to town A after.

ZS the Coder considers the roads in the Udayland confusing, if there is a sequence of distinct towns A1, A2, ..., Ak (k>

1) such that for every 1 ≤ i<k there is a road from town Ai to town Ai + 1 and another road from town Ak to town A1. In
   other words, the roads are confusing if some of them form a directed cycle of some towns.

Now ZS the Coder wonders how many sets of roads (there are 2n variants) in initial configuration can he choose to flip
such that after flipping each road in the set exactly once, the resulting network will not be confusing.

Note that it is allowed that after the flipping there are more than one directed road from some town and possibly some
towns with no roads leading out of it, or multiple roads between any pair of cities.

### ideas

1. 这个city只有一个环，要找到这个环
2. 不在这个环上的road，可以随便flip
3. 在这个环上的road，要避免flip成同一个方向（顺时针或逆时针）
4. 如果它们的方向一致，减去2（不flip或者都flip)
5. 有不一致的情况，也是减去2