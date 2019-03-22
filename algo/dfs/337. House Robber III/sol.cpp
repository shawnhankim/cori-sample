/*

337. House Robber III

Medium

The thief has found himself a new place for his thievery again. There is only one entrance to this area, called the "root." Besides the root, each house has one and only one parent house. After a tour, the smart thief realized that "all houses in this place forms a binary tree". It will automatically contact the police if two directly-linked houses were broken into on the same night.

Determine the maximum amount of money the thief can rob tonight without alerting the police.

Example 1:

Input: [3,2,3,null,3,null,1]

     3
    / \
   2   3
    \   \ 
     3   1

Output: 7 
Explanation: Maximum amount of money the thief can rob = 3 + 3 + 1 = 7.
Example 2:

Input: [3,4,5,1,3,null,1]

     3
    / \
   4   5
  / \   \ 
 1   3   1

Output: 9
Explanation: Maximum amount of money the thief can rob = 4 + 5 = 9.


*/

/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    int rob(TreeNode* root) {
        int l = 0, r = 0;
        return Rob(root, l, r);
    }
    int Rob(TreeNode* root, int& l, int& r) {
        if (!root) return 0;
        int ll = 0, lr = 0, rl = 0, rr = 0;
        l = Rob(root->left , ll, lr);
        r = Rob(root->right, rl, rr);
        return max(root->val + ll + lr + rl + rr, l + r);
    }
    int rob1(TreeNode* root) {
        int sodd = 0, seven = 0;
        r(root, sodd, seven, 0);
        return max(sodd, seven);
    }
    void r(TreeNode* root, int& sodd, int& seven, int depth) {
        if (!root) return;
        if (depth & 1) sodd  += root->val;
        else           seven += root->val;
        r(root->left , sodd, seven, depth+1);
        r(root->right, sodd, seven, depth+1);
    }
};
