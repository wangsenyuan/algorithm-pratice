### description

There is a square matrix, consisting of 𝑛
rows and 𝑛
columns of cells, both numbered from 1
to 𝑛
. The cells are colored white or black. Cells from 1
to 𝑎𝑖
are black, and cells from 𝑎𝑖+1
to 𝑛
are white, in the 𝑖
-th column.

You want to place 𝑚
integers in the matrix, from 1
to 𝑚
. There are two rules:

each cell should contain at most one integer;
black cells should not contain integers.
The beauty of the matrix is the number of such 𝑗
that 𝑗+1
is written in the same row, in the next column as 𝑗
(in the neighbouring cell to the right).

What's the maximum possible beauty of the matrix?

### solution

Notice that the rows of the matrix are basically independent. When we fill the matrix with integers, the values from the
rows are just added together. Moreover, in a single row, the segments of white cells separated by black cells are also
independent in the same way.

And how to solve the problem for one segment of white cells? Obviously, put the numbers one after another. Placing a
non-zero amount of integers 𝑘
will yield the result of 𝑘−1
. Thus, the answer is ∑𝑖=1𝑠𝑘𝑖−1
, where 𝑠
is the number of used segments. We know that the sum of 𝑘𝑖
is just 𝑚
. So it becomes 𝑚+∑𝑖=1𝑠−1=𝑚−𝑠
. So, the problem asks us to minimize the number of used segments.

In order to use the smallest number of segments, we should pick the longest ones. We only have to find them.

In the worst case, there are 𝑂(𝑛2)
segments of white cells in the matrix. However, we can store them in a compressed form. Since their lengths are from 1
to 𝑛
, we can calculate how many segments of each length there are.

Look at the matrix from the bottom to the top. First, there are some fully white rows. Then some black cell appears and
splits the segment of length 𝑛
into two segments. Maybe more if there are more black cells in that row. After that the split segments behave
independently of each other.

Let's record these intervals of rows each segment exists at. So, let some segment exist from row 𝑙
to row 𝑟
. What does that mean? This segment appeared because of some split at row 𝑟
. At row 𝑙
, some black cell appeared that split the segment in parts. If we knew the 𝑟
for the segment, we could've saved the information about it during the event at row 𝑙
.

So, we need a data structure that can support the following operations:

find a segment that covers cell 𝑥
;
erase a segment;
insert a segment.
Additionally, that data structure should store a value associated with a segment. Thus, let's use a map. For a segment,
which is a pair of integers, store another integer — the bottommost row this segment exists.

Make events (𝑎𝑖,𝑖)
for each column 𝑖
. Sort them and process in a non-increasing order. During the processing of the event, find the segment this black cell
splits. Save the information about this segment. Then remove it and add two new segments (or less if some are empty):
the one to the left and the one to the right.

At the end, for each length from 1
to 𝑛
, we will have the number of segments of such length. In order to fill them with 𝑚
integers, start with the longest segments and use as many of each length as possible. So, while having 𝑚
more integers to place and cnt
segments of length len
, you can use min(⌊𝑚len⌋,𝑐𝑛𝑡)
of segments.

Overall complexity: 𝑂(𝑛log𝑛)
per testcase.
