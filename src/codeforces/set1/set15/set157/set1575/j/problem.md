Mr. Chanek has a new game called Dropping Balls. Initially, Mr. Chanek has a grid 𝑎
of size 𝑛×𝑚
Each cell (𝑥,𝑦)
contains an integer 𝑎𝑥,𝑦
denoting the direction of how the ball will move.

𝑎𝑥,𝑦=1
— the ball will move to the right (the next cell is (𝑥,𝑦+1)
);
𝑎𝑥,𝑦=2
— the ball will move to the bottom (the next cell is (𝑥+1,𝑦)
);
𝑎𝑥,𝑦=3
— the ball will move to the left (the next cell is (𝑥,𝑦−1)
).
Every time a ball leaves a cell (𝑥,𝑦)
, the integer 𝑎𝑥,𝑦
will change to 2
. Mr. Chanek will drop 𝑘
balls sequentially, each starting from the first row, and on the 𝑐1,𝑐2,…,𝑐𝑘
-th (1≤𝑐𝑖≤𝑚
) columns.

Determine in which column each ball will end up in (position of the ball after leaving the grid).

