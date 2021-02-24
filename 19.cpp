class Solution {
public:
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        ListNode* fast=head;
        ListNode* low=head;
        for(int i=0;i<n;i++){
            fast=fast->next;
        }
        if(!fast){
            return head->next;
        }
        while(fast->next!=nullptr){
            fast=fast->next;
            low=low->next;
        }
        low->next=low->next->next;
        return head;
    }
};