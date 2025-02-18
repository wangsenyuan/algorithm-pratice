Let's imagine the surface of Mars as an infinite coordinate plane. Initially, the rover Perseverance-2 and the helicopter Ingenuity-2 are located at the point with coordinates (0,0)
. A set of instructions 𝑠
 consisting of 𝑛
 instructions of the following types was specially developed for them:

N: move one meter north (from point (𝑥,𝑦)
 to (𝑥,𝑦+1)
);
S: move one meter south (from point (𝑥,𝑦)
 to (𝑥,𝑦−1)
);
E: move one meter east (from point (𝑥,𝑦)
 to (𝑥+1,𝑦)
);
W: move one meter west (from point (𝑥,𝑦)
 to (𝑥−1,𝑦)
).
Each instruction must be executed either by the rover or by the helicopter. Moreover, each device must execute at least one instruction. Your task is to distribute the instructions in such a way that after executing all 𝑛
 instructions, the helicopter and the rover end up at the same point, or determine that this is impossible.

