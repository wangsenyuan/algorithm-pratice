Bizon the Champion isn't just attentive, he also is very hardworking.

Bizon the Champion decided to paint his old fence his favorite color, orange. The fence is represented as n vertical
planks, put in a row. Adjacent planks have no gap between them. The planks are numbered from the left to the right
starting from one, the i-th plank has the width of 1 meter and the height of ai meters.

Bizon the Champion bought a brush in the shop, the brush's width is 1 meter. He can make vertical and horizontal strokes
with the brush. During a stroke the brush's full surface must touch the fence at all the time (see the samples for the
better understanding). What minimum number of strokes should Bizon the Champion do to fully paint the fence? Note that
you are allowed to paint the same area of the fence multiple times.

### ideas

1. n <= 5000, 可以考虑n**2的算法
2. 答案不会差于n,只使用竖直方向的，最多只需要n次
3. 可以考虑这样一个情况，如果在i前面不管怎么处理，需要dp[i]次操作
4. dp[i+1]怎么计算？那么dp[i] = dp[i-1] + 1
5. 如果a[i] <= a[i-1], 那么 dp[i] = dp[j] + a[i]如果a[j+1...i]  >= a[i]
6. 还有一种情况是 dp[j] + min(a[j+1...i]) + a[i] + 1