class Solution {
public:
    bool cmp(vector<int> a,vector<int> b){
        return a[1]<b[1];
    } 
    int findMinArrowShots(vector<vector<int>>& points) {
        if (points.size()==0){
            return 0;
        }
        sort(points.begin(),points.end(),[](const vector<int>& a,const vector<int>& b){
            return a[1]<b[1];
        });
        int ans=1,arrow=points[0][1];
        for (int i=0;i<points.size();i++){
            if (points[i][0]>arrow){
                ans++;
                arrow=points[i][1];
            }else{
                continue;
            }
        }
        return ans;
    }
};