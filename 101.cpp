class Solution {
public:
    bool istrue(TreeNode* R,TreeNode*L){
        if((R==NULL)&&(L==NULL))
            return true;
        else if(R&&L){
            if(R->val==L->val){
                return istrue(R->right,L->left)&&istrue(R->left,L->right);
            }else{
                return false;
            }
        }else{
            return false;
        }
    }
    bool isSymmetric(TreeNode* root) {
        if(root==NULL)
            return true;
        else 
            return istrue(root->left,root->right);
    }
};
