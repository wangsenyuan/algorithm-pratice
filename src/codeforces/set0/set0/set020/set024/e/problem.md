Recently the construction of Berland collider has been completed. Collider can be represented as a long narrow tunnel that contains n particles. We associate with collider 1-dimensional coordinate system, going from left to right. For each particle we know its coordinate and velocity at the moment of start of the collider. The velocities of the particles don't change after the launch of the collider. Berland scientists think that the big bang will happen at the first collision of particles, whose velocities differs in directions. Help them to determine how much time elapses after the launch of the collider before the big bang happens.

Input
The first line contains single integer n (1 ≤ n ≤ 5·105) — amount of particles in the collider. Next n lines contain description of particles. Each particle is described by two integers xi, vi ( - 109 ≤ xi, vi ≤ 109, vi ≠ 0) — coordinate and velocity respectively. All the coordinates are distinct. The particles are listed in order of increasing of coordinates. All the coordinates are in meters, and all the velocities — in meters per second. The negative velocity means that after the start of collider the particle will move to the left, and the positive — that the particle will move to the right.

Output
If there will be no big bang, output -1. Otherwise output one number — how much time in seconds elapses after the launch of the collider before the big bang happens. Your answer must have a relative or absolute error less than 10 - 9.


### ideas
1. t = (s1 - s2) / (v1 + v2) (v1和v2都是绝对值，且符号相反)
2. 固定t的情况下， t >= (s1 - s2) / (v1 + v2) 时，这两个部分就可以相遇
3. v1 * t + v2 * t >= s1 - s2
4. s1 - v1 * t <= s2 + v2 * t
5. 