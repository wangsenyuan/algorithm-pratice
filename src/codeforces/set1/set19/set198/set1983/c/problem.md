Alice, Bob and Charlie want to share a rectangular cake cut into 𝑛
 pieces. Each person considers every piece to be worth a different value. The 𝑖
-th piece is considered to be of value 𝑎𝑖
 by Alice, 𝑏𝑖
 by Bob and 𝑐𝑖
 by Charlie.

The sum over all 𝑎𝑖
, all 𝑏𝑖
 and all 𝑐𝑖
 individually is the same, equal to 𝑡𝑜𝑡
.

Given the values of each piece of the cake for each person, you need to give each person a contiguous slice of cake. In other words, the indices at the left and right ends of these subarrays (the slices given to each person) can be represented as (𝑙𝑎,𝑟𝑎)
, (𝑙𝑏,𝑟𝑏)
 and (𝑙𝑐,𝑟𝑐)
 respectively for Alice, Bob and Charlie. The division needs to satisfy the following constraints:

No piece is assigned to more than one person, i.e., no two subarrays among [𝑙𝑎,…,𝑟𝑎]
, [𝑙𝑏,…,𝑟𝑏]
 and [𝑙𝑐,…,𝑟𝑐]
 intersect.
∑𝑟𝑎𝑖=𝑙𝑎𝑎𝑖,∑𝑟𝑏𝑖=𝑙𝑏𝑏𝑖,∑𝑟𝑐𝑖=𝑙𝑐𝑐𝑖≥⌈𝑡𝑜𝑡3⌉
.
Here, the notation ⌈𝑎𝑏⌉
 represents ceiling division. It is defined as the smallest integer greater than or equal to the exact division of 𝑎
 by 𝑏
. In other words, it rounds up the division result to the nearest integer. For instance ⌈103⌉=4
 and ⌈153⌉=5
.