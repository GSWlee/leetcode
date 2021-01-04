class Solution {
public:
    int fib(int n) {
        int l2=0,l1=1;
        if(n==0){
            return 0;
        }else if(n==1){
            return 1;
        }else{
            while(n>2){
                n--;
                auto temp=l1;
                l1=l1+l2;
                l2=temp;
            }
            return l1+l2;
        }
    }
};