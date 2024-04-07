A bus moves along the coordinate line Ox from the point x = 0 to the point x = a. After starting from the point x = 0,
it reaches the point x = a, immediately turns back and then moves to the point x = 0. After returning to the point x = 0
it immediately goes back to the point x = a and so on. Thus, the bus moves from x = 0 to x = a and back. Moving from the
point x = 0 to x = a or from the point x = a to x = 0 is called a bus journey. In total, the bus must make k journeys.

The petrol tank of the bus can hold b liters of gasoline. To pass a single unit of distance the bus needs to spend
exactly one liter of gasoline. The bus starts its first journey with a full petrol tank.

There is a gas station in point x = f. This point is between points x = 0 and x = a. There are no other gas stations on
the bus route. While passing by a gas station in either direction the bus can stop and completely refuel its tank. Thus,
after stopping to refuel the tank will contain b liters of gasoline.

What is the minimum number of times the bus needs to refuel at the point x = f to make k journeys? The first journey
starts in the point x = 0.

### ideas

1. 先考虑第一次加油时的时机，假设行驶了距离x <= b,
2. 如果 let cnt = b / a, rem = b % a
3. 如果 rem >= a - f， 且 cnt % 2 = 1
4. 那么第一次加油时，bus是向左的，否则只能是在向右的时候
5. 如果 rem >= f 且 cnt % 2 == 0,
6. 那么此时可以计算处在第一次加油前跑了几趟
7. 然后就是从f到f，看一箱油能跑几趟