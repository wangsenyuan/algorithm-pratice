The Berland Forest can be represented as an infinite cell plane. Every cell contains a tree. That is, contained before the recent events.

A destructive fire raged through the Forest, and several trees were damaged by it. Precisely speaking, you have a 𝑛×𝑚
 rectangle map which represents the damaged part of the Forest. The damaged trees were marked as "X" while the remaining ones were marked as ".". You are sure that all burnt trees are shown on the map. All the trees outside the map are undamaged.

The firemen quickly extinguished the fire, and now they are investigating the cause of it. The main version is that there was an arson: at some moment of time (let's consider it as 0
) some trees were set on fire. At the beginning of minute 0
, only the trees that were set on fire initially were burning. At the end of each minute, the fire spread from every burning tree to each of 8
 neighboring trees. At the beginning of minute 𝑇
, the fire was extinguished.

The firemen want to find the arsonists as quickly as possible. The problem is, they know neither the value of 𝑇
 (how long the fire has been raging) nor the coordinates of the trees that were initially set on fire. They want you to find the maximum value of 𝑇
 (to know how far could the arsonists escape) and a possible set of trees that could be initially set on fire.

Note that you'd like to maximize value 𝑇
 but the set of trees can be arbitrary.

 ### ideas
 1. 对于给定的T，相当于用2 * T- 1的正方形去覆盖所有着火的区域，不能多也不能少