/*

777. Swap Adjacent in LR String

In a string composed of 'L', 'R', and 'X' characters, like "RXXLRXRXL", a move consists of either replacing one occurrence of "XL" with "LX", or replacing one occurrence of "RX" with "XR". Given the starting string start and the ending string end, return True if and only if there exists a sequence of moves to transform one string to the other.

Example:

Input: start = "RXXLRXRXL", end = "XRLXXRRLX"
Output: True
Explanation:
We can transform start to end following these steps:
RXXLRXRXL ->
XRXLRXRXL ->
XRLXRXRXL ->
XRLXXRRXL ->
XRLXXRRLX
Note:

1 <= len(start) = len(end) <= 10000.
Both start and end will only consist of characters in {'L', 'R', 'X'}.

*/


class Solution {
public:

    // passed all of test cases
    bool canTransform(string s, string e) {
        int n = s.size(), i = 0, j = 0;
        
        while (i < n && j < n) {
            while(i < n && s[i] == 'X') ++i;
            while(j < n && e[j] == 'X') ++j;
            if (s[i] != e[j]) return false;
            if ((s[i] == 'L' && i < j) || (s[i] == 'R' && i > j)) return false;
            ++i; ++j;
        }
        return true;
    }

    // 43 / 78 test cases passed
    bool canTransform1(string s, string e) {
        int n = s.size();
        int sx = 0, sr = 0, sl = 0;
        int ex = 0, er = 0, el = 0;
        for (auto start : s)
            switch(start) {
                case 'X': ++sx; break;
                case 'R': ++sr; break;
                case 'L': ++sl; break;
            }
        for (auto end : e)
            switch(end) {
                case 'X': ++ex; break;
                case 'R': ++er; break;
                case 'L': ++el; break;
            }
        if (sx != ex || sr != er || sl != el || sx == 0 || ex == 0)
            return false;
        
        int i, j;
        for (i = 0, j = 1; i < n; ++i, ++j) {
            if (i < n - 1) {
                if ((s[i] == 'R' && s[j] == 'X' && e[i] == 'X' && e[j] == 'R') || 
                    (s[i] == 'L' && s[j] == 'X' && e[i] == 'X' && e[j] == 'L') || 
                    (s[i] == 'X' && s[j] == 'R' && e[i] == 'R' && e[j] == 'X') || 
                    (s[i] == 'X' && s[j] == 'L' && e[i] == 'L' && e[j] == 'X') ||
                    (s[i] == 'X' && s[j] == 'L' && e[i] == 'X' && e[j] == 'L') || 
                    (s[i] == 'X' && s[j] == 'R' && e[i] == 'X' && e[j] == 'R') ) {
                    ++i; ++j;
                } else if (s[i] == e[i]) {
                    continue;
                } else {
                    return false;
                }
            } else {
                if (s[i] != e[i])
                    return false;
            }
        }
        
        return true;
    }
};
