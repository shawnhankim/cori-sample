/*

452. Minimum Number of Arrows to Burst Balloons

Medium

There are a number of spherical balloons spread in two-dimensional space. For each balloon, provided input is the start and end coordinates of the horizontal diameter. Since it's horizontal, y-coordinates don't matter and hence the x-coordinates of start and end of the diameter suffice. Start is always smaller than end. There will be at most 104 balloons.

An arrow can be shot up exactly vertically from different points along the x-axis. A balloon with xstart and xend bursts by an arrow shot at x if xstart ≤ x ≤ xend. There is no limit to the number of arrows that can be shot. An arrow once shot keeps travelling up infinitely. The problem is to find the minimum number of arrows that must be shot to burst all balloons.

Example:

Input:
[[10,16], [2,8], [1,6], [7,12]]

Output:
2

Explanation:
One way is to shoot one arrow for example at x = 6 (bursting the balloons [2,8] and [1,6]) and another arrow at x = 11 (bursting the other two balloons).

*/

func findMinArrowShots(points [][]int) int {
    var res int
    arrow := math.MaxInt32
    
    sort.Sort(ByLen(points))
    for _, pt := range points {
        if arrow != math.MaxInt32 && pt[0] <= arrow {
            continue
        }
        arrow = pt[1]
        res += 1
    }
    return res
}

type ByLen [][]int

func (a ByLen) Len() int {
    return len(a)
}

func (a ByLen) Less(i, j int) bool {
    return a[i][1] < a[j][1]
}

func (a ByLen) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}


