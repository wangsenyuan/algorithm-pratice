Theofanis really likes to play with the bits of numbers. He has an array 𝑎
of size 𝑛
and an integer 𝑘
. He can make at most 𝑘
operations in the array. In each operation, he picks a single element and increases it by 1
.

He found the maximum bitwise AND that array 𝑎
can have after at most 𝑘
operations.

Theofanis has put a lot of work into finding this value and was very happy with his result. Unfortunately, Adaś, being
the evil person that he is, decided to bully him by repeatedly changing the value of 𝑘
.

Help Theofanis by calculating the maximum possible bitwise AND for 𝑞
different values of 𝑘
. Note that queries are independent.

### solution

Let 𝑆=∑(220−𝑎𝑖)
. If 𝑘≥𝑆
then the answer is 220+⌊𝑘−𝑆𝑛⌋
.

Similarly to D1 let's construct the answer bit by bit. Let 𝑥
be the current answer and 𝑏
be the bit we want to add.

Let's look at the amount of operations we need to do on the 𝑖
-th element to change our answer from 𝑥
to 𝑥+2𝑏
.

if 𝑥
is not a submask of 𝑎𝑖
, then after constructing answer 𝑥
it has 0
s on all bits not greater than 𝑏
. In this case we need to increase the 𝑖
-th element by 2𝑏
.
if 𝑥+2𝑏
is a submask of 𝑎𝑖
, then we do not need to increase the 𝑖
-th element.
otherwise we need to increase the 𝑖
-th element by 2𝑏−𝑎𝑖
𝑚𝑜𝑑
2𝑏
.
We can handle all three cases efficiently if we precompute the following two arrays:

𝐴(𝑚𝑎𝑠𝑘)
– how many elements from the array is 𝑚𝑎𝑠𝑘
a submask of.

𝐵(𝑚𝑎𝑠𝑘,𝑏)
– sum of 𝑎𝑖
𝑚𝑜𝑑
2𝑏
over all 𝑎𝑖
for which 𝑥
is a submask.

Both arrays can be calculated efficiently using SOS dp.

This allows us to answer the queries in 𝑂
(𝑞×𝑙𝑜𝑔
𝑎
) with 𝑂
(𝑎
𝑙𝑜𝑔2
𝑎
) preprocessing