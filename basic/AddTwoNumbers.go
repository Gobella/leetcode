package basic

/*
You are given two linked lists representing two non-negative numbers. The digits are stored
 in reverse order and each of their nodes contain a single digit. Add the two numbers and
 return it as a linked list.
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head,tail,currentNode *ListNode
	jinwei := 0
	if  l1==nil && l2 == nil {
		return nil;
	}
	for (l1!=nil&&l2!=nil){
		current :=l1.Val + l2.Val + jinwei
		if head==nil  {
			head = &ListNode{current%10,nil}
			tail = head
			currentNode = head
			jinwei =  current /10
		}else{
			currentNode = &ListNode{current%10,nil}
			tail.Next=currentNode
			tail = currentNode
			jinwei = current/10
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1!=nil{
		current :=l1.Val + jinwei
		if head==nil  {
			head = &ListNode{current%10,nil}
			tail = head
			currentNode = head
			jinwei = current /10
		}else{
			currentNode = &ListNode{current%10,nil}
			tail.Next=currentNode
			tail = currentNode
			jinwei = current/10
		}
		l1 = l1.Next
	}
	for l2!=nil{
		current :=l2.Val + jinwei
		if head==nil  {
			head = &ListNode{current%10,nil}
			tail = head
			currentNode = head
			jinwei = current /10
		}else{
			currentNode = &ListNode{current%10,nil}
			tail.Next=currentNode
			tail = currentNode
			jinwei = current/10
		}
		l2 = l2.Next
	}
	if jinwei!=0{
		current := jinwei
		if head==nil  {
			head = &ListNode{current%10,nil}
			tail = head
			currentNode = head
			jinwei = current /10
		}else{
			currentNode = &ListNode{current%10,nil}
			tail.Next=currentNode
			tail = currentNode
			jinwei = current/10
		}
	}
	return head


}