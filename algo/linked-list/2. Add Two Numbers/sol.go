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
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    t1 := l1
    t2 := l2
    var size1, size2 int
    
    for t1 != nil {
        t1 = t1.Next;
        size1++;
    }
    for t2 != nil {
        t2 = t2.Next;
        size2++;
    }
    if size1 >= size2 {
        return addList(l1, l2)
    }
    return addList(l2, l1)
}

func addList(l1 *ListNode, l2 *ListNode) *ListNode {
    res := l1
    t1  := l1
    t2  := l2
    
    var carry, val int
    
    for t2 != nil {
        val = t1.Val + t2.Val + carry
        t1.Val = val % 10
        
        if val >= 10 {
            carry = 1
        } else {
            carry = 0
        }
        
        if t1.Next == nil && carry == 1 {
            t1.Next = new(ListNode)
            t1.Next.Val = carry
            return res
        } 
        t2 = t2.Next
        t1 = t1.Next
    }
    for t1 != nil {
        val = t1.Val + carry
        t1.Val = val % 10
        
        if val >= 10 {
            carry = 1
        } else {
            carry = 0
        }
        
        if t1.Next == nil && carry == 1 {
            t1.Next = new(ListNode)
            t1.Next.Val = carry
            return res
        } 
        t1 = t1.Next
    }
    return res
}
~
