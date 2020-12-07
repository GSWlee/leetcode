class Solution {
public:
    int btoi(vector<int> bits){
        int num=0;
        for(int i=bits.size()-1;i>=0;i--){
            if (bits[i]==1){
                num+=1<<(bits.size()-1-i);
            }
        }
        return num;
    }
    void reverse(vector<int> &A){
        for(int i=0;i<A.size();i++){
            if(A[i]==0){
                A[i]=1;
            }else{
                A[i]=0;
            }
        }
    }
    void reverse(vector<vector<int>>& A,int j){
        int one=0,zero=0;
        for(int i=0;i<A.size();i++){
            if(A[i][j]==0){
                zero++;
            }else{
                one++;
            }
        }
        if(zero>one){
            for(int i=0;i<A.size();i++){
                if(A[i][j]==0){
                    A[i][j]=1;
                }else{
                    A[i][j]=0;
                }
            }
        }
    }
    int matrixScore(vector<vector<int>>& A) {
        for(int i=0;i<A.size();i++){
            if(A[i][0]==0){
                reverse(A[i]);
            }
        }
        for(int j=0;j<A[0].size();j++){
            reverse(A,j);
        }
        int ans=0;
        for(auto i:A){
            ans+=btoi(i);
        }
        return ans;
    }
};