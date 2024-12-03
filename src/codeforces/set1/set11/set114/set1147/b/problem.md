Inaka has a disc, the circumference of which is 𝑛
 units. The circumference is equally divided by 𝑛
 points numbered clockwise from 1
 to 𝑛
, such that points 𝑖
 and 𝑖+1
 (1≤𝑖<𝑛
) are adjacent, and so are points 𝑛
 and 1
.

There are 𝑚
 straight segments on the disc, the endpoints of which are all among the aforementioned 𝑛
 points.

Inaka wants to know if her image is rotationally symmetrical, i.e. if there is an integer 𝑘
 (1≤𝑘<𝑛
), such that if all segments are rotated clockwise around the center of the circle by 𝑘
 units, the new image will be the same as the original one.

Input
The first line contains two space-separated integers 𝑛
 and 𝑚
 (2≤𝑛≤100000
, 1≤𝑚≤200000
) — the number of points and the number of segments, respectively.

The 𝑖
-th of the following 𝑚
 lines contains two space-separated integers 𝑎𝑖
 and 𝑏𝑖
 (1≤𝑎𝑖,𝑏𝑖≤𝑛
, 𝑎𝑖≠𝑏𝑖
) that describe a segment connecting points 𝑎𝑖
 and 𝑏𝑖
.

It is guaranteed that no segments coincide.