After a successful field test, Heidi is considering deploying a trap along some Corridor, possibly not the first one. She wants to avoid meeting the Daleks inside the Time Vortex, so for abundance of caution she considers placing the traps only along those Corridors that are not going to be used according to the current Daleks' plan – which is to use a minimum spanning tree of Corridors. Heidi knows that all energy requirements for different Corridors are now different, and that the Daleks have a single unique plan which they are intending to use.

Your task is to calculate the number 𝐸𝑚𝑎𝑥(𝑐)
, which is defined in the same way as in the easy version – i.e., the largest 𝑒≤109
 such that if we changed the energy of corridor 𝑐
 to 𝑒
, the Daleks might use it – but now for every corridor that Heidi considers.

### ideas
1. 没想法了～
2. 所有的edge边上的weight是不同的，所以mst是唯一的
3. 需要处理的是那些不会被用到的边，通过改变它们的weight，使的它们有可能被用到（而不是必须被用到）
4. lcp... 这条路径上的最大值