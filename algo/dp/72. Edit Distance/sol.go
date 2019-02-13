/*

72. Edit Distance

Hard

Given two words word1 and word2, find the minimum number of operations required to convert word1 to word2.

You have the following 3 operations permitted on a word:

Insert a character
Delete a character
Replace a character
Example 1:

Input: word1 = "horse", word2 = "ros"
Output: 3
Explanation: 
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e')
Example 2:

Input: word1 = "intention", word2 = "execution"
Output: 5
Explanation: 
intention -> inention (remove 't')
inention -> enention (replace 'i' with 'e')
enention -> exention (replace 'n' with 'x')
exention -> exection (replace 'n' with 'c')
exection -> execution (insert 'u')

*/

func minDistance(word1 string, word2 string) int {
    m  := int(len(word1))
    n  := int(len(word2))
    dp := make([][]int, n+1)
        
    for i := 0; i <= n; i++ {
        dp[i] = make([]int, m+1)
        dp[i][0] = i
    }
    for j := 0; j <= m; j++ {
        dp[0][j] = j
    }
    for i := 1; i <= n; i++ {
        for j := 1; j <= m; j++ {
            if word1[j-1] == word2[i-1] {
                dp[i][j] = 1 + min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1]-1))
            } else {
                dp[i][j] = 1 + min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1]))
            }
        }
    }
    return dp[n][m]
}

func min(a int, b int) int {
    if a > b {
        return b
    }
    return a
}

