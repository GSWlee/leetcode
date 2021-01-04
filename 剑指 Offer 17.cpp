class Solution {
public:
    vector<int> printNumbers(int n) {
        auto temp=pow(10,n);
        vector<int> ans;
        for(int i=1;i<temp;i++){
            ans.push_back(i);
        }
        return ans;
    }
};