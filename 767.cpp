class Solution {
public:
    string reorganizeString(string S) {
        int length=S.size();
        vector<int> count(26,0);
        if (length<3){
            return S;
        }
        int max=0;
        for(auto i :S){
            count[i-'a']++;
            if(count[i-'a']>max){
                max=count[i-'a'];
            }
        }
        string ans="";
        if(max>(length+1)/2){
            return ans;
        }
        auto cmp = [&](const char & a,const char & b){return count[a-'a']<count[b-'a'];};
        priority_queue <char,vector<char>,decltype(cmp)> queue(cmp);
        for(int i=0;i<26;i++){
            if (count[i]!=0){
                queue.push('a'+i);
            }
        }
        while(queue.size()>1){
            auto l1=queue.top();queue.pop();
            auto l2=queue.top();queue.pop();
            ans+=l1;
            ans+=l2;
            count[l1-'a']--;
            if (count[l1-'a']>0){
                queue.push(l1);
            }
            count[l2-'a']--;
            if (count[l2-'a']>0){
                queue.push(l2);
            }
        }
        if (queue.size() > 0) {
            ans += queue.top();
        }
        return ans;
    }
};