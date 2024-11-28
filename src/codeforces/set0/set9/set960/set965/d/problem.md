A lot of frogs want to cross a river. A river is 𝑤
 units width, but frogs can only jump 𝑙
 units long, where 𝑙<𝑤
. Frogs can also jump on lengths shorter than 𝑙
. but can't jump longer. Hopefully, there are some stones in the river to help them.

The stones are located at integer distances from the banks. There are 𝑎𝑖
 stones at the distance of 𝑖
 units from the bank the frogs are currently at. Each stone can only be used once by one frog, after that it drowns in the water.

What is the maximum number of frogs that can cross the river, given that then can only jump on the stones?

### ideas
1. 从后面开始，如果在w - 1 处，有a[w-1]个frog到达了这里
2. 那么这些frog肯定（如果 l >= 1)肯定能完成
3. 同样的= a[w-l]处
4. 但是 w - l - 1 处的，最远只能跳到w-1处，
5. 如果 a[w-l-1] <= a[w-1], 那么没有问题，全部可以跳过去，但是如果 a[w-l-1] > a[w-1]
6. 也不是完全不能，部分，可以跳到某个a[i]处
7. 所以这里出现一个贪心的策略
8. 从bank开始，如果尽量的填满接下来能填满的位置
9. 然后接下来的，也尽量去填满。依次类推，就可以了。
10. 