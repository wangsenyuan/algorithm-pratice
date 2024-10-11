As you might already know, space has always been a problem in ICPC Jakarta. To cope with this, ICPC Jakarta is planning to build two new buildings. These buildings should have a shape of a rectangle of the same size. Now, their problem is to find land to build the buildings.

There are 𝑁
 lands available for sale. The 𝑖𝑡ℎ
 land has a rectangular shape of size 𝐿𝑖×𝑊𝑖
. For a good feng shui, the building's side should be parallel to the land's sides.

One way is to build the two buildings on two different lands, one on each land (not necessarily with the same orientation). A building of size 𝐴×𝐵
 can be build on the 𝑖𝑡ℎ
 land if and only if at least one of the following is satisfied:

𝐴≤𝐿𝑖
 and 𝐵≤𝑊𝑖
, or
𝐴≤𝑊𝑖
 and 𝐵≤𝐿𝑖
.
Alternatively, it is also possible to build two buildings of 𝐴×𝐵
 on the 𝑖𝑡ℎ
 land with the same orientation. Formally, it is possible to build two buildings of 𝐴×𝐵
 on the 𝑖𝑡ℎ
 land if and only if at least one of the following is satisfied:
𝐴×2≤𝐿𝑖
 and 𝐵≤𝑊𝑖
, or
𝐴×2≤𝑊𝑖
 and 𝐵≤𝐿𝑖
, or
𝐴≤𝐿𝑖
 and 𝐵×2≤𝑊𝑖
, or
𝐴≤𝑊𝑖
 and 𝐵×2≤𝐿𝑖
.
Your task in this problem is to help ICPC Jakarta to figure out the largest possible buildings they can build given 𝑁
 available lands. Note that ICPC Jakarta has to build two buildings of 𝐴×𝐵
; output the largest possible for 𝐴×𝐵
.

