class Solution {
public:
    int firstUniqChar(string s) {
        for(auto i : s){
            auto first=s.find_first_of(i);
            auto last=s.find_last_of(i);
            if(first!=string::npos&&first==last){
                return first;
            }
        }
        return -1;
    }
};