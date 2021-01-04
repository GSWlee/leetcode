/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    int findlast(TreeNode* root){
        while(root->left!=nullptr){
            root=root->left;
        }
        return root->val;
    }
    TreeNode* deleteNode(TreeNode* root, int key) {
        if(!root){
            return root;
        }else{
            if(root->val==key){
                if(root->right&&root->left){
                    auto t=findlast(root->right);
                    root->val=t;
                    root->right=deleteNode(root->right,t);
                }else if(!root->right&&!root->left){
                    return nullptr;
                }else if(root->right&&!root->left){
                    return root->right;
                }else{
                    return root->left;
                }
            }else if(root->val>key){
                root->left=deleteNode(root->left,key);

            }else{
                root->right=deleteNode(root->right,key);
            }
        }
        return root;
    }
};