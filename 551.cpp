class Solution {
public:
    bool checkRecord(string s) {
        int A=0,L=0;
        for(auto i : s){
            if(A==2||L==3){
                return false;
            }
            if(i=='L'){
                L++;
            }else if(i=='A'){
                A++;
                L=0;
            }else{
                L=0;
            }
        }
        if(A==2||L==3){
            return false;
        }else{
            return true;
        }
    }
};