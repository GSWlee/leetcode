class Solution {
public:
    int hammingDistance(int x, int y) {
        int count=0;
        for(int i=0;i<31;i++){
            if((x%2)!=(y%2)){
                count++;
            }
            x=x>>1;
            y=y>>1;
        }
        return count;
    }
};
