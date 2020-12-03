class Solution {
public:
    int countPrimes(int n) {
        if(n<2){
            return 0;
        }
        vector<bool> flag(n,true);
        flag[0]=false;
        flag[1]=false;
        for(int i=2;i<sqrt(n);i++){
            if(flag[i]==false){
                continue;
            }else{
                for(int j=i+i;j<n;j+=i){
                    flag[j]=false;
                }
            }
        }
        int res=0;
        for(auto i:flag){
            if(i==true){
                res++;
            }
        }
        return res;
    }
};