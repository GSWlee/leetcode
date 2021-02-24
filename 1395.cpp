class Solution {
public:
    int numTeams(vector<int>& rating) {
        int ans=0,n=rating.size();
        for(int i=1;i<n-1;i++){
            int ll=0,lm=0,rl=0,rm=0;
            for(int j=0;j<i;j++){
                if(rating[j]<rating[i]){
                    ll++;
                }else if(rating[j]>rating[i]){
                    lm++;
                }
            }
            for(int j=i+1;j<n;j++){
                if(rating[j]<rating[i]){
                    rl++;
                }else if(rating[j]>rating[i]){
                    rm++;
                }
            }
            ans+=ll*rm+lm*rl;
        }
        return ans;
    }
};