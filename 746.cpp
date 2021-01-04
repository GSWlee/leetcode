class Solution {
public:
    int minCostClimbingStairs(vector<int>& cost) {
        vector<int> ans(cost.size()+1,0);
        
        for(int i=2;i<cost.size()+1;i++){
            ans[i]=min(cost[i-1]+ans[i-1],cost[i-2]+ans[i-2]);
        }
        return ans[cost.size()];
    }
};