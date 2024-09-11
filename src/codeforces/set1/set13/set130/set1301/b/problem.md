Dark is going to attend Motarack's birthday. Dark decided that the gift he is going to give to Motarack is an array 𝑎
 of 𝑛
 non-negative integers.

Dark created that array 1000
 years ago, so some elements in that array disappeared. Dark knows that Motarack hates to see an array that has two adjacent elements with a high absolute difference between them. He doesn't have much time so he wants to choose an integer 𝑘
 (0≤𝑘≤109
) and replaces all missing elements in the array 𝑎
 with 𝑘
.

Let 𝑚
 be the maximum absolute difference between all adjacent elements (i.e. the maximum value of |𝑎𝑖−𝑎𝑖+1|
 for all 1≤𝑖≤𝑛−1
) in the array 𝑎
 after Dark replaces all missing elements with 𝑘
.

Dark should choose an integer 𝑘
 so that 𝑚
 is minimized. Can you help him?


 ### ideas
 1. 先可以计算出一个m0, 就是那些没有被删除的连续的数字的最大的difference
 2. 然后将数组进行压缩后，将只剩下 x, -1, u, v, -1, y, .. 这样的结构
 3. 然后考察是否可以达到某个最大值是m的预期结果
 4. 如果m成立，那么也肯定有办法得到m+1的预期结果，只需要改变某个-1即可
 5. 因为-1只在两个数的中间，所以可以很容易的判断
 6. x, -1, y 那么就是要求, 假设这个数为z
 7. 当z >= x, y 时，  z - x <= m, z - m <= y, z <= min(x + m, y + m)即可
 8. 只有当 x < y, x + m + m < y 时，m才不成立