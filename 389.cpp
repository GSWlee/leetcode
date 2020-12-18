class Solution {
public:
    char findTheDifference(string s, string t) {
        for(auto a : t){
            auto temp=s.find(a);
            if(temp==s.npos){
                return a;
            }else{
                s.erase(temp,1);
            }
        }
        return 0;
    }
};