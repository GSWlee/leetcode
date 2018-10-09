class Solution {
public:
    string reverseword(string s) {
        int len =s.length();
        string ans;
        for (int i=0;i<len;i++){
            ans.append(s,len-i-1,1);
        }
        return ans;
    }
    string reverseWords(string s) {
        int i=0,l=0;
        string ans;
        for(;l<s.length();l++){
            if (s[l]==' '){
                string temp(s,i,l-i);
                temp=this->reverseword(temp);
                ans.append(temp+" ");
                i=l+1;
            }
            else if(l==s.length()-1){
                string temp(s,i,l-i+1);
                temp=this->reverseword(temp);
                ans.append(temp);
            }
            
        }
        return ans;
    }
};