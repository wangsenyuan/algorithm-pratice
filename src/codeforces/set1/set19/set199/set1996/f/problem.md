Sparkle gives you two arrays a
 and b
 of length n
. Initially, your score is 0
. In one operation, you can choose an integer i
 and add ai
 to your score. Then, you must set ai
 = max(0,ai−bi)
.

You only have time to perform k
 operations before Sparkle sets off a nuclear bomb! What is the maximum score you can acquire after k
 operations?

 ### ideas
 1. ai一旦变成0了，就没法操作了
 2. a -> a - b -> a - 2 * b -> ...
 3. 假设操作了x次后 a - x * b <= 0, x = (a + b - 1) / b
 4. 在最后一次前， a - (x - 1) * b > 0
 5. 这里是不是模拟就可以了？每次都选择最大的a[i], 将它减少到第二大a[j]小？
 6. 这样子，有可能会tle。
 7. 按照a派个序，假设后面的所有的数字，都只能处理到 比a[i]小，计算需要多少次操作
 8. 如果这个操作次数>k, 说明还需要找更大的数字
 9. 如果这个操作次数 <= k, 说明这个下限太大了，可以往前找更小的数字
 10. 也就是说在左边的i 需要的次数超过 k, 在右边的i，需要的次数少于等于k
 11. 找到这样一个i，先使用足够的次数，将右边的都变成 < a[i]的数字（有可能变成0，但是在操作前，都是 >= a[i])的
 12. 如果还剩余一些次数，这些可以只能用来处理a[i]及后面的部分。 再处理一次？