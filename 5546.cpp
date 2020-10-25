class Solution {
public:
    char slowestKey(vector<int>& releaseTimes, string keysPressed) {
        long long a[26];
        for (int j=0;j<26;j++){
            a[j]=0;
        }
        for (int i=0;i<releaseTimes.size();i++){
            if (i==0){
                a[keysPressed[i]-'a']=releaseTimes[i];
            }else if(a[keysPressed[i]-'a']<(releaseTimes[i]-releaseTimes[i-1])){
                a[keysPressed[i]-'a']=releaseTimes[i]-releaseTimes[i-1];
            }
        }
        int max=0;
        for (int j=0;j<26;j++){
            if(a[max]<=a[j]){
                max=j;
            }
        }
        return 'a'+max;
    }
};