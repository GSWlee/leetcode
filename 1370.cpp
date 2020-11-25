class Solution {
public:
    string sortString(string s) {
        sort(s.begin(),s.end());
        string ans="";
        bool flag=true;
        while(s.length()){
            if (flag){
                ans.append(string(1,s[0]));
                s.erase(0,1);
                for (int i=0;i<s.length();i++){
                    if (s[i]>ans[ans.length()-1]){
                        ans.append(string(1,s[i]));
                        s.erase(i,1);
                        i--;
                    }
                }
                flag=false;
            }else{
                ans.append(string(1,s[s.length()-1]));
                s.erase(s.length()-1,1);
                for (int i=s.length()-1;i>=0;i--){
                    if (s[i]<ans[ans.length()-1]){
                        ans.append(string(1,s[i]));
                        s.erase(i,1);
                    }
                }
                flag=true;
            }
        }
        return ans;
    }
};