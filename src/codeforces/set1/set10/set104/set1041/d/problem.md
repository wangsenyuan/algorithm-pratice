A plane is flying at a constant height of ℎ
meters above the ground surface. Let's consider that it is flying from the point (−109,ℎ)
to the point (109,ℎ)
parallel with 𝑂𝑥
axis.

A glider is inside the plane, ready to start his flight at any moment (for the sake of simplicity let's consider that he
may start only when the plane's coordinates are integers). After jumping from the plane, he will fly in the same
direction as the plane, parallel to 𝑂𝑥
axis, covering a unit of distance every second. Naturally, he will also descend; thus his second coordinate will
decrease by one unit every second.

There are ascending air flows on certain segments, each such segment is characterized by two numbers 𝑥1
and 𝑥2
(𝑥1<𝑥2
) representing its endpoints. No two segments share any common points. When the glider is inside one of such segments,
he doesn't descend, so his second coordinate stays the same each second. The glider still flies along 𝑂𝑥
axis, covering one unit of distance every second.

If the glider jumps out at 1
, he will stop at 10
. Otherwise, if he jumps out at 2
, he will stop at 12
.
Determine the maximum distance along 𝑂𝑥
axis from the point where the glider's flight starts to the point where his flight ends if the glider can choose any
integer coordinate to jump from the plane and start his flight. After touching the ground the glider stops altogether,
so he cannot glide through an ascending airflow segment if his second coordinate is 0
.

### ideas

1. start at one of the left point may produce the best result
2. glider 下落的距离是定的 = h, 在没有上升气流的情况下，他可以向右边移动的距离就是h，所以就是要在右边有h空的地方即可
3. 可以用双指针