/*

863. All Nodes Distance K in Binary Tree

We are given a binary tree (with root node root), a target node, and an integer value K.

Return a list of the values of all nodes that have a distance K from the target node.  The answer can be returned in any order.

 

Example 1:

Input: root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, K = 2

Output: [7,4,1]

Explanation: 
The nodes that are a distance 2 from the target node (with value 5)
have values 7, 4, and 1.

Note that the inputs "root" and "target" are actually TreeNodes.
The descriptions of the inputs above are just serializations of these objects.
 

Note:

The given tree is non-empty.
Each node in the tree has unique values 0 <= node.val <= 500.
The target node is a node in the tree.
0 <= K <= 1000.


Test Cases:

[3,5,1,6,2,0,8,null,null,7,4]
5
2

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
    vector<int> distanceK(TreeNode* root, TreeNode* target, int K) {
        dfs1(root);
        dfs2(target->val, -1, 0, K);
        return ret_;
    }
    void dfs1(TreeNode* root) {
        if (!root) return;
        if (root->left) {
            dfs1(root->left);
            edge_[root->val].push_back(root->left->val);
            edge_[root->left->val].push_back(root->val);
        }
        if (root->right) {
            dfs1(root->right);
            edge_[root->val].push_back(root->right->val);
            edge_[root->right->val].push_back(root->val);
        }
    }
    void dfs2(int u, int fa, int depth, int K) {
        if (depth == K) {
            ret_.push_back(u); return;
        }
        for (int i = 0; i < edge_[u].size(); ++i) {
            int v = edge_[u][i];
            if (v == fa) continue;
            dfs2(v, u, depth+1, K);
        }
    }
public:
    vector<int> ret_;
    map<int, vector<int>> edge_;
};
