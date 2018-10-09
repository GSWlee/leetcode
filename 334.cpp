class Solution {
public:
    string reverseString(string s) {
        int len =s.length();
        string ans;
        for (int i=0;i<len;i++){
            ans.append(s,len-i-1,1);
        }
        return ans;
    }
};