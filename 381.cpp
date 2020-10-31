class RandomizedCollection {
public:
    vector<int> nums;
    int sum;
    map<int,vector<int>> tag;
    /** Initialize your data structure here. */
    RandomizedCollection() {
        sum=0;
    }
    
    /** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
    bool insert(int val) {
        nums.push_back(val);
        sum++;
        auto iter=tag.find(val);
        if (iter!=tag.end()){
            tag[val].push_back(nums.size()-1);
            if (tag[val].size()==1){
                return true;
            }else{
                return false;
            }
        }else{
            tag.insert((pair<int, vector<int>>(val, vector<int>(1,nums.size()-1))));
            return true;
        }
        
    }
    
    /** Removes a value from the collection. Returns true if the collection contained the specified element. */
    bool remove(int val) {
        auto iter=tag.find(val);
        if (iter==tag.end()){
            return false;
        }else{
            if (tag[val].size()==0){
                return false;
            }else{
                auto temp_v=nums[sum-1];
                sum--;
                nums.pop_back();
                auto temp_p=tag[val].back();
                tag[val].pop_back();
                if (sum==temp_p){
                    return true;
                }else{
                    nums[temp_p]=temp_v;
                    tag[temp_v].pop_back();
                    tag[temp_v].push_back(temp_p);
                    for(int i=tag[temp_v].size()-1;i>0;i--){
                        if(tag[temp_v][i]<tag[temp_v][i-1]){
                            auto temp=tag[temp_v][i-1];
                            tag[temp_v][i-1]=tag[temp_v][i];
                            tag[temp_v][i]=temp;
                        }
                    }
                    return true;
                }
            }
        }
    }
        
    
    
    /** Get a random element from the collection. */
    int getRandom() {
        return nums[rand()%sum];
    }
};