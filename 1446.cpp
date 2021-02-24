class Solution {
public:
    int maxPower(string s) {
        int max=0,now=0;
        auto temp=s[0];
        for(int i=0;i<s.size();i++){
            if(s[i]==temp){
                now++;
            }else{
                temp=s[i];
                if(now>max){
                    max=now;
                }
                now=1;
            }
        }
        return max>now?max:now;
    }
};