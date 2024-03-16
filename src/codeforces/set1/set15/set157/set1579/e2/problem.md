You are given an integer array 𝑎[1…𝑛]=[𝑎1,𝑎2,…,𝑎𝑛]
.

Let us consider an empty deque (double-ended queue). A deque is a data structure that supports adding elements to both
the beginning and the end. So, if there are elements [3,4,4]
currently in the deque, adding an element 1
to the beginning will produce the sequence [1,3,4,4]
, and adding the same element to the end will produce [3,4,4,1]
.

The elements of the array are sequentially added to the initially empty deque, starting with 𝑎1
and finishing with 𝑎𝑛
. Before adding each element to the deque, you may choose whether to add it to the beginning or to the end.

For example, if we consider an array 𝑎=[3,7,5,5]
, one of the possible sequences of actions looks like this:

1. add 3
   to the beginning of the deque:    deque has a sequence [3]
   in it;
2. add 7
   to the end of the deque:    deque has a sequence [3,7]
   in it;
3. add 5
   to the end of the deque:    deque has a sequence [3,7,5]
   in it;
4. add 5
   to the beginning of the deque:    deque has a sequence [5,3,7,5]
   in it;
   Find the minimal possible number of inversions in the deque after the whole array is processed.

An inversion in sequence 𝑑
is a pair of indices (𝑖,𝑗)
such that 𝑖<𝑗
and 𝑑𝑖>𝑑𝑗
. For example, the array 𝑑=[5,3,7,5]
has exactly two inversions — (1,2)
and (3,4)
, since 𝑑1=5>3=𝑑2
and 𝑑3=7>5=𝑑4
.

### ideas

1. 还没有抓住题眼～
2. 对于最后一个数来说，它前面的要们在它的后面（它放在最前面）或者在它的前面（它放在最后面）
3. 无论哪种情况来说，已经是确定的
4. 可以直接放入对应的位置即可
5. 对于第二个数来说，前面的情况依然适用，
6. 但是最后一个数的位置，它是控制不了的（如果已经占了第一个，它没得选）
7. 所以，后面的数，对前面数的位置没有影响，只需要考虑前面数，对后面数的影响