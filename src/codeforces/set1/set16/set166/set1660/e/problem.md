You are given a binary matrix 𝐴
of size 𝑛×𝑛
. Rows are numbered from top to bottom from 1
to 𝑛
, columns are numbered from left to right from 1
to 𝑛
. The element located at the intersection of row 𝑖
and column 𝑗
is called 𝐴𝑖𝑗
. Consider a set of 4
operations:

Cyclically shift all rows up. The row with index 𝑖
will be written in place of the row 𝑖−1
(2≤𝑖≤𝑛
), the row with index 1
will be written in place of the row 𝑛
.
Cyclically shift all rows down. The row with index 𝑖
will be written in place of the row 𝑖+1
(1≤𝑖≤𝑛−1
), the row with index 𝑛
will be written in place of the row 1
.
Cyclically shift all columns to the left. The column with index 𝑗
will be written in place of the column 𝑗−1
(2≤𝑗≤𝑛
), the column with index 1
will be written in place of the column 𝑛
.
Cyclically shift all columns to the right. The column with index 𝑗
will be written in place of the column 𝑗+1
(1≤𝑗≤𝑛−1
), the column with index 𝑛
will be written in place of the column 1
.
The 3×3
matrix is shown on the left before the 3
-rd operation is applied to it, on the right — after.
You can perform an arbitrary (possibly zero) number of operations on the matrix; the operations can be performed in any
order.

After that, you can perform an arbitrary (possibly zero) number of new xor-operations:

Select any element 𝐴𝑖𝑗
and assign it with new value 𝐴𝑖𝑗⊕1
. In other words, the value of (𝐴𝑖𝑗+1)mod2
will have to be written into element 𝐴𝑖𝑗
.
Each application of this xor-operation costs one burl. Note that the 4
shift operations — are free. These 4
operations can only be performed before xor-operations are performed.

Output the minimum number of burles you would have to pay to make the 𝐴
matrix unitary. A unitary matrix is a matrix with ones on the main diagonal and the rest of its elements are zeros (that
is, 𝐴𝑖𝑗=1
if 𝑖=𝑗
and 𝐴𝑖𝑗=0
otherwise).

### ideas

1. 同一行/列的元素貌似不会改变，但是位置可以变
2. 假设有一整行都没有1，但是它的下一行有两个1，那无论怎么处理，这一行都不会出现1
3. 