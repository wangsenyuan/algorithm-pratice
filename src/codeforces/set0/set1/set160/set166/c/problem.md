A median in an array with the length of n is an element which occupies position number after we sort the elements in the
non-decreasing order (the array elements are numbered starting with 1). A median of an array (2, 6, 1, 2, 3) is the
number 2, and a median of array (0, 96, 17, 23) — the number 17.

We define an expression as the integer part of dividing number a by number b.

One day Vasya showed Petya an array consisting of n integers and suggested finding the array's median. Petya didn't even
look at the array and said that it equals x. Petya is a very honest boy, so he decided to add several numbers to the
given array so that the median of the resulting array would be equal to x.

Petya can add any integers from 1 to 105 to the array, including the same numbers. Of course, he can add nothing to the
array. If a number is added multiple times, then we should consider it the number of times it occurs. It is not allowed
to delete of change initial numbers of the array.

While Petya is busy distracting Vasya, your task is to find the minimum number of elements he will need.

### ideas

1. sort it of course
2. 至多增加n个x
3. 需要知道有多少个a = 大于x, b = 小于x, c= 等于x的部分
4. 让x为中位数，需要 b + c >= (a + b + c) / 2, (a + c) >= (a + b + c) / 2