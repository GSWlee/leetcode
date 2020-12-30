int lastStoneWeight(vector<int>& stones) {
    while(stones.size()>1){
        auto max1=stones.begin();
        auto max2=++stones.begin();
        if(*max1<*max2){
            auto temp=max1;
            max1=max2;
            max2=temp;
        }
        for(auto it=stones.begin()+2;it<stones.end();it++){
            if(*it>=*max1){
                max2=max1;
                max1=it;
            }else if(*it>=*max2){
                max2=it;
            }else{
                continue;
            }
        }
        if(*max1==*max2){
            auto t=*max1;
            stones.erase(find(stones.begin(),stones.end(),t));
            stones.erase(find(stones.begin(),stones.end(),t));
        }else{
            *max1=*max1-*max2;
            stones.erase(find(stones.begin(),stones.end(),*max2));
        }
    }
    if(stones.size()==0){
        return 0;
    }else{
        return stones[0];
    }
}