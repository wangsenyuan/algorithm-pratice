Carol is currently curling.

She has n disks each with radius r on the 2D plane.

Initially she has all these disks above the line y = 10100.

She then will slide the disks towards the line y = 0 one by one in order from 1 to n.

When she slides the i-th disk, she will place its center at the point (xi, 10100). She will then push it so the disk’s y
coordinate continuously decreases, and x coordinate stays constant. The disk stops once it touches the line y = 0 or it
touches any previous disk. Note that once a disk stops moving, it will not move again, even if hit by another disk.

Compute the y-coordinates of centers of all the disks after all disks have been pushed.

### ideas

1. 对于i, 需要知道 xi - r, xi + r
2. 假设对于前面的所有的circle计算出了它们的圆心位置 (x, y)
3. 且它们在i的接触范围内
4. 那么需要计算它的y[i]， 满足的条件是
5. dx * dx + dy * dy = 4 * r * r
6. dx是确定的， 然后可以计算出dy