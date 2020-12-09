class Solution {
public:
    int uniquePaths(int m, int n) {
        vector<vector<int>> map(m,vector<int>(n,0));
        for(int i=0;i<m;i++){
            for(int j=0;j<n;j++){
                if(i==0){
                    if(j==0){
                        map[i][j]=1;
                    }else{
                        map[i][j]=map[i][j-1];
                    }
                }else{
                    if(j==0){
                        map[i][j]=map[i-1][j];
                    }else{
                        map[i][j]=map[i-1][j]+map[i][j-1];
                    }
                }
            }
        }
        return map[m-1][n-1];
    }
};