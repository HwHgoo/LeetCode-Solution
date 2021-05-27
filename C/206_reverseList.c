
struct ListNode {
	int val;
	struct ListNode *next;
};

struct ListNode *reverseList(struct ListNode *head) {
	int * vals = (int*)malloc(sizeof(int) * 5000);
	struct ListNode* ltmp = head;
	int count = 0;
	while(ltmp){
		vals[count++] = head->val;
		ltmp = ltmp->next;
	}
	ltmp = head;
	for(int i = count - 1; i >= 0; --i){
		ltmp->val = vals[i];
		ltmp = ltmp->next;
	}
	return head;
}
