class Solution {
public:
    bool find132pattern(vector<int>& nums) {
        int two=INT_MIN;
        stack<int> s;
        for(int i=nums.size()-1;i>=0;i--){
            if (nums[i]<two){
                return true;
            }else{
                while(!s.empty()&&nums[s.top()]<nums[i]){
                    two=nums[s.top()];
                    s.pop();
                }
                s.push(i);
            }
        }
        return false;
    }
};