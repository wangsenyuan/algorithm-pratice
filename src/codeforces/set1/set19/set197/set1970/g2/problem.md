In the medium version, 2≤∑𝑛≤300
 and 1≤∑𝑚≤300

After a worker's strike organized by the Dementors asking for equal rights, the prison of Azkaban has suffered some damage. After settling the spirits, the Ministry of Magic is looking to renovate the prison to ensure that the Dementors are kept in check. The prison consists of 𝑛
 prison cells and 𝑚
 bi-directional corridors. The 𝑖𝑡ℎ
 corridor is from cells 𝑢𝑖
 to 𝑣𝑖
. A subset of these cells 𝑆
 is called a complex if any cell in 𝑆
 is reachable from any other cell in 𝑆
. Formally, a subset of cells 𝑆
 is a complex if 𝑥
 and 𝑦
 are reachable from each other for all 𝑥,𝑦∈𝑆
, using only cells from 𝑆
 on the way. The funding required for a complex 𝑆
 consisting of 𝑘
 cells is defined as 𝑘2
.

As part of your Intro to Magical Interior Design course at Hogwarts, you have been tasked with designing the prison. The Ministry of Magic has asked that you divide the prison into 2
 complexes with 𝐞𝐱𝐚𝐜𝐭𝐥𝐲 𝐨𝐧𝐞 𝐜𝐨𝐫𝐫𝐢𝐝𝐨𝐫
 connecting them, so that the Dementors can't organize union meetings. For this purpose, you are allowed to build bi-directional corridors. The funding required to build a corridor between any 2
 cells is 𝑐
.

Due to budget cuts and the ongoing fight against the Death Eaters, you must find the 𝐦𝐢𝐧𝐢𝐦𝐮𝐦 𝐭𝐨𝐭𝐚𝐥 𝐟𝐮𝐧𝐝𝐢𝐧𝐠
 required to divide the prison as per the Ministry's requirements or −1
 if no division is possible.

Note: The total funding is the sum of the funding required for the 2
 complexes and the corridors built. If after the division, the two complexes have 𝑥
 and 𝑦
 cells respectively and you have built a total of 𝑎
 corridors, the total funding will be 𝑥2+𝑦2+𝑐×𝑎
. Note that 𝑥+𝑦=𝑛
.


### ideas
1. 现在需要增加新的边了，原来已经连在一起的，可以不用增加边，但有可能需要切分
2. 要增加的边的数量是确定的，就是把这些分离的component连在一起的数量，假设有x个component，那么 a = x - 1
3. 然后就是看怎么连，可以让 x和y最接近
4. 那些强连通的，必须分在一组（因为不能满足，两个complex通过一条边连通的要求）
5. 这些可以变成一个点。剩余的就是一条线，或者一个点，或者一棵树
6. 还是有一些问题。首先，我们希望x和y足够的接近，这样子 x**2 + y ** y 才最小
7. 所以，在哪里分隔是个大问题。假设分割边的两个端点点是(u, v)
8. 那么(u, v)可能把它所在的边分成了两半(su, sv), 然后其他的部分，可以连接到u或者v上
9. 所以要知道其他部分的和，可能的组合数 
10. 其他边，是sum dp的方式