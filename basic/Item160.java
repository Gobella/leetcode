package leetcode;

public class Item160 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub

	}
	public ListNode getIntersectionNode(ListNode headA, ListNode headB) {
		if( null==headA || null==headB )
				return null;
			
		ListNode curA = headA, curB = headB;
		while( curA!=curB){
			curA = curA==null?headB:curA.next;
			curB = curB==null?headA:curB.next;
		}
		return curA;
    }

}
