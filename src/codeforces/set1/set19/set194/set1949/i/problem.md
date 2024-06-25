You are given 𝑛
 disks in the plane. The center of each disk has integer coordinates, and the radius of each disk is a positive integer. No two disks overlap in a region of positive area, but it is possible for disks to be tangent to each other.

Your task is to determine whether it is possible to change the radii of the disks in such a way that:

Disks that were tangent to each other remain tangent to each other.
No two disks overlap in a region of positive area.
The sum of all radii strictly decreases.
The new radii are allowed to be arbitrary positive real numbers. The centers of the disks cannot be changed.

### ideas
1. 圆形的面积= pi * r * r， pi是常量，所以面积的改变，之和r有有关
2. 如果是两个圆相切在一起，那么一个的面积变小delta，另外一个就要增加delta，所以总面积不会变
3. 但是如果是3个，依次在一起，那么就可以通过增大中间的，减小两边的，从而减小总面积
4. 但是如果这3个是两辆相切，或者组成了一个环，那么是不能改变半径的
5. 组成线型的可以改变。
6. 组成树形的呢？也是可以改变的。因为树形，可以二分成两种颜色的，如果红色多一个，那么红色的变小，蓝色的变大即可
7. 所以，进而推广，就是是否可以对图进行二分，如果可以，就可以变化，看个数