class Solution {
public:
    int countNodes(TreeNode* root) {
        int right=0,left=0;
        if (root==nullptr){
            return 0;
        }
        if (root->left!=nullptr){
            left=countNodes(root->left);
        }
        if (root->right!=nullptr){
            right=countNodes(root->right);
        }
        return right+left+1;
    }
};