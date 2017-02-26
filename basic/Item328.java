package leetcode;

//
//Given a singly linked list, group all odd nodes together followed by the even nodes. Please note here we are talking about the node number and not the value in the nodes.
//
//You should try to do it in place. The program should run in O(1) space complexity and O(nodes) time complexity.
//
//Example:
//Given 1->2->3->4->5->NULL,
//return 1->3->5->2->4->NULL.
//
//Note:
//The relative order inside both the even and odd groups should remain as it was in the input. 
//The first node is considered odd, the second node even and so on ...

public class Item328 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub

	}
    public ListNode oddEvenList(ListNode head) {
        if(head ==null)
            return head;
        ListNode headc=head.next;
        ListNode even=head.next;
        ListNode oc=head;
        ListNode ec=head.next;
        while(headc!=null){
            if(ec!=null){
                oc.next=ec.next;
                if(oc.next!=null)
                    oc=oc.next;
                headc=headc.next;
            }
            if(ec!=null&&headc!=null){
                ec.next=oc.next;
                ec=ec.next;
                headc=headc.next;
            }
        }
        oc.next=even;
        if(ec!=null)
            ec.next=null;
        return head;
    }

}
