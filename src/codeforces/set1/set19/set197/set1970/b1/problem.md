After some recent attacks on Hogwarts Castle by the Death Eaters, the Order of the Phoenix has decided to station 𝑛
 members in Hogsmead Village. The houses will be situated on a picturesque 𝑛×𝑛
 square field. Each wizard will have their own house, and every house will belong to some wizard. Each house will take up the space of one square.

However, as you might know wizards are very superstitious. During the weekends, each wizard 𝑖
 will want to visit the house that is exactly 𝑎𝑖
 (0≤𝑎𝑖≤𝑛)
 away from their own house. The roads in the village are built horizontally and vertically, so the distance between points (𝑥𝑖,𝑦𝑖)
 and (𝑥𝑗,𝑦𝑗)
 on the 𝑛×𝑛
 field is |𝑥𝑖−𝑥𝑗|+|𝑦𝑖−𝑦𝑗|
. The wizards know and trust each other, so one wizard can visit another wizard's house when the second wizard is away. The houses to be built will be big enough for all 𝑛
 wizards to simultaneously visit any house.

Apart from that, each wizard is mandated to have a view of the Hogwarts Castle in the north and the Forbidden Forest in the south, so the house of no other wizard should block the view. In terms of the village, it means that in each column of the 𝑛×𝑛
 field, there can be at most one house, i.e. if the 𝑖
-th house has coordinates (𝑥𝑖,𝑦𝑖)
, then 𝑥𝑖≠𝑥𝑗
 for all 𝑖≠𝑗
.

The Order of the Phoenix doesn't yet know if it is possible to place 𝑛
 houses in such a way that will satisfy the visit and view requirements of all 𝑛
 wizards, so they are asking for your help in designing such a plan.

If it is possible to have a correct placement, where for the 𝑖
-th wizard there is a house that is 𝑎𝑖
 away from it and the house of the 𝑖
-th wizard is the only house in their column, output YES, the position of houses for each wizard, and to the house of which wizard should each wizard go during the weekends.

If it is impossible to have a correct placement, output NO.


### ideas
1. 这个题目好拗口呐
2. 假设这些房子的位置是(x1, y1) (x2, y2) ... (xn, yn)
3. 因为x1....xn 不同，我们可以按照列排序（后面再考虑分配的问题）
4. 房子的位置，简化为(y1, y2, y3, .... yn)
5. a[i]先排个序
6. 对于i来说，如果要在a[i]的距离内存在一个另外一个房子
7. 如果a[i] <= n - 1
8. 如果有一个房子处在(1, 1)处，那么最远的距离是 (n - 1) * 2
9. 然后任何一个偶数位的位置，都可以(最多）有两个 
10. 因为a[i]是偶数，那只要放在对角线就可以了