Heidi found out that the Daleks have created a network of bidirectional Time Corridors connecting different destinations (at different times!). She suspects that they are planning another invasion on the entire Space and Time. In order to counter the invasion, she plans to deploy a trap in the Time Vortex, along a carefully chosen Time Corridor. She knows that tinkering with the Time Vortex is dangerous, so she consulted the Doctor on how to proceed. She has learned the following:

Different Time Corridors require different amounts of energy to keep stable.
Daleks are unlikely to use all corridors in their invasion. They will pick a set of Corridors that requires the smallest total energy to maintain, yet still makes (time) travel possible between any two destinations (for those in the know: they will use a minimum spanning tree).
Setting the trap may modify the energy required to keep the Corridor stable.
Heidi decided to carry out a field test and deploy one trap, placing it along the first Corridor. But she needs to know whether the Daleks are going to use this corridor after the deployment of the trap.

She gives you a map of Time Corridors (an undirected graph) with energy requirements for each Corridor.

For a Corridor 𝑐
, 𝐸𝑚𝑎𝑥(𝑐)
 is the largest 𝑒≤109
 such that if we changed the required amount of energy of 𝑐
 to 𝑒
, then the Daleks may still be using 𝑐
 in their invasion (that is, it belongs to some minimum spanning tree). Your task is to calculate 𝐸𝑚𝑎𝑥(𝑐1)
 for the Corridor 𝑐1
 that Heidi plans to arm with a trap, which is the first edge in the graph.

### ideas
1. binary search?