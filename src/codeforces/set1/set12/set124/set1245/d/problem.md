Shichikuji is the new resident deity of the South Black Snail Temple. Her first job is as follows:

There are 𝑛
 new cities located in Prefecture X. Cities are numbered from 1
 to 𝑛
. City 𝑖
 is located 𝑥𝑖
 km North of the shrine and 𝑦𝑖
 km East of the shrine. It is possible that (𝑥𝑖,𝑦𝑖)=(𝑥𝑗,𝑦𝑗)
 even when 𝑖≠𝑗
.

Shichikuji must provide electricity to each city either by building a power station in that city, or by making a connection between that city and another one that already has electricity. So the City has electricity if it has a power station in it or it is connected to a City which has electricity by a direct connection or via a chain of connections.

Building a power station in City 𝑖
 will cost 𝑐𝑖
 yen;
Making a connection between City 𝑖
 and City 𝑗
 will cost 𝑘𝑖+𝑘𝑗
 yen per km of wire used for the connection. However, wires can only go the cardinal directions (North, South, East, West). Wires can cross each other. Each wire must have both of its endpoints in some cities. If City 𝑖
 and City 𝑗
 are connected by a wire, the wire will go through any shortest path from City 𝑖
 to City 𝑗
. Thus, the length of the wire if City 𝑖
 and City 𝑗
 are connected is |𝑥𝑖−𝑥𝑗|+|𝑦𝑖−𝑦𝑗|
 km.
Shichikuji wants to do this job spending as little money as possible, since according to her, there isn't really anything else in the world other than money. However, she died when she was only in fifth grade so she is not smart enough for this. And thus, the new resident deity asks for your help.

And so, you have to provide Shichikuji with the following information: minimum amount of yen needed to provide electricity to all cities, the cities in which power stations will be built, and the connections to be made.

If there are multiple ways to choose the cities and the connections to obtain the construction of minimum price, then print any of them.


### ideas
1. 把一组city连接起来后，只需要在cost 最低的城市建造一个station就可以了