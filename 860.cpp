class Solution {
public:
    bool lemonadeChange(vector<int>& bills) {
        int five=0,ten=0;
        for(auto temp:bills){
            if(temp==5){
                five++;
            }else if(temp==10){
                if(five<1){
                    return false;
                }else{
                    five--;
                    ten++;
                }
            }else{
                if(ten>0&&five>0){
                    ten--;
                    five--;
                }else if(ten<1&&five>2){
                    five--;
                    five--;
                    five--;
                }else{
                    return false;
                }
            }
        }
        return true;
    }
};