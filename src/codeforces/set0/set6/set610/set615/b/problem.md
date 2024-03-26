This Christmas Santa gave Masha a magic picture and a pencil. The picture consists of n points connected by m segments (
they might cross in any way, that doesn't matter). No two segments connect the same pair of points, and no segment
connects the point to itself. Masha wants to color some segments in order paint a hedgehog. In Mashas mind every
hedgehog consists of a tail and some spines. She wants to paint the tail that satisfies the following conditions:

Only segments already presented on the picture can be painted;
The tail should be continuous, i.e. consists of some sequence of points, such that every two neighbouring points are
connected by a colored segment;
The numbers of points from the beginning of the tail to the end should strictly increase.
Masha defines the length of the tail as the number of points in it. Also, she wants to paint some spines. To do so,
Masha will paint all the segments, such that one of their ends is the endpoint of the tail. Masha defines the beauty of
a hedgehog as the length of the tail multiplied by the number of spines. Masha wants to color the most beautiful
hedgehog. Help her calculate what result she may hope to get.

Note that according to Masha's definition of a hedgehog, one segment may simultaneously serve as a spine and a part of
the tail (she is a little girl after all). Take a look at the picture for further clarifications.

### ideas

1. 一个tail是一个连续的一段，且数字是严格递增的
2. spin只能连接到tail上面两头（起点/或者终点）
3. 对于一个u 以它结尾的最长的tail是多少（这个很好算）* deg[u]