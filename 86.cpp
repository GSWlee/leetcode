ListNode* partition(ListNode* head, int x) {
    ListNode* small= new ListNode;
    ListNode* large= new ListNode;
    small->next=NULL;
    large->next=NULL;
    auto temp=head;
    auto templ=large;
    auto temps=small;
    while(temp!=NULL){
        if(temp->val<x){
            temps->next=temp;
            temp=temp->next;
            temps=temps->next;
        }else{
            templ->next=temp;
            temp=temp->next;
            templ=templ->next;
        }
    }
    temps->next=large->next;
    templ->next=NULL;
    return small->next;
}