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
    TreeNode* sortedArrayToBST(vector<int>& nums) {
        sort(nums.begin(),nums.end());
        if(nums.empty()){
            return NULL;
        }else{
            TreeNode* temp =new TreeNode;
            temp->val=nums[nums.size()/2];
            if(nums.size()%2){
                vector<int> numsl(nums.begin(),nums.begin()+nums.size()/2);
                temp->left=sortedArrayToBST(numsl);
                vector<int> numsr(nums.begin()+nums.size()/2+1,nums.end());
                temp->right=sortedArrayToBST(numsr);
            }else{
                vector<int> numsl(nums.begin(),nums.begin()+nums.size()/2);
                temp->left=sortedArrayToBST(numsl);
                vector<int> numsr(nums.begin()+nums.size()/2+1,nums.end());
                temp->right=sortedArrayToBST(numsr);
            }
            return temp;
        }
    }
};