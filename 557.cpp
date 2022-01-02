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

class Solution {
public:
    string reverseWords(string s) {
        int start=0;
        for(int j =0;j<s.length();j++){
            if(s[j]==' '){
                int end =j-1;
                while(start<end){
                    auto temp=s[start];
                    s[start]=s[end];
                    s[end]=temp;
                    start++;
                    end--;
                }
                start=j+1;
            }else{
                continue;
            }
        }
        int end =s.length()-1;
        while(start<end){
            auto temp=s[start];
            s[start]=s[end];
            s[end]=temp;
            start++;
            end--;
        }
        return s;
    }
};