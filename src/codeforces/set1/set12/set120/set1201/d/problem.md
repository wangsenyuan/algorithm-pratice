You are on the island which can be represented as a 𝑛×𝑚
 table. The rows are numbered from 1
 to 𝑛
 and the columns are numbered from 1
 to 𝑚
. There are 𝑘
 treasures on the island, the 𝑖
-th of them is located at the position (𝑟𝑖,𝑐𝑖)
.

Initially you stand at the lower left corner of the island, at the position (1,1)
. If at any moment you are at the cell with a treasure, you can pick it up without any extra time. In one move you can move up (from (𝑟,𝑐)
 to (𝑟+1,𝑐)
), left (from (𝑟,𝑐)
 to (𝑟,𝑐−1)
), or right (from position (𝑟,𝑐)
 to (𝑟,𝑐+1)
). Because of the traps, you can't move down.

However, moving up is also risky. You can move up only if you are in a safe column. There are 𝑞
 safe columns: 𝑏1,𝑏2,…,𝑏𝑞
. You want to collect all the treasures as fast as possible. Count the minimum number of moves required to collect all the treasures.


### ideas
1. 显然只能一直往上移动
2. 假设在某一行计算出来，停在那些安全列上的最小值，然后计算下一行
3. 从任何一列开始，必然是向着一个方向移动，然后到最远的宝藏点，然后再折回移动的另外一边最远的宝藏点
4. 所以，对于任何一个安全列i，到了下上一行以后，对于它左边的安全列j, 如果它右边存在宝藏的话，必须先处理右边的
5. 然后再处理左边的，然后移动到j，那么这个结果是 pr - pl + pr - i + j - pl 这个可以用来做区间update
6. 对于j来说 dp[i] - i + pr - pl + pr - pl + j, 要找到最小的 dp[i] - i