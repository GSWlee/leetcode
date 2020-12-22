class Solution {
public:
    vector<vector<int>> zigzagLevelOrder(TreeNode* root) {
        vector<vector<int>> ans;
        vector<TreeNode*> last;
        if (root!=nullptr){
            last.push_back(root);
        }
        int n=0;
        while(!last.empty()){
            vector<TreeNode*> temp;
            vector<int> nums;
            for(auto i : last){
                if(i->left!=nullptr){
                    temp.push_back(i->left);
                }
                if(i->right!=nullptr){
                    temp.push_back(i->right);
                }
                nums.push_back(i->val);
            }
            if(n%2){
                reverse(nums.begin(),nums.end());
            }
            n++;
            ans.push_back(nums);
            last=temp;
        }
        return ans;
    }
};