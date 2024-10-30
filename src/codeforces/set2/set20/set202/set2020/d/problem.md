To play the game, Alice draws a straight line and marks 𝑛
 points on it, indexed from 1
 to 𝑛
. Initially, there are no arcs between the points, so they are all disjoint. After that, Alice performs 𝑚
 operations of the following type:

She picks three integers 𝑎𝑖
, 𝑑𝑖
 (1≤𝑑𝑖≤10
), and 𝑘𝑖
.
She selects points 𝑎𝑖,𝑎𝑖+𝑑𝑖,𝑎𝑖+2𝑑𝑖,𝑎𝑖+3𝑑𝑖,…,𝑎𝑖+𝑘𝑖⋅𝑑𝑖
 and connects each pair of these points with arcs.
After performing all 𝑚
 operations, she wants to know the number of connected components†
 these points form. Please help her find this number.


 ### ideas
 1. d is small enough