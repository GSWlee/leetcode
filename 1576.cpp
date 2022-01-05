class Solution {
public:
    string modifyString(string s) {
        for(int i=0;i<s.size();i++){
            if (s[i]!='?'){
                continue;
            }else{
                if(i==0){
                    if (s[i+1]!='a'){
                        s[i]='a';
                    }else{
                        s[i]='b';
                    }
                }else if( i==s.size()-1){
                    if(s[i-1]!='a'){
                        s[i]='a';
                    }else{
                        s[i]='b';
                    }
                }else{
                    auto temp = (s[i-1]-'a'+1)%26+'a';
                    if (s[i+1]!=temp){
                        s[i]=temp;
                    }else{
                        temp = (s[i-1]-'a'+2)%26+'a';
                        s[i]=temp;
                    }
                }
            }
        }
        return s;
    }
};