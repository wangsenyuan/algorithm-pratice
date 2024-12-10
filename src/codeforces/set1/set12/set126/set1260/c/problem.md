You are a rebel leader and you are planning to start a revolution in your country. But the evil Government found out about your plans and set your punishment in the form of correctional labor.

You must paint a fence which consists of 10100
 planks in two colors in the following way (suppose planks are numbered from left to right from 0
):

if the index of the plank is divisible by 𝑟
 (such planks have indices 0
, 𝑟
, 2𝑟
 and so on) then you must paint it red;
if the index of the plank is divisible by 𝑏
 (such planks have indices 0
, 𝑏
, 2𝑏
 and so on) then you must paint it blue;
if the index is divisible both by 𝑟
 and 𝑏
 you can choose the color to paint the plank;
otherwise, you don't need to paint the plank at all (and it is forbidden to spent paint on it).
Furthermore, the Government added one additional restriction to make your punishment worse. Let's list all painted planks of the fence in ascending order: if there are 𝑘
 consecutive planks with the same color in this list, then the Government will state that you failed the labor and execute you immediately. If you don't paint the fence according to the four aforementioned conditions, you will also be executed.

The question is: will you be able to accomplish the labor (the time is not important) or the execution is unavoidable and you need to escape at all costs.


### ideas
1. let c = lcm(b, r)
2. 在x % c的部分，可以交替paint，所以没有问题
3. 然后就是c前面的部分，就是看是否会出现连续k个的red/blue
