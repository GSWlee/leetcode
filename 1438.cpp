class Solution {
public:
    int longestSubarray(vector<int>& nums, int limit) {
        int longest=0;
        if(nums.empty()){
            return 0;
        }
        int max=nums[0],min=nums[0],maxp=0,minp=0,temp=0;
        for(int i=0;i<nums.size();i++){
            if(abs(nums[i]-max)<=limit&&abs(nums[i]-min)<=limit){
                temp++;
                if(max<=nums[i]){
                    max=nums[i];
                    maxp=i;
                }
                if(min>=nums[i]){
                    min=nums[i];
                    minp=i;
                }
            }else{
                longest=temp>longest?temp:longest;
                if(abs(nums[i]-max)>limit){
                    i=maxp;
                    max=nums[maxp+1];
                    min=nums[maxp+1];
                    maxp=maxp+1;
                    minp=maxp+1;
                    temp=0;
                }else{
                    i=minp;
                    max=nums[minp+1];
                    min=nums[minp+1];
                    maxp=minp+1;
                    minp=minp+1;
                    temp=0;
                }
            }
        }
        longest=temp>longest?temp:longest;
        return longest;
    }
};