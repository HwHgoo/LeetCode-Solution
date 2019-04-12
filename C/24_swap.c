#include<stdio.h>

struct ListNode {
    int val;
    struct ListNode *next;
};

struct ListNode* swapPairs(struct ListNode*);
struct ListNode* genrateList(int*,int);

int main(){
    int a[4] = {1,2,3,4};
    struct ListNode* list = genrateList(a,4);
    swapPairs(list);
    return 0;
}

struct ListNode* genrateList(int* nums, int n){
    struct ListNode* head = (struct ListNode*)malloc(sizeof(struct ListNode));
    head->val = *nums;
    struct ListNode* ptr = head;
    int index = 0;
    while(index < n - 1){
        struct ListNode* node = (struct ListNode*)malloc(sizeof(struct ListNode));
        node->val = *(nums + index + 1);
        ptr->next = node;
        node->next = NULL;
        ptr = ptr->next;
        index++;
    }
    return head;
}

struct ListNode* swapPairs(struct ListNode* head) {
    if (!head || !head->next)
        return head;
    struct ListNode* ptr = head->next;
    head->next = ptr->next;
    ptr->next = head;
    head->next = swapPairs(head->next);
    return ptr;
}