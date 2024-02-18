Initially, you have an empty multiset of segments. You need to process 𝑞
operations of two types:

+

𝑙
𝑟
— Add the segment (𝑙,𝑟)
to the multiset,
−
𝑙
𝑟
— Remove exactly one segment (𝑙,𝑟)
from the multiset. It is guaranteed that this segment exists in the multiset.
After each operation, you need to determine if there exists a pair of segments in the multiset that do not intersect. A
pair of segments (𝑙,𝑟)
and (𝑎,𝑏)
do not intersect if there does not exist a point 𝑥
such that 𝑙≤𝑥≤𝑟
and 𝑎≤𝑥≤𝑏
.