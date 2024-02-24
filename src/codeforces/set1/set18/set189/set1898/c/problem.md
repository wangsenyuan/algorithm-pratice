Elena has a grid formed by 𝑛
horizontal lines and 𝑚
vertical lines. The horizontal lines are numbered by integers from 1
to 𝑛
from top to bottom. The vertical lines are numbered by integers from 1
to 𝑚
from left to right. For each 𝑥
and 𝑦
(1≤𝑥≤𝑛
, 1≤𝑦≤𝑚
), the notation (𝑥,𝑦)
denotes the point at the intersection of the 𝑥
-th horizontal line and 𝑦
-th vertical line.

Two points (𝑥1,𝑦1)
and (𝑥2,𝑦2)
are adjacent if and only if |𝑥1−𝑥2|+|𝑦1−𝑦2|=1
.

The grid formed by 𝑛=4
horizontal lines and 𝑚=5
vertical lines.
Elena calls a sequence of points 𝑝1,𝑝2,…,𝑝𝑔
of length 𝑔
a walk if and only if all the following conditions hold:

The first point 𝑝1
in this sequence is (1,1)
.
The last point 𝑝𝑔
in this sequence is (𝑛,𝑚)
.
For each 1≤𝑖<𝑔
, the points 𝑝𝑖
and 𝑝𝑖+1
are adjacent.
Note that the walk may contain the same point more than once. In particular, it may contain point (1,1)
or (𝑛,𝑚)
multiple times.

There are 𝑛(𝑚−1)+(𝑛−1)𝑚
segments connecting the adjacent points in Elena's grid. Elena wants to color each of these segments in blue or red
color so that there exists a walk 𝑝1,𝑝2,…,𝑝𝑘+1
of length 𝑘+1
such that

out of 𝑘
segments connecting two consecutive points in this walk, no two consecutive segments have the same color (in other
words, for each 1≤𝑖<𝑘
, the color of the segment between points 𝑝𝑖
and 𝑝𝑖+1
differs from the color of the segment between points 𝑝𝑖+1
and 𝑝𝑖+2
).
Please find any such coloring or report that there is no such coloring.