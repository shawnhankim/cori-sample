/*

145. Binary Tree Postorder Traversal

Given a binary tree, return the postorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [3,2,1]
Follow up: Recursive solution is trivial, could you do it iteratively?

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
    vector<int> postorderTraversal(TreeNode* root) {
        vector<int> res;
        PostOrder(root, res);
        return res;
    }
    void PostOrder(TreeNode* root, vector<int>& res) {
        if (!root) return;
        PostOrder(root->left, res);
        PostOrder(root->right, res);
        res.push_back(root->val);
    }
};
