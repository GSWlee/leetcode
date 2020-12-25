class Solution {
public:
    void binaryTreePaths(TreeNode* root,vector<string> &ans,string temp){
        if (!root){
            return;
        }
        if (root->left==nullptr&&root->right==nullptr){
            temp=temp+to_string(root->val);
            ans.push_back(temp);
            return;
        }
        temp=temp+to_string(root->val)+"->";
        if(root->left!=nullptr){
            binaryTreePaths(root->left,ans,temp);
        }
        if(root->right!=nullptr){
            binaryTreePaths(root->right,ans,temp);
        }

    }
    vector<string> binaryTreePaths(TreeNode* root) {
        vector<string> ans;
        binaryTreePaths(root,ans,"");
        return ans;
    }
};