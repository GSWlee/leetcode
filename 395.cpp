class Solution {
public:
    int longestSubstring(string s, int k) {
            bool flag=true;
    vector<string> temp{s};
    while (flag){
        flag=false;
        vector<string> newtemp;
        for(int i=0;i<temp.size();i++){
            unordered_map<char,int> map;
            for(int j=0;j<temp[i].size();j++){
                map[temp[i][j]]+=1;
            }
            vector<int> position;
            for(int j=0;j<temp[i].size();j++){
                if(map[temp[i][j]]<k){
                    position.push_back(j);
                    flag=true;
                }
            }
            int target=0;
            for(auto item : position){
                if(item-target<k){
                    target=item+1;
                }else{
                    newtemp.emplace(newtemp.end(),temp[i].begin()+target,temp[i].begin()+item);
                    target=item+1;
                }
            }
            if(temp[i].size()-target>=k){
                newtemp.emplace(newtemp.end(),temp[i].begin()+target,temp[i].end());
            }
        }
        if(flag){
            temp=newtemp;
        }
    }
    int max=0;
    for(auto item:temp){
        if(max<item.size()){
            max=item.size();
        }
    }
    return max;
}                
};