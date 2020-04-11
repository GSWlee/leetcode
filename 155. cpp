class MinStack {
public:
    /** initialize your data structure here. */
    int a[10000];
    int min=-1;
    int t=0;

    void push(int x) {
        a[t++]=x;
        if(min==-1){
            min=t-1;
       }else{
           if(x<=a[min]){
               min=t-1;
           }
       }
    }

    void pop() {
        t--;
        if(t==0){
            min=-1;
        }else{
            min=0;
            for(int i=1;i<t;i++){
                if(a[i]<=a[min]){
                    min=i;
                }
            }

        }

    }
    
    int top() {
        return a[t-1];
    }
    
    int getMin() {
        return a[min];
    }
};
