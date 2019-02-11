/*

298. Binary Tree Longest Consecutive Sequence

Given a binary tree, find the length of the longest consecutive sequence path.

The path refers to any sequence of nodes from some starting node to any node in the tree along the parent-child connections. The longest consecutive path need to be from parent to child (cannot be the reverse).

Example 1:

Input:

   1
    \
     3
    / \
   2   4
        \
         5

Output: 3

Explanation: Longest consecutive sequence path is 3-4-5, so return 3.
Example 2:

Input:

   2
    \
     3
    / 
   2    
  / 
 1

Output: 2 

Explanation: Longest consecutive sequence path is 2-3, not 3-2-1, so return 2.

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
    int longestConsecutive(TreeNode* root) {
        if (!root)
            return 0;
        int res = 1, tmp = 1;
        lc(root, res, tmp);
        return res;
    }
    void lc(TreeNode *root, int &res, int tmp) {
        if (!root) {
            return;
        }
        int ltmp = tmp, rtmp = tmp, val = 0;
        if (root->left) {
            val = (root->val - root->left->val);
            if (val == -1) {
                ltmp += 1;
            } else {
                ltmp = 1;
            }
            res = max(res, ltmp);
            lc(root->left, res, ltmp);
        }
        if (root->right) {
            val = (root->val - root->right->val);
            if (val == -1) {
                rtmp += 1;
            } else {
                rtmp = 1;
            }
            res = max(res, rtmp);
            lc(root->right, res, rtmp);
        }
    }
};


class Solution {
public:
    int longestConsecutive(TreeNode* root) {
        int res = 0;
        lc(root, nullptr, res, 1);
        return res;
    }
    void lc(TreeNode* root, TreeNode* par, int& res, int len) {
        if (!root) return;
        len = (par && (root->val - par->val) == 1) ? len + 1 : 1;
        res = max(res, len);
        lc(root->left , root, res, len);
        lc(root->right, root, res, len);
    }
};


