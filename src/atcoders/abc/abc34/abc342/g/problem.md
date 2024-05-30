You are given an integer sequence 
A=(A 
1
​
 ,A 
2
​
 ,…,A 
N
​
 ) of length 
N.

Process 
Q operations in order. There are three types of operations:

A type-1 operation is represented by a triple of integers 
(l,r,x), which corresponds to replacing 
A 
i
​
  with 
max{A 
i
​
 ,x} for each 
i=l,l+1,…,r.
A type-2 operation is represented by an integer 
i, which corresponds to canceling the 
i-th operation (it is guaranteed that the 
i-th operation is of type 1 and has not already been canceled). The new sequence 
A can be obtained by performing all type-1 operations that have not been canceled, starting from the initial state.
A type-3 operation is represented by an integer 
i, which corresponds to printing the current value of 
A 
i
​
 .
Constraints
1≤N≤2×10 
5
 
1≤A 
i
​
 ≤10 
9
  (1≤i≤N)
1≤Q≤2×10 
5
 
In a type-1 operation, 
1≤l≤r≤N and 
1≤x≤10 
9
 .
In a type-2 operation, 
i is not greater than the number of operations given before, and 
1≤i.
In a type-2 operation, the 
i-th operation is of type 1.
In type-2 operations, the same 
i does not appear multiple times.
In a type-3 operation, 
1≤i≤N.
All input values are integers.


### ideas
1. 如果没有取消操作，是可以用segment tree的
2. 有点没办法呐～
3. 假设对于位置pos，如果它在操作i的覆盖范围内，
4. 如果它后面有新的type1的操作覆盖到它，那么取消i对它的值没有影响
5. 否则取消i后，pos的值要回退到i之前的状态
6. 假设，对于每个操作i，记录了它执行前的这一段的状态，那么取消操作i，就返回回去？这个不行，特别是上面情况1很难处理
7. 看看，能不能保存历史，通过快速的查询来实现
8. 不考虑取消的情况，多次操作，对一个位置的影响，相当于给它赋值它们的最大值
9. 而且操作的顺序也无关


### solution

First, we consider applying and canceling 
A←max{A,x} against a single integer, instead of a sequence.

Since the operations commute, one can manage the multiset of the values of 
x for the operations not canceled so far in a heap (so-called “erasable priority queue” or a balanced binary search tree) to solve it fast.

This idea can be adopted to the original problem too.

As in a segment tree, consider 
O(N) segments contained in 
[0,N) and written in the form of 
[i2 
k
 ,(i+1)2 
k
 ) (i,k are integers
). Then, for any segment 
[l,r),

updating 
A 
i
​
 ←max{A 
i
​
 ,x} and
canceling it
can be realized as addition and removal of 
x against 
O(logN) multisets.

The response to a retrieval query of 
A 
i
​
  can also be found by inspecting the maximum values of the 
O(logN) multisets containing 
A 
i
​
 , and then finding the maximum among them.

The time complexity is 
O((N+Q)logNlogQ), using appropriate data structures.

