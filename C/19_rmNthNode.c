#include<stdio.h>
#include<stdlib.h>

struct ListNode{
    int val;
    struct ListNode *next;
};

struct ListNode* removeNthFromEnd(struct ListNode*, int);

int main(){
    struct ListNode* head = (struct ListNode*)malloc(sizeof(struct ListNode));
    head->val = 1;
    head->next = NULL;
    head = removeNthFromEnd(head,1);
    if (head != NULL)
        printf("!!!!!!!!!!");
    else{
        printf("************");
    }
    getchar();
    return 0;
}


int len(struct ListNode* head){
    int result = 0;
    struct ListNode* ptr = head;
    while(ptr){
        ptr = ptr->next;
        result++;
    }
    return result;
}

struct ListNode* removeNthFromEnd(struct ListNode* head, int n) {
    int length = len(head);
    if(n == length)
        return head->next;
    struct ListNode* ptr = head;
    for(int i = 0; i < length - n - 1; i++)
        ptr = ptr->next;
    ptr->next = ptr->next->next;
    return head;
}