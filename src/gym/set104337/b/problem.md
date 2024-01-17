Walk Alone likes to go with the flow, so he designs a function on digits about mode.

Let 𝑓(𝑥)
be the maximal occurrence among the digits in decimal expression of the number 𝑥
. For example, 𝑓(133)=2
since digit 3
occurs twice, while 𝑓(213)=𝑓(0)=1
since every digit appears exactly once in both numbers. Walk Alone gives you a task to calculate the range sum of
function 𝑓
, i.e. ∑𝑟𝑖=𝑙𝑓(𝑖)
.

### thoughts

1. 数位dp
2. 在给定上界时，计算当mode = x时的计数吗？
3. 0...9个数，每个数最多就20次
4. dp[i][x][f]表示前i个数，且x出现f次时的mode
5. 但这里似乎还需要一个状态就是mode = ？
6. 这个状态不大对， 因为新加入一个数字，比如2时，如何变更mode呢？
7. 如果2就是最多的那个数，那么mode + 1， 如果不是，那么增加f，但是如何知道最多的数是哪个呢？