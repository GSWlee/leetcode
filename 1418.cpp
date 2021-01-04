vector<vector<string>> displayTable(vector<vector<string>>& orders) {
    map<string,int> m;
    int sum=0;
    for(auto i : orders){
        if(m.find(i[2])==m.end()){
            m[i[2]]=sum;
            sum++;
        }
    }

    int q=0;
    for(auto& i : m){
        i.second=q;
        q++;
    }
    map<int,vector<int>> set;
    for(auto i :orders){
        if(set.find(stoi(i[1]))==set.end()){
            set[stoi(i[1])]=vector<int>(m.size()+1,0);
            set[stoi(i[1])][0]=stoi(i[1]);
            set[stoi(i[1])][m[i[2]]+1]++;
        }else{
            set[stoi(i[1])][m[i[2]]+1]++;
        }
    }
    vector<vector<string>> ans;
    vector<string> first;
    first.push_back("Table");
    for(auto i : m){
        first.push_back(i.first);
    }
    ans.push_back(first);
    for(auto i : set){
        vector<string> temp;
        for(auto j : i.second){
            temp.push_back(to_string(j));
        }
        ans.push_back(temp);
    }
    return ans;
}