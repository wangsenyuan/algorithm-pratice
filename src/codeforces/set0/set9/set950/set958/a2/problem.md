https://codeforces.com/problemset/problem/958/A2

The stardate is 1983, and Princess Heidi is getting better at detecting the Death Stars. This time, two Rebel spies have yet again given Heidi two maps with the possible locations of the Death Star. Since she got rid of all double agents last time, she knows that both maps are correct, and indeed show the map of the solar system that contains the Death Star. However, this time the Empire has hidden the Death Star very well, and Heidi needs to find a place that appears on both maps in order to detect the Death Star.

The first map is an N × M grid, each cell of which shows some type of cosmic object that is present in the corresponding quadrant of space. The second map is an M × N grid. Heidi needs to align those two maps in such a way that they overlap over some M × M section in which all cosmic objects are identical. Help Heidi by identifying where such an M × M section lies within both maps.

### ideas
1. 在第一个map中，计算所有m * m 的hash
2. 