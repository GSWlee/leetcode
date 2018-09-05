/*
 * 寻找两个数组的中位数。
 * 参照算法导论中的中位数算法，结合了快排的思想
 * */

#include <iostream>
#include <vector>
#include <string>

using namespace std;

void swap(vector<int>& num,int a,int b){
    int t=num[a];
    num[a]=num[b];
    num[b]=t;
    return;
}

int sortn(vector<int>& num,int begin,int end){
    int x = num[begin];
    for(int i = begin+1;i<end;){
        if(num[i]<=x)
            i++;
        else {
            swap(num,i,--end);
        }
    }
    swap(num,begin,end-1);
    return end-1;
}

float findz(vector<int>& num,int begin,int end, int target) {
    if (begin==end)
        return begin;
    int k=sortn(num,begin,end);
    if (k==target){
        return k;
    } else if(k<target){
        return findz(num,k+1,end,target);
    } else{
        return findz(num,begin,k,target);
    }
}

int main() {
    int a[1]={2};
    vector<int> nums1(a,a+1);
    //int b[9]={10,11,12,13,14,15,16,17,18};
    vector<int> nums2;//(b,b+9);
    vector<int> num;
    num.insert(num.end(),nums1.begin(),nums1.end());
    num.insert(num.end(),nums2.begin(),nums2.end());
    float ans;
    int  q=num.size()/2;
    ans= num.size()%2 ? ans=num[findz(num,0,num.size(),q)]:ans=(num[findz(num,0,num.size(),q+1)]+num[findz(num,0,num.size(),q)])/2;
    cout<<ans;
    return 0;
}
