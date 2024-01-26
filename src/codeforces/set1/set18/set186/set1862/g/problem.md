Tema bought an old device with a small screen and a worn-out inscription "The Great Equalizer" on the side.

The seller said that the device needs to be given an array 𝑎
of integers as input, after which "The Great Equalizer" will work as follows:

Sort the current array in non-decreasing order and remove duplicate elements leaving only one occurrence of each
element.
If the current length of the array is equal to 1
, the device stops working and outputs the single number in the array — output value of the device.
Add an arithmetic progression {𝑛, 𝑛−1, 𝑛−2, …, 1
} to the current array, where 𝑛
is the length of the current array. In other words, 𝑛−𝑖
is added to the 𝑖
-th element of the array, when indexed from zero.
Go to the first step.
To test the operation of the device, Tema came up with a certain array of integers 𝑎
, and then wanted to perform 𝑞
operations on the array 𝑎
of the following type:

Assign the value 𝑥
(1≤𝑥≤109
) to the element 𝑎𝑖
(1≤𝑖≤𝑛
).
Give the array 𝑎
as input to the device and find out the result of the device's operation, while the array 𝑎
remains unchanged during the operation of the device.
Help Tema find out the output values of the device after each operation.

### thoughts

1. 假设在最后一次操作后，元素的个数变成了1，在此之前的数组长度为n
2. 那么必然有 a[0] + n = a[1] +n - 1 = ... = a[n-1] + 1
3. 进而可以得到，a是一个相差为1的等差数列
4. 还有一个观察是，a只能增大，而不能变小
5. 还有一个观察就是，相邻两个数至少差1，但每次操作后，它们的差最多缩小1
6. 所以，数字前后的关系是不变的，换句话说，每次操作都是在变更最大值和最小值的差
7. 直到把它们变成0，如果进行了x轮，那么最后的结果就是 最大值 + x
8. 但是这个x怎么计算出来？
9. ai + n0 + n1 + n2 + ... + nk = ax + x
10. n0 + n1 + n2 + ... + nk = ax + x - ai
11. 假设一开始按照相邻数的差，假设最大的是10，那么10次后，它们肯定会变成相同的