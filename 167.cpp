class Solution {
public:
    vector<int> twoSum(vector<int>& numbers, int target) {
        unordered_map<int,int> map;
        for(int i=0;i<numbers.size();i++){
            auto temp=target-numbers[i];
            auto j=map.find(temp);
            if(j!=map.end()){
                return vector<int>{map[temp]+1,i+1};
            } else{
                map[numbers[i]]=i;
            }
        }
        return vector<int>{-1,-1};
    }
};