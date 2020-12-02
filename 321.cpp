vector<int> maxseq(vector<int> nums,int k){
    vector<int> ans;
    for(int i=nums.size()-1;i>=0;i--){
        if(ans.size()<k){
            ans.push_back(nums[i]);
        }else{
            int nowtag=k-1;
            int now=nums[i];
            while(nowtag>=0&&ans[nowtag]<=now){
                int temp=ans[nowtag];
                ans[nowtag]=now;
                now=temp;
                nowtag--;
            }
        }
    }
    reverse(ans.begin(),ans.end());
    return ans;
}
bool cmp(vector<int> a,int ta,vector<int> b,int tb){
    while(ta<a.size()&&tb<b.size()){
        if(a[ta]<b[tb]){
            return false;
        }else if (a[ta]>b[tb]){
            return true;
        }else{
            tb++;
            ta++;
        }
    }
    if(ta==a.size()){
        return false;
    }
    return true;
}
vector<int> merge(vector<int>a,vector<int> b){
    vector<int> ans;
    for(int ta=0,tb=0;ta<a.size()||tb<b.size();){
        if(ta<a.size()&&tb<b.size()){
            if(cmp(a,ta,b,tb)){
                ans.push_back(a[ta]);
                ta++;
            }else{
                ans.push_back(b[tb]);
                tb++;
            }
        }else{
            for(;ta<a.size();ta++){
                ans.push_back(a[ta]);
            }
            for(;tb<b.size();tb++){
                ans.push_back(b[tb]);
            }
        }
    }
    return ans;
}

bool compare(vector<int> a,vector<int> b){
    for(int i=0;i<a.size();i++){
        if(a[i]>b[i]){
            return true;
        }else if(a[i]<b[i]){
            return false;
        }
    }
    return true;
}
vector<int> maxNumber(vector<int>& nums1, vector<int>& nums2, int k) {
    vector<int> ans(k,0);
    for(int i=max(0,int(k-nums2.size()));i<=min(k,int(nums1.size()));i++){
        auto seq1=maxseq(nums1,i);
        auto seq2=maxseq(nums2,k-i);
        auto seq=merge(seq1,seq2);
        if(!compare(ans,seq)){
            ans=seq;
        }
    }
    return ans;
}