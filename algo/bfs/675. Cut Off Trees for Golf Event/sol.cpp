/*

675. Cut Off Trees for Golf Event

Hard

You are asked to cut off trees in a forest for a golf event. The forest is represented as a non-negative 2D map, in this map:

0 represents the obstacle can't be reached.
1 represents the ground can be walked through.
The place with number bigger than 1 represents a tree can be walked through, and this positive number represents the tree's height.
 

You are asked to cut off all the trees in this forest in the order of tree's height - always cut off the tree with lowest height first. And after cutting, the original place has the tree will become a grass (value 1).

You will start from the point (0, 0) and you should output the minimum steps you need to walk to cut off all the trees. If you can't cut off all the trees, output -1 in that situation.

You are guaranteed that no two trees have the same height and there is at least one tree needs to be cut off.

Example 1:

Input: 
[
 [1,2,3],
 [0,0,4],
 [7,6,5]
]
Output: 6
 

Example 2:

Input: 
[
 [1,2,3],
 [0,0,0],
 [7,6,5]
]
Output: -1
 

Example 3:

Input: 
[
 [2,3,4],
 [0,0,5],
 [8,7,6]
]
Output: 6
Explanation: You started from the point (0,0) and you can cut off the tree in (0,0) directly without walking.
 

Hint: size of the given matrix will not exceed 50x50

*/

// Wrong answer
class Solution {
public:
    int cutOffTree(vector<vector<int>>& forest) {
        int dx[4] = {1, 0, -1,  0};
        int dy[4] = {0, 1,  0, -1};
        int m = forest.size(), n = forest[0].size();
        int cnt0 = 0, cnt1 = 0, x, y, nx, ny;
        
        queue<int> qx, qy;
        qx.push(0); qy.push(0);
        
        while(!qx.empty()) {
            x = qx.front(); qx.pop();
            y = qy.front(); qy.pop();
            
            for (int i = 0; i < 4; ++i) {
                nx = x + dx[i]; if (nx < 0 || nx >= m) continue;
                ny = y + dy[i]; if (ny < 0 || ny >= n) continue;
                if (forest[nx][ny] == 0) { ++cnt0; continue; }
                if (forest[x][y] < forest[nx][ny]) {
                    ++cnt1; qx.push(nx); qy.push(ny);
                }
            }
        }
        printf("cnt0 : %d, cnt1 : %d, m: %d, n: %d\n", cnt0, cnt1, m, n);
        if (cnt0 + cnt1 == (m*n))
            return cnt1;
        return -1;
    }
};



