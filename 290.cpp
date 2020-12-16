void SplitString(const std::string& s, std::vector<std::string>& v, string c)
{
    std::string::size_type pos1, pos2;
    pos2 = s.find(c);
    pos1 = 0;
    while(std::string::npos != pos2)
    {
        v.push_back(s.substr(pos1, pos2-pos1));

        pos1 = pos2 + c.size();
        pos2 = s.find(c, pos1);
    }
    if(pos1 != s.length())
        v.push_back(s.substr(pos1));
}
bool wordPattern(string pattern, string s) {
    unordered_map<char,string> map1;
    unordered_map<string,char> map2;
    vector<string> vStr;
    SplitString(s,vStr," ");
    if(vStr.size()!=pattern.size()){
        return false;
    }else{
        for(int i=0;i<vStr.size();i++){
            auto flag1 = map1.find (pattern[i]);
            auto flag2 = map2.find(vStr[i]);
            if ( flag1 == map1.end()&&flag2==map2.end() ) {
                map1[pattern[i]]=vStr[i];
                map2[vStr[i]]=pattern[i];
            }else if(flag1 != map1.end()&&flag2!=map2.end() ){
                if(flag1->second!=vStr[i]||flag2->second!=pattern[i]){
                    return false;
                }
            }else{
                return false;
            }
        }
    }
    return true;
}