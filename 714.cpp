class Solution {
public:
    int maxProfit(vector<int>& prices, int fee) {
        vector<int> have(prices.size(),0);
        vector<int> no(prices.size(),0);
        have[0]=0-prices[0]-fee;
        for(int i=1;i<prices.size();i++){
            have[i]=max(have[i-1],no[i-1]-prices[i]-fee);
            no[i]=max(have[i-1]+prices[i],no[i-1]);
        }
        return max(have[prices.size()-1],no[prices.size()-1]);
    }
};