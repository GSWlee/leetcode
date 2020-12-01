class Solution {
public:
    vector<int> searchRange(vector<int>& nums, int target) {
        bool find=false;
        vector<int> ans(2,-1);
        int length=nums.size();
        if (length==0){
            return ans;
        }
        int start=0,end=length-1;
        while(start+1<end){
            if(nums[(end+start)/2]>target){
                end=(end+start)/2;
            }else if(nums[(end+start)/2]==target){
                start=(end+start)/2;
                end=start;
                break;
            }else{
                start=(end+start)/2;
            }
        }
        if(nums[start]!=target&&nums[end]!=target){
            return ans;
        }
        if(nums[start]==target){
            end=start;
        }else{
            start=end;
        }
        while(start>=0&&nums[start]==target){
            ans[0]=start;
            start--;
        }
        while(end<length&&nums[end]==target&&end<length){
            ans[1]=end;
            end++;
        }
        return ans;
    }
};