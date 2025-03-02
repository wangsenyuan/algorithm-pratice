The travel network is represented as 𝑛
 vertices numbered from 1
 to 𝑛
 and 𝑚
 edges. The 𝑖
-th edge connects vertices 𝑢𝑖
 and 𝑣𝑖
, and takes 𝑤𝑖
 seconds to travel (all 𝑤𝑖
 are even). Marian starts at vertex 1
 (Market) and Robin starts at vertex 𝑛
 (Major Oak).

In addition, ℎ
 of the 𝑛
 vertices each has a single horse available. Both Marian and Robin are capable riders, and could mount horses in no time (i.e. in 0
 seconds). Travel times are halved when riding. Once mounted, a horse lasts the remainder of the travel. Meeting must take place on a vertex (i.e. not on an edge). Either could choose to wait on any vertex.

Output the earliest time Robin and Marian can meet. If vertices 1
 and 𝑛
 are disconnected, output −1
 as the meeting is cancelled.