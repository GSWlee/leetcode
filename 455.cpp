class Solution {
public:
    int findContentChildren(vector<int>& g, vector<int>& s) {
        sort(g.begin(),g.end());
        sort(s.begin(),s.end());
        int ans=0;
        for(auto i : g){
            bool find=false;
            for(auto j=s.begin();j!=s.end();j++){
                if(*j>=i){
                    ans++;
                    find=true;
                    s.erase(j);
                    break;
                }
            }
            if(!find){
                return ans;
            }
        }
        return ans;
    }
};