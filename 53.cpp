class Solution {
public:
    int maxSubArray(vector<int>& nums,int start,int end) {
        if(start==end)
            return nums[start];
        if(start==end-1)
            return nums[start];
        int mid=(start+end)/2;
        int left_max=this->maxSubArray(nums,start,mid);
        int right_max=this->maxSubArray(nums,mid,end);
        int sum=0;
        int left_mid_max=INT_MIN;
        for (int i=mid-1;i>=start;i--){
            sum+=nums[i];
            if(sum>left_mid_max)
                left_mid_max=sum;
        }
        sum=0;
        int right_mid_max=INT_MIN;
        for (int i=mid;i<end;i++){
            sum+=nums[i];
            if(sum>right_mid_max)
                right_mid_max=sum;
        }
        int max=left_mid_max+right_mid_max;
        if(left_max>max)
            max=left_max;
        if(right_max>max)
            max=right_max;
        return max;
    }
    int maxSubArray(vector<int>& nums) {
        return this->maxSubArray(nums,0,nums.size());
    }
};