Filled with optimism, Hyunuk will host a conference about how great this new year will be!

The conference will have 𝑛
 lectures. Hyunuk has two candidate venues 𝑎
 and 𝑏
. For each of the 𝑛
 lectures, the speaker specified two time intervals [𝑠𝑎𝑖,𝑒𝑎𝑖]
 (𝑠𝑎𝑖≤𝑒𝑎𝑖
) and [𝑠𝑏𝑖,𝑒𝑏𝑖]
 (𝑠𝑏𝑖≤𝑒𝑏𝑖
). If the conference is situated in venue 𝑎
, the lecture will be held from 𝑠𝑎𝑖
 to 𝑒𝑎𝑖
, and if the conference is situated in venue 𝑏
, the lecture will be held from 𝑠𝑏𝑖
 to 𝑒𝑏𝑖
. Hyunuk will choose one of these venues and all lectures will be held at that venue.

Two lectures are said to overlap if they share any point in time in common. Formally, a lecture held in interval [𝑥,𝑦]
 overlaps with a lecture held in interval [𝑢,𝑣]
 if and only if max(𝑥,𝑢)≤min(𝑦,𝑣)
.

We say that a participant can attend a subset 𝑠
 of the lectures if the lectures in 𝑠
 do not pairwise overlap (i.e. no two lectures overlap). Note that the possibility of attending may depend on whether Hyunuk selected venue 𝑎
 or venue 𝑏
 to hold the conference.

A subset of lectures 𝑠
 is said to be venue-sensitive if, for one of the venues, the participant can attend 𝑠
, but for the other venue, the participant cannot attend 𝑠
.

A venue-sensitive set is problematic for a participant who is interested in attending the lectures in 𝑠
 because the participant cannot be sure whether the lecture times will overlap. Hyunuk will be happy if and only if there are no venue-sensitive sets. Determine whether Hyunuk will be happy.

 ### ideas
 1. 检查是否存在一个venue-sensitive的set？
 2. 假设有一个set，包括(i,j,k)几个lectures， 它们在a中是没有overlap的，但是在b中是overlap的，那么就是存在的
 3. 或者反过来，（i，j, k)在a中存在overlap，但是在b中不overlap
 4. 假设存在这样的一对(i, j)满足它在a中overlap，但是在b中不overlap，这个肯定是venue-sensitive的
 5. 但是假设不存在，是不是就一定没有呢？（不考虑b，a的情况）
 6. 不存在，也就是说如果在a中（i，j)重叠，但是在b中不重叠
 7. 先按照左端点升序排，对于j，那些仍然active的区域，在b上查找是否有和j重叠的就可以了