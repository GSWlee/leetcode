vector<vector<string>> groupAnagrams(vector<string>& strs) {
        vector<vector<string>> ans;
        unordered_map<string,vector<string>> mp;
        for(auto str :strs){
            auto key=str;
            sort(key.begin(),key.end());
            mp[key].emplace_back(str);
        }
        for(auto it=mp.begin();it!=mp.end();it++){
            ans.emplace_back(it->second);
        }
        return ans;
    }