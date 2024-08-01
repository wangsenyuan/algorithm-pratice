ikrpprpp found an array 𝑎
 consisting of integers. He likes justice, so he wants to make 𝑎
 fair — that is, make it non-decreasing. To do that, he can perform an act of justice on an index 1≤𝑖≤𝑛
 of the array, which will replace 𝑎𝑖
 with 𝑎2𝑖
 (the element at position 𝑖
 with its square). For example, if 𝑎=[2,4,3,3,5,3]
 and ikrpprpp chooses to perform an act of justice on 𝑖=4
, 𝑎
 becomes [2,4,3,9,5,3]
.

What is the minimum number of acts of justice needed to make the array non-decreasing?

### ideas
1. 对a[i]的操作，不会影响i左边的情况，但会影响右边的。所以应该尽可能的让a[i]小，当然前提是比a[i-1]大
2. 这里的问题，应该是对a[i]操作几次后，会造成这个数，非常大
3. 所以需要一种表示，且能够比较两个数的大小
4. 2, 4, 8, 16, ....
5. 要用到对数