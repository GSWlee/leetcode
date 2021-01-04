class Solution {
public:
    vector<int> nextGreaterElement(vector<int>& nums1, vector<int>& nums2) {
        vector<int> ans;
        for(auto i : nums1){
            auto it=find(nums2.begin(),nums2.end(),i);
            if(it==nums2.end()){
                ans.push_back(-1);
            }else{
                auto f=false;
                for(;it<nums2.end();it++){
                    if(*it>i){
                        ans.push_back(*it);
                        f=true;
                        break;
                    }
                }
                if(!f){
                    ans.push_back(-1);
                }
            }
        }
        return ans;
    }
};