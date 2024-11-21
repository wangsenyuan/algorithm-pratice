You are an all-powerful being and you have created a rectangular world. In fact, your world is so bland that it could be represented by a 𝑟×𝑐
 grid. Each cell on the grid represents a country. Each country has a dominant religion. There are only two religions in your world. One of the religions is called Beingawesomeism, who do good for the sake of being good. The other religion is called Pushingittoofarism, who do murders for the sake of being bad.

Oh, and you are actually not really all-powerful. You just have one power, which you can use infinitely many times! Your power involves missionary groups. When a missionary group of a certain country, say 𝑎
, passes by another country 𝑏
, they change the dominant religion of country 𝑏
 to the dominant religion of country 𝑎
.

In particular, a single use of your power is this:

You choose a horizontal 1×𝑥
 subgrid or a vertical 𝑥×1
 subgrid. That value of 𝑥
 is up to you;
You choose a direction 𝑑
. If you chose a horizontal subgrid, your choices will either be NORTH or SOUTH. If you choose a vertical subgrid, your choices will either be EAST or WEST;
You choose the number 𝑠
 of steps;
You command each country in the subgrid to send a missionary group that will travel 𝑠
 steps towards direction 𝑑
. In each step, they will visit (and in effect convert the dominant religion of) all 𝑠
 countries they pass through, as detailed above.
The parameters 𝑥
, 𝑑
, 𝑠
 must be chosen in such a way that any of the missionary groups won't leave the grid.
The following image illustrates one possible single usage of your power. Here, A represents a country with dominant religion Beingawesomeism and P represents a country with dominant religion Pushingittoofarism. Here, we've chosen a 1×4
 subgrid, the direction NORTH, and 𝑠=2
 steps.


You are a being which believes in free will, for the most part. However, you just really want to stop receiving murders that are attributed to your name. Hence, you decide to use your powers and try to make Beingawesomeism the dominant religion in every country.

What is the minimum number of usages of your power needed to convert everyone to Beingawesomeism?

With god, nothing is impossible. But maybe you're not god? If it is impossible to make Beingawesomeism the dominant religion in all countries, you must also admit your mortality and say so.

### ideas
1. 显然，如果全部是P => mortal
2. 如果至少有一个A，那么最多是4次。如果A在最中间，选择x = 1, s = i - 1, north
3. 然后选择x = 1, s = n - i, sorth
4. 然后选择x = n, s = j - 1, west
5. 然后选择x = n, s = m - j, east
6. 如果在任何一个角落有A，那么这个答案可以到2
7. 如果有一个边是全A，那么就可以到1
8. 如果全部是A，那么就是0
9. 有没有3的可能？也有，就是在某个边上的中间