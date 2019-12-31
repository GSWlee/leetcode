//利用位运算来判断相同的数是否出现，如果出现两次，这ａ与ａ异或结果不会变
int singleNumber(vector<int>& nums) {
    int flag=0;
    for(auto i : nums)
        flag=flag^i;
    return flag;
}