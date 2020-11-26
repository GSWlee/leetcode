class Solution {
public:
    int maximumGap(vector<int>& nums) {
        sort(nums.begin(),nums.end());
        int max=0;
        for(int i=1;i<nums.size();i++){
            if ((nums[i]-nums[i-1])>max){
                max=nums[i]-nums[i-1];
            }
        }
        return max;
    }
};