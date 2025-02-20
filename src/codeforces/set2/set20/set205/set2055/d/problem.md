A crow is sitting at position 0
 of the number line. There are 𝑛
 scarecrows positioned at integer coordinates 𝑎1,𝑎2,…,𝑎𝑛
 along the number line. These scarecrows have been enchanted, allowing them to move left and right at a speed of 1
 unit per second.

The crow is afraid of scarecrows and wants to stay at least a distance of 𝑘
 ahead of the nearest scarecrow positioned at or before it. To do so, the crow uses its teleportation ability as follows:

Let 𝑥
 be the current position of the crow, and let 𝑦
 be the largest position of a scarecrow such that 𝑦≤𝑥
. If 𝑥−𝑦<𝑘
, meaning the scarecrow is too close, the crow will instantly teleport to position 𝑦+𝑘
.
This teleportation happens instantly and continuously. The crow will keep checking for scarecrows positioned at or to the left of him and teleport whenever one gets too close (which could happen at non-integral times). Note that besides this teleportation ability, the crow will not move on its own.

Your task is to determine the minimum time required to make the crow teleport to a position greater than or equal to ℓ
, assuming the scarecrows move optimally to allow the crow to reach its goal. For convenience, you are asked to output twice the minimum time needed for the crow to reach the target position ℓ
. It can be proven that this value will always be an integer.

Note that the scarecrows can start, stop, or change direction at any time (possibly at non-integral times).

### ideas
1. 假设经过一段时间后， crow到达了位置p，那么此时所有它左边的，不管距离，都应该往右边移动
2. 当左边的最靠近的移动到位p - k + 1 时，推动 crow 移动到 p + 1,直到遇到下一个a[?]位置时，会被一次性推到 a[?] + k的地方
3. 右边的应该向左移动（但不一定全部都移动）
4. 最好的效果就是，大家都距离正好时k，crow到达a[i]的时候，可以一次性的移动到 a[j], a[j] = a[i] + w * k
5. 当a[1] > 0 的时候， 必须先要花费a[1]的时候，让第一个移动到0的位置
6. 然后此时的位置是k，如果右边距离p的距离是a[2] - p, 左边是最近的是0
7. 如果 a[2] - k <= k, 那么a[2] 可以停留（花费小于k的时间内到到位置k）从而让下次push更快的发生
8. 否则 a[2] - k > k, 这时候就需要考虑，是让左边的最近的继续push，还是让a[2]花非 a[2] - k - p 的时间