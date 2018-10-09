class Solution {
public:
    bool isPalindrome(int x) {
        int i=0;
        long num=1;
        while(num<=x){
            num*=10;
            i++;
        }
        int q=x;
        i--;
        int ans=0;
        for(int j=i;j>=0;j--){
            ans+=int((q/pow(10,j)))*pow(10,i-j);
            q=q%int(pow(10,j));
        }
        return x==ans;
    }
};