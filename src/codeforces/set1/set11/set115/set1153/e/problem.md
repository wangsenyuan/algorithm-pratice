Now Serval is a senior high school student in Japari Middle School. However, on the way to the school, he must go across a pond, in which there is a dangerous snake. The pond can be represented as a 𝑛×𝑛
 grid. The snake has a head and a tail in different cells, and its body is a series of adjacent cells connecting the head and the tail without self-intersecting. If Serval hits its head or tail, the snake will bite him and he will die.

Luckily, he has a special device which can answer the following question: you can pick a rectangle, it will tell you the number of times one needs to cross the border of the rectangle walking cell by cell along the snake from the head to the tail. The pictures below show a possible snake and a possible query to it, which will get an answer of 4
.



Today Serval got up too late and only have time to make 2019
 queries. As his best friend, can you help him find the positions of the head and the tail?

Note that two cells are adjacent if and only if they have a common edge in the grid, and a snake can have a body of length 0
, that means it only has adjacent head and tail.

Also note that the snake is sleeping, so it won't move while Serval using his device. And what's obvious is that the snake position does not depend on your queries.

### ideas
1. 如果一个区域和这条蛇的交互区域是偶数 > 0, 那么头、尾肯定在区域外, 如果 = 0, 那么就不好说了
2. 如果为奇数，那么肯定一个在区域外，一个在区域内
3. 