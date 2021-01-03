class Solution {
public:
    int pivotIndex(vector<int>& nums) {
        int sum=0;
        for(auto i : nums){
            sum+=i;
        }
        int temp=0;
        for(int i=0;i<nums.size();i++){
            if(2*temp==(sum-nums[i])){
                return i;
            }else{
                temp+=nums[i];
            }
        }
        return -1;
    }
};