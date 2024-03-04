Polycarp is designing a level for a game. The level consists of 𝑛
segments on the number line, where the 𝑖
-th segment starts at the point with coordinate 𝑙𝑖
and ends at the point with coordinate 𝑟𝑖
.

The player starts the level at the point with coordinate 0
. In one move, they can move to any point that is within a distance of no more than 𝑘
. After their 𝑖
-th move, the player must land within the 𝑖
-th segment, that is, at a coordinate 𝑥
such that 𝑙𝑖≤𝑥≤𝑟𝑖
. This means:

After the first move, they must be inside the first segment (from 𝑙1
to 𝑟1
);
After the second move, they must be inside the second segment (from 𝑙2
to 𝑟2
);
...
After the 𝑛
-th move, they must be inside the 𝑛
-th segment (from 𝑙𝑛
to 𝑟𝑛
).
The level is considered completed if the player reaches the 𝑛
-th segment, following the rules described above. After some thought, Polycarp realized that it is impossible to
complete the level with some values of 𝑘
.

Polycarp does not want the level to be too easy, so he asks you to determine the minimum integer 𝑘
with which it is possible to complete the level.