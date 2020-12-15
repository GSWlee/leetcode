class Solution {
public:
    int monotoneIncreasingDigits(int N) {
        string strN = to_string(N);
        for(int i=strN.size()-1;i>0;i--){
            if(strN[i]>=strN[i-1]){
                continue;
            } else{
                strN[i-1]--;
                for(int j=i;j<strN.size();j++){
                    strN[j]='9';
                }
            }
        }
        return stoi(strN);
    }
};