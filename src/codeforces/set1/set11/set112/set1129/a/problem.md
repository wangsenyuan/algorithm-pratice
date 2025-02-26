Alice received a set of Toy Train™ from Bob. It consists of one train and a connected railway network of 𝑛
 stations, enumerated from 1
 through 𝑛
. The train occupies one station at a time and travels around the network of stations in a circular manner. More precisely, the immediate station that the train will visit after station 𝑖
 is station 𝑖+1
 if 1≤𝑖<𝑛
 or station 1
 if 𝑖=𝑛
. It takes the train 1
 second to travel to its next station as described.

Bob gave Alice a fun task before he left: to deliver 𝑚
 candies that are initially at some stations to their independent destinations using the train. The candies are enumerated from 1
 through 𝑚
. Candy 𝑖
 (1≤𝑖≤𝑚
), now at station 𝑎𝑖
, should be delivered to station 𝑏𝑖
 (𝑎𝑖≠𝑏𝑖
).

The blue numbers on the candies correspond to 𝑏𝑖
 values. The image corresponds to the 1
-st example.
The train has infinite capacity, and it is possible to load off any number of candies at a station. However, only at most one candy can be loaded from a station onto the train before it leaves the station. You can choose any candy at this station. The time it takes to move the candies is negligible.

Now, Alice wonders how much time is needed for the train to deliver all candies. Your task is to find, for each station, the minimum time the train would need to deliver all the candies were it to start from there.