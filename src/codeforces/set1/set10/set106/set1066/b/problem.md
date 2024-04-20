Vova's house is an array consisting of 𝑛
elements (yeah, this is the first problem, I think, where someone lives in the array). There are heaters in some
positions of the array. The 𝑖
-th element of the array is 1
if there is a heater in the position 𝑖
, otherwise the 𝑖
-th element of the array is 0
.

Each heater has a value 𝑟
(𝑟
is the same for all heaters). This value means that the heater at the position 𝑝𝑜𝑠
can warm up all the elements in range [𝑝𝑜𝑠−𝑟+1;𝑝𝑜𝑠+𝑟−1]
.

Vova likes to walk through his house while he thinks, and he hates cold positions of his house. Vova wants to switch
some of his heaters on in such a way that each element of his house will be warmed up by at least one heater.

Vova's target is to warm up the whole house (all the elements of the array), i.e. if 𝑛=6
, 𝑟=2
and heaters are at positions 2
and 5
, then Vova can warm up the whole house if he switches all the heaters in the house on (then the first 3
elements will be warmed up by the first heater and the last 3
elements will be warmed up by the second heater).

Initially, all the heaters are off.

But from the other hand, Vova didn't like to pay much for the electricity. So he wants to switch the minimum number of
heaters on in such a way that each element of his house is warmed up by at least one heater.

Your task is to find this number of heaters or say that it is impossible to warm up the whole house.