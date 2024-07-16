To satisfy his love of matching socks, Phoenix has brought his 𝑛
 socks (𝑛
 is even) to the sock store. Each of his socks has a color 𝑐𝑖
 and is either a left sock or right sock.

Phoenix can pay one dollar to the sock store to either:

recolor a sock to any color 𝑐′
 (1≤𝑐′≤𝑛)
turn a left sock into a right sock
turn a right sock into a left sock
The sock store may perform each of these changes any number of times. Note that the color of a left sock doesn't change when it turns into a right sock, and vice versa.

A matching pair of socks is a left and right sock with the same color. What is the minimum cost for Phoenix to make 𝑛/2
 matching pairs? Each sock must be included in exactly one matching pair. 


 ### ideas
 1. 那些已经matching的，是不是不用考虑了？
 2. 假设i, j是一对，也就是它们是相反的，且有相同的颜色
 3. 那么在最后的答案中，是否存在把其中的一个从左变成右？
 4. 假设需要，那么可以得到一个结论是，右 < 左（否则就不需要转换）假设这时，还有一只同样颜色的，左袜子c，那么用c转换就没有区别。但是如果c的颜色不一样，还是可以得到用c转换（颜色+左右）没有区别；
 5. 那么也就是说，一开始把所有能匹配的，都匹配掉。然后剩下不能匹配的
 6. 要么颜色不一样，要么极性一致，要么两个同时存在
 7. 所以，先把袜子分成两类，左/右
 8. 多出来的，肯定要转换， 然后转换的时候，应该尽量使用颜色多的那种
 9. 假设要转换x个，然后对于每种颜色，计算左边比右边多的那部分（颜色的差值），选择一半的（不超过x）转换到另外一边