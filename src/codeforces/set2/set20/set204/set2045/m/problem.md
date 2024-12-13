You are given a grid of 𝑅
 rows (numbered from 1
 to 𝑅
 from north to south) and 𝐶
 columns (numbered from 1
 to 𝐶
 from west to east). Every cell in this grid is a square of the same size. The cell located at row 𝑟
 and column 𝑐
 is denoted as (𝑟,𝑐)
. Each cell can either be empty or have a mirror in one of the cell's diagonals. Each mirror is represented by a line segment. The mirror is type 1
 if it is positioned diagonally from the southwest corner to the northeast corner of the cell, or type 2
 for the other diagonal.

These mirrors follow the law of reflection, that is, the angle of reflection equals the angle of incidence. Formally, for type 1
 mirror, if a beam of light comes from the north, south, west, or east of the cell, then it will be reflected to the west, east, north, and south of the cell, respectively. Similarly, for type 2
 mirror, if a beam of light comes from the north, south, west, or east of the cell, then it will be reflected to the east, west, south, and north of the cell, respectively.


You want to put a laser from outside the grid such that all mirrors are hit by the laser beam. There are 2⋅(𝑅+𝐶)
 possible locations to put the laser:

from the north side of the grid at column 𝑐
, for 1≤𝑐≤𝐶
, shooting a laser beam to the south;
from the south side of the grid at column 𝑐
, for 1≤𝑐≤𝐶
, shooting a laser beam to the north;
from the east side of the grid at row 𝑟
, for 1≤𝑟≤𝑅
, shooting a laser beam to the west; and
from the west side of the grid at row 𝑟
, for 1≤𝑟≤𝑅
, shooting a laser beam to the east.

Determine all possible locations for the laser such that all mirrors are hit by the laser beam.

### ideas
1. type 1 从右上连接到左下
2. type 2 从左上连接到右下
3. 从任何一个位置开始，一直连接过去，如果能碰到所有的mirror，那么就可以算进去
4. 否则，光线肯定会在某个时刻离开了区域
5. 但是这样的复杂性可能是 (n + m) * n * m 差不多 1e7
6. 似乎也是可行的
7. 如果答案存在，那么肯定是2