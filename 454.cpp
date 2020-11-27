class Solution {
public:
    int fourSumCount(vector<int>& A, vector<int>& B, vector<int>& C, vector<int>& D) {
    int ans=0;
    map<int,int> AB;
    for (int i=0;i<A.size();i++){
        int temp=A[i];
        for(int j=0;j<B.size();j++){
            int tem=temp+B[j];
            auto iter=AB.find(tem);
            if (iter==AB.end()){
                AB.insert(pair<int,int>(tem,1));
            }else {
                iter->second++;
            }
        }
    }
    for(int m=0;m<C.size();m++){
        int temp=C[m];
        for(int n=0;n<D.size();n++){
            int tem=temp+D[n];
            auto iter=AB.find(-tem);
            if (iter==AB.end()){
                continue;
            }else {
                ans+=iter->second;
            }
        }
    }
    return ans;
}
};