As Jaehyun's house is not large enough to have a bookshelf, he keeps the n books by stacking them vertically. When he
wants to read a certain book x, he follows the steps described below.

He lifts all the books above book x.
He pushes book x out of the stack.
He puts down the lifted books without changing their order.
After reading book x, he puts book x on the top of the stack.

He decided to read books for m days. In the j-th (1 ≤ j ≤ m) day, he will read the book that is numbered with integer
bj (1 ≤ bj ≤ n). To read the book, he has to use the process described in the paragraph above. It is possible that he
decides to re-read the same book several times.

After making this plan, he realized that the total weight of books he should lift during m days would be too heavy. So,
he decided to change the order of the stacked books before the New Year comes, and minimize the total weight. You may
assume that books can be stacked in any possible order. Note that book that he is going to read on certain step isn't
considered as lifted on that step. Can you help him?

### thoughts

1. 是对书架进行排序，不是对读书的顺序进行重拍
2. b[i] = x， 那么在阅读到b[i]时，它前面的书籍都被阅读了一次，所以此时不论怎么排，它们都是在x的上面，所以书架里面数书的位置，就是它第一次出现的位置
3. 然后就是怎么计算的问题了