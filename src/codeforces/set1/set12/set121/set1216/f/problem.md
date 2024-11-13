You work as a system administrator in a dormitory, which has 𝑛
 rooms one after another along a straight hallway. Rooms are numbered from 1
 to 𝑛
.

You have to connect all 𝑛
 rooms to the Internet.

You can connect each room to the Internet directly, the cost of such connection for the 𝑖
-th room is 𝑖
 coins.

Some rooms also have a spot for a router. The cost of placing a router in the 𝑖
-th room is also 𝑖
 coins. You cannot place a router in a room which does not have a spot for it. When you place a router in the room 𝑖
, you connect all rooms with the numbers from 𝑚𝑎𝑥(1, 𝑖−𝑘)
 to 𝑚𝑖𝑛(𝑛, 𝑖+𝑘)
 inclusive to the Internet, where 𝑘
 is the range of router. The value of 𝑘
 is the same for all routers.

Calculate the minimum total cost of connecting all 𝑛
 rooms to the Internet. You can assume that the number of rooms which have a spot for a router is not greater than the number of routers you have.

 ### ideas
 1. 常规动态规划考虑一下
 2. dp[i] = dp[i-1] + i 如果用room i处直接连接
 3. dp[i] = (?) + dp[l] 如果能够在i....l中间放置一个router
 4.       =  如果在 i - k 处有一个router 那么 l = max(0, i - 2 * k - 1)
 5.       =  i - k + dp[i - 2 * k - 1]
 6. 如果在i处放置了一个router，那么前面k个位置的dp也会被影响到
 7. 有个感觉是，如果可以放置router的地方，始终应该放置router，
 8. 除非这个router被cover住了？似乎也不对， 0110 (k = 1)
 9. 显然应该放置一个router在后面
 10. 从后往前，记录上一个没有被连接的位置，j，如果当前有个router可以放置，那么就使用router，替换j
 11. 如果有两个router呢？0110 （k = 2)时，最好使用前面的
 12. 有个办法，就是记录目前可以用的router，如果上一个连接的能够cover住当前的，就push进去。
 13. 否则就用那个最能cover住当前位置的router
 14. 但是问题在于，有可能使用后面的一个，也能cover住当前的，且还是个更好的选择