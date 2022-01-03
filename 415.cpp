class Solution {
public:
    string addStrings(string num1, string num2) {
        reverse(num1.begin(),num1.end());
        reverse(num2.begin(),num2.end());
        int flag =0;
        int i=0;
        string ans="";
        for(;i<num1.length()&&i<num2.length();i++){
            auto temp=num1[i]+num2[i]-'0'-'0'+flag;
            flag=0;
            if (temp/10>0){
                flag=1;
            }
            temp=temp%10;
            ans+=(temp+'0');
        }
        while(i<num1.length()){
            auto temp=num1[i]-'0'+flag;
            flag=0;
            if (temp/10>0){
                flag=1;
            }
            temp=temp%10;
            ans+=(temp+'0');
            i++;
        }
        while(i<num2.length()){
            auto temp=num2[i]-'0'+flag;
            flag=0;
            if (temp/10>0){
                flag=1;
            }
            temp=temp%10;
            ans+=(temp+'0');
            i++;
        }
        if(flag){
            ans+="1";
        }
        reverse(ans.begin(),ans.end());
        return ans;
    } 
};