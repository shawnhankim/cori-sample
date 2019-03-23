/*

2. Add Two Numbers

Medium

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.

*/

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        ListNode *t1 = l1, *t2 = l2;
        int size1 = 0, size2 = 0;
        while(t1) { t1 = t1->next; ++size1; }
        while(t2) { t2 = t2->next; ++size2; }
        
        return size1 >= size2 ? addNode(l1, l2) : addNode(l2, l1);
    }
    
    ListNode* addNode(ListNode* l1, ListNode* l2) {
        ListNode* res = l1;
        ListNode* t1 = l1, *t2 = l2;
        int val = 0, carry = 0;
        
        while(t2) {
            val = t1->val + t2->val + carry;
            t1->val = val % 10;
            if (val >= 10) carry = 1;
            else           carry = 0;
            if (!t1->next && carry == 1) {
                t1->next = new ListNode(1);
                return res;
            }
            t1 = t1->next;
            t2 = t2->next;
        }
        
        while(t1) {
            val = t1->val + carry;
            t1->val = val % 10;
            if (val >= 10) carry = 1;
            else           carry = 0;
            
            if (!t1->next && carry == 1) {
                t1->next = new ListNode(1);
                return res;
            }
            t1 = t1->next;
        }
        return res;
    }
};
