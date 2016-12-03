package basic

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head:=&ListNode{}
	thead:=head
	for ;l1!=nil&&l2!=nil;{
		if l1.Val>l2.Val{
			head.Next=l2
			l2=l2.Next

		}else{
			head.Next=l1
			l1=l1.Next
		}

		head=head.Next
	}
	if l1==nil&&l2!=nil{
		head.Next=l2
	}else if l2==nil&&l1!=nil{
		head.Next=l1
	}
	return thead.Next
}
