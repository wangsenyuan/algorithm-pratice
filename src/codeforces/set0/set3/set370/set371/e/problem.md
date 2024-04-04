Berland is going through tough times — the dirt price has dropped and that is a blow to the country's economy. Everybody
knows that Berland is the top world dirt exporter!

The President of Berland was forced to leave only k of the currently existing n subway stations.

The subway stations are located on a straight line one after another, the trains consecutively visit the stations as
they move. You can assume that the stations are on the Ox axis, the i-th station is at point with coordinate xi. In such
case the distance between stations i and j is calculated by a simple formula |xi - xj|.

Currently, the Ministry of Transport is choosing which stations to close and which ones to leave. Obviously, the
residents of the capital won't be too enthusiastic about the innovation, so it was decided to show the best side to the
people. The Ministry of Transport wants to choose such k stations that minimize the average commute time in the subway!

Assuming that the train speed is constant (it is a fixed value), the average commute time in the subway is calculated as
the sum of pairwise distances between stations, divided by the number of pairs (that is ) and divided by the speed of
the train.

Help the Minister of Transport to solve this difficult problem. Write a program that, given the location of the stations
selects such k stations that the average commute time in the subway is minimized.

### ideas

1. sort and get the min sum 